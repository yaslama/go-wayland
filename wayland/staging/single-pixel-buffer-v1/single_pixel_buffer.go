// Generated by go-wayland-scanner
// https://github.com/yaslama/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/staging/single-pixel-buffer/single-pixel-buffer-v1.xml?ref_type=tags
//
// single_pixel_buffer_v1 Protocol Copyright:
//
// Copyright © 2022 Simon Ser
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

package single_pixel_buffer

import "github.com/yaslama/go-wayland/wayland/client"

// WpSinglePixelBufferManagerInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const WpSinglePixelBufferManagerInterfaceName = "wp_single_pixel_buffer_manager_v1"

// WpSinglePixelBufferManager : global factory for single-pixel buffers
//
// The wp_single_pixel_buffer_manager_v1 interface is a factory for
// single-pixel buffers.
type WpSinglePixelBufferManager struct {
	client.BaseProxy
}

// NewWpSinglePixelBufferManager : global factory for single-pixel buffers
//
// The wp_single_pixel_buffer_manager_v1 interface is a factory for
// single-pixel buffers.
func NewWpSinglePixelBufferManager(ctx *client.Context) *WpSinglePixelBufferManager {
	wpSinglePixelBufferManagerV1 := &WpSinglePixelBufferManager{}
	ctx.Register(wpSinglePixelBufferManagerV1)
	return wpSinglePixelBufferManagerV1
}

// Destroy : destroy the manager
//
// Destroy the wp_single_pixel_buffer_manager_v1 object.
//
// The child objects created via this interface are unaffected.
func (i *WpSinglePixelBufferManager) Destroy() error {
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

// CreateU32RgbaBuffer : create a 1×1 buffer from 32-bit RGBA values
//
// Create a single-pixel buffer from four 32-bit RGBA values.
//
// Unless specified in another protocol extension, the RGBA values use
// pre-multiplied alpha.
//
// The width and height of the buffer are 1.
//
//	r: value of the buffer's red channel
//	g: value of the buffer's green channel
//	b: value of the buffer's blue channel
//	a: value of the buffer's alpha channel
func (i *WpSinglePixelBufferManager) CreateU32RgbaBuffer(r, g, b, a uint32) (*client.Buffer, error) {
	id := client.NewBuffer(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(r))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(g))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(b))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(a))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}
