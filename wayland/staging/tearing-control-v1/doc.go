package tearing_control

//go:generate go run github.com/yaslama/go-wayland/cmd/go-wayland-scanner -pkg tearing_control -prefix wp -suffix v1 -o tearing_control.go -i https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/staging/tearing-control/tearing-control-v1.xml?ref_type=tags
