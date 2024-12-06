package main

import (
	"fmt"
	"github.com/MatthiasKunnen/go-wayland/wayland/client"
	idlenotify "github.com/MatthiasKunnen/go-wayland/wayland/staging/ext-idle-notify-v1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Global app state
type appState struct {
	exit     bool
	display  *client.Display
	registry *client.Registry
	seat     *client.Seat
	notifier *idlenotify.IdleNotifier
}

func main() {
	app := &appState{}

	app.initWindow()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		app.exit = true
		app.cleanup()
		os.Exit(1)
	}()

	for {
		app.dispatch()
	}
}

func (app *appState) initWindow() {
	// Connect to wayland server
	display, err := client.Connect("")
	if err != nil {
		log.Fatalf("unable to connect to wayland server: %v", err)
	}
	app.display = display

	display.SetErrorHandler(app.HandleDisplayError)

	// Get global interfaces registry
	registry, err := app.display.GetRegistry()
	if err != nil {
		log.Fatalf("unable to get global registry object: %v", err)
	}
	app.registry = registry

	// Add global interfaces registrar handler
	registry.SetGlobalHandler(app.HandleRegistryGlobal)
	// Wait for interfaces to register
	app.displayRoundTrip()
	// Wait for handler events
	app.displayRoundTrip()

	notification, err := app.notifier.GetIdleNotification(5000, app.seat)
	if err != nil {
		log.Fatalf("Error getting idle notification: %v", err)
	}
	notification.SetIdledHandler(func(event idlenotify.IdleNotificationIdledEvent) {
		fmt.Printf("Idle\n")
	})
	notification.SetResumedHandler(func(event idlenotify.IdleNotificationResumedEvent) {
		fmt.Printf("Resumed\n")
	})
}

func (app *appState) dispatch() {
	err := app.display.Context().Dispatch()
	if err != nil && !app.exit {
		log.Printf("Error dispatching: %v\n", err)
	}
}

func (app *appState) context() *client.Context {
	return app.display.Context()
}

func (app *appState) HandleRegistryGlobal(e client.RegistryGlobalEvent) {
	switch e.Interface {
	case idlenotify.IdleNotifierInterfaceName:
		app.notifier = idlenotify.NewIdleNotifier(app.context())
		err := app.registry.Bind(e.Name, e.Interface, e.Version, app.notifier)
		if err != nil {
			log.Fatalf("Unable to bind %s interface: %v", idlenotify.IdleNotifierInterfaceName, err)
		}
	case client.SeatInterfaceName:
		seat := client.NewSeat(app.context())
		err := app.registry.Bind(e.Name, e.Interface, e.Version, seat)
		if err != nil {
			log.Fatalf("unable to bind %s interface: %v", client.SeatInterfaceName, err)
		}
		app.seat = seat
	}
}

// HandleDisplayError handles client.Display errors
func (*appState) HandleDisplayError(e client.DisplayErrorEvent) {
	// Just log.Fatal for now
	log.Fatalf("display error event: %v", e)
}

func (app *appState) displayRoundTrip() {
	// Get display sync callback
	callback, err := app.display.Sync()
	if err != nil {
		log.Fatalf("unable to get sync callback: %v", err)
	}
	defer func() {
		if err2 := callback.Destroy(); err2 != nil {
			log.Println("unable to destroy callback:", err2)
		}
	}()

	done := false
	callback.SetDoneHandler(func(_ client.CallbackDoneEvent) {
		done = true
	})

	// Wait for callback to return
	for !done {
		app.dispatch()
	}
}

func (app *appState) cleanup() {
	// Release wl_seat handlers
	if app.seat != nil {
		if err := app.seat.Release(); err != nil {
			log.Println("unable to release wl_seat:", err)
		}
		app.seat = nil
	}

	if app.registry != nil {
		if err := app.registry.Destroy(); err != nil {
			log.Println("unable to destroy wl_registry:", err)
		}
		app.registry = nil
	}

	if app.display != nil {
		if err := app.display.Destroy(); err != nil {
			log.Println("unable to destroy wl_display:", err)
		}
	}

	if app.notifier != nil {
		if err := app.notifier.Destroy(); err != nil {
			log.Printf("unable to destroy %s, %v:\n", idlenotify.IdleNotifierInterfaceName, err)
		}
	}

	// Close the wayland server connection
	if err := app.context().Close(); err != nil {
		log.Println("unable to close wayland context:", err)
	}
}
