package tablet

//go:generate go run github.com/yaslama/go-wayland/cmd/go-wayland-scanner -pkg tablet -prefix zwp -suffix v2 -o tablet.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/tablet/tablet-unstable-v2.xml?ref_type=tags
