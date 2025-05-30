// Generated by go-wayland-scanner
// https://github.com/yaslama/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/idle-inhibit/idle-inhibit-unstable-v1.xml?ref_type=tags
//
// idle_inhibit_unstable_v1 Protocol Copyright:
//
// Copyright © 2015 Samsung Electronics Co., Ltd
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice (including the next
// paragraph) shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package idle_inhibit

import "github.com/yaslama/go-wayland/wayland/client"

// IdleInhibitManagerInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const IdleInhibitManagerInterfaceName = "zwp_idle_inhibit_manager_v1"

// IdleInhibitManager : control behavior when display idles
//
// This interface permits inhibiting the idle behavior such as screen
// blanking, locking, and screensaving.  The client binds the idle manager
// globally, then creates idle-inhibitor objects for each surface.
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
type IdleInhibitManager struct {
	client.BaseProxy
}

// NewIdleInhibitManager : control behavior when display idles
//
// This interface permits inhibiting the idle behavior such as screen
// blanking, locking, and screensaving.  The client binds the idle manager
// globally, then creates idle-inhibitor objects for each surface.
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
func NewIdleInhibitManager(ctx *client.Context) *IdleInhibitManager {
	zwpIdleInhibitManagerV1 := &IdleInhibitManager{}
	ctx.Register(zwpIdleInhibitManagerV1)
	return zwpIdleInhibitManagerV1
}

// Destroy : destroy the idle inhibitor object
//
// Destroy the inhibit manager.
func (i *IdleInhibitManager) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// CreateInhibitor : create a new inhibitor object
//
// Create a new inhibitor object associated with the given surface.
//
//	surface: the surface that inhibits the idle behavior
func (i *IdleInhibitManager) CreateInhibitor(surface *client.Surface) (*IdleInhibitor, error) {
	id := NewIdleInhibitor(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], surface.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// IdleInhibitorInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const IdleInhibitorInterfaceName = "zwp_idle_inhibitor_v1"

// IdleInhibitor : context object for inhibiting idle behavior
//
// An idle inhibitor prevents the output that the associated surface is
// visible on from being set to a state where it is not visually usable due
// to lack of user interaction (e.g. blanked, dimmed, locked, set to power
// save, etc.)  Any screensaver processes are also blocked from displaying.
//
// If the surface is destroyed, unmapped, becomes occluded, loses
// visibility, or otherwise becomes not visually relevant for the user, the
// idle inhibitor will not be honored by the compositor; if the surface
// subsequently regains visibility the inhibitor takes effect once again.
// Likewise, the inhibitor isn't honored if the system was already idled at
// the time the inhibitor was established, although if the system later
// de-idles and re-idles the inhibitor will take effect.
type IdleInhibitor struct {
	client.BaseProxy
}

// NewIdleInhibitor : context object for inhibiting idle behavior
//
// An idle inhibitor prevents the output that the associated surface is
// visible on from being set to a state where it is not visually usable due
// to lack of user interaction (e.g. blanked, dimmed, locked, set to power
// save, etc.)  Any screensaver processes are also blocked from displaying.
//
// If the surface is destroyed, unmapped, becomes occluded, loses
// visibility, or otherwise becomes not visually relevant for the user, the
// idle inhibitor will not be honored by the compositor; if the surface
// subsequently regains visibility the inhibitor takes effect once again.
// Likewise, the inhibitor isn't honored if the system was already idled at
// the time the inhibitor was established, although if the system later
// de-idles and re-idles the inhibitor will take effect.
func NewIdleInhibitor(ctx *client.Context) *IdleInhibitor {
	zwpIdleInhibitorV1 := &IdleInhibitor{}
	ctx.Register(zwpIdleInhibitorV1)
	return zwpIdleInhibitorV1
}

// Destroy : destroy the idle inhibitor object
//
// Remove the inhibitor effect from the associated wl_surface.
func (i *IdleInhibitor) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}
