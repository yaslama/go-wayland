// Generated by go-wayland-scanner
// https://github.com/yaslama/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.44/staging/xdg-system-bell/xdg-system-bell-v1.xml?ref_type=tags
//
// xdg_system_bell_v1 Protocol Copyright:
//
// Copyright © 2016, 2023 Red Hat
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

package xdg_system_bell

import "github.com/yaslama/go-wayland/wayland/client"

// XdgSystemBellInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const XdgSystemBellInterfaceName = "xdg_system_bell_v1"

// XdgSystemBell : system bell
//
// This global interface enables clients to ring the system bell.
//
// Warning! The protocol described in this file is currently in the testing
// phase. Backward compatible changes may be added together with the
// corresponding interface version bump. Backward incompatible changes can
// only be done by creating a new major version of the extension.
type XdgSystemBell struct {
	client.BaseProxy
}

// NewXdgSystemBell : system bell
//
// This global interface enables clients to ring the system bell.
//
// Warning! The protocol described in this file is currently in the testing
// phase. Backward compatible changes may be added together with the
// corresponding interface version bump. Backward incompatible changes can
// only be done by creating a new major version of the extension.
func NewXdgSystemBell(ctx *client.Context) *XdgSystemBell {
	xdgSystemBellV1 := &XdgSystemBell{}
	ctx.Register(xdgSystemBellV1)
	return xdgSystemBellV1
}

// Destroy : destroy the system bell object
//
// Notify that the object will no longer be used.
func (i *XdgSystemBell) Destroy() error {
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

// Ring : ring the system bell
//
// This requests rings the system bell on behalf of a client. How ringing
// the bell is implemented is up to the compositor. It may be an audible
// sound, a visual feedback of some kind, or any other thing including
// nothing.
//
// The passed surface should correspond to a toplevel like surface role,
// or be null, meaning the client doesn't have a particular toplevel it
// wants to associate the bell ringing with. See the xdg-shell protocol
// extension for a toplevel like surface role.
//
//	surface: associated surface
func (i *XdgSystemBell) Ring(surface *client.Surface) error {
	const opcode = 1
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	if surface == nil {
		client.PutUint32(_reqBuf[l:l+4], 0)
		l += 4
	} else {
		client.PutUint32(_reqBuf[l:l+4], surface.ID())
		l += 4
	}
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}
