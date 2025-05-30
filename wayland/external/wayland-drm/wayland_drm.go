// Generated by go-wayland-scanner
// https://github.com/yaslama/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/mesa/mesa/-/raw/25.0/src/egl/wayland/wayland-drm/wayland-drm.xml
//
// drm Protocol Copyright:
//
// Copyright © 2008-2011 Kristian Høgsberg
// Copyright © 2010-2011 Intel Corporation
//
// Permission to use, copy, modify, distribute, and sell this
// software and its documentation for any purpose is hereby granted
// without fee, provided that\n the above copyright notice appear in
// all copies and that both that copyright notice and this permission
// notice appear in supporting documentation, and that the name of
// the copyright holders not be used in advertising or publicity
// pertaining to distribution of the software without specific,
// written prior permission.  The copyright holders make no
// representations about the suitability of this software for any
// purpose.  It is provided "as is" without express or implied
// warranty.
//
// THE COPYRIGHT HOLDERS DISCLAIM ALL WARRANTIES WITH REGARD TO THIS
// SOFTWARE, INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS, IN NO EVENT SHALL THE COPYRIGHT HOLDERS BE LIABLE FOR ANY
// SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN
// AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION,
// ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
// THIS SOFTWARE.

package wayland_drm

import (
	"github.com/yaslama/go-wayland/wayland/client"
	"golang.org/x/sys/unix"
)

// DrmInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const DrmInterfaceName = "wl_drm"

// Drm :
type Drm struct {
	client.BaseProxy
	deviceHandler        DrmDeviceHandlerFunc
	formatHandler        DrmFormatHandlerFunc
	authenticatedHandler DrmAuthenticatedHandlerFunc
	capabilitiesHandler  DrmCapabilitiesHandlerFunc
}

// NewDrm :
func NewDrm(ctx *client.Context) *Drm {
	wlDrm := &Drm{}
	ctx.Register(wlDrm)
	return wlDrm
}

// Authenticate :
func (i *Drm) Authenticate(id uint32) error {
	const opcode = 0
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(id))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// CreateBuffer :
func (i *Drm) CreateBuffer(name uint32, width, height int32, stride, format uint32) (*client.Buffer, error) {
	id := client.NewBuffer(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4 + 4 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(name))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(width))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(height))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(stride))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(format))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// CreatePlanarBuffer :
func (i *Drm) CreatePlanarBuffer(name uint32, width, height int32, format uint32, offset0, stride0, offset1, stride1, offset2, stride2 int32) (*client.Buffer, error) {
	id := client.NewBuffer(i.Context())
	const opcode = 2
	const _reqBufLen = 8 + 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(name))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(width))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(height))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(format))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(offset0))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(stride0))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(offset1))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(stride1))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(offset2))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(stride2))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// CreatePrimeBuffer :
func (i *Drm) CreatePrimeBuffer(name int, width, height int32, format uint32, offset0, stride0, offset1, stride1, offset2, stride2 int32) (*client.Buffer, error) {
	id := client.NewBuffer(i.Context())
	const opcode = 3
	const _reqBufLen = 8 + 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(width))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(height))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(format))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(offset0))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(stride0))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(offset1))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(stride1))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(offset2))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(stride2))
	l += 4
	oob := unix.UnixRights(int(name))
	err := i.Context().WriteMsg(_reqBuf[:], oob)
	return id, err
}

func (i *Drm) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

type DrmError uint32

// DrmError :
const (
	DrmErrorAuthenticateFail DrmError = 0
	DrmErrorInvalidFormat    DrmError = 1
	DrmErrorInvalidName      DrmError = 2
)

func (e DrmError) Name() string {
	switch e {
	case DrmErrorAuthenticateFail:
		return "authenticate_fail"
	case DrmErrorInvalidFormat:
		return "invalid_format"
	case DrmErrorInvalidName:
		return "invalid_name"
	default:
		return ""
	}
}

func (e DrmError) Value() string {
	switch e {
	case DrmErrorAuthenticateFail:
		return "0"
	case DrmErrorInvalidFormat:
		return "1"
	case DrmErrorInvalidName:
		return "2"
	default:
		return ""
	}
}

func (e DrmError) String() string {
	return e.Name() + "=" + e.Value()
}

type DrmFormat uint32

// DrmFormat :
const (
	DrmFormatC8          DrmFormat = 0x20203843
	DrmFormatRgb332      DrmFormat = 0x38424752
	DrmFormatBgr233      DrmFormat = 0x38524742
	DrmFormatXrgb4444    DrmFormat = 0x32315258
	DrmFormatXbgr4444    DrmFormat = 0x32314258
	DrmFormatRgbx4444    DrmFormat = 0x32315852
	DrmFormatBgrx4444    DrmFormat = 0x32315842
	DrmFormatArgb4444    DrmFormat = 0x32315241
	DrmFormatAbgr4444    DrmFormat = 0x32314241
	DrmFormatRgba4444    DrmFormat = 0x32314152
	DrmFormatBgra4444    DrmFormat = 0x32314142
	DrmFormatXrgb1555    DrmFormat = 0x35315258
	DrmFormatXbgr1555    DrmFormat = 0x35314258
	DrmFormatRgbx5551    DrmFormat = 0x35315852
	DrmFormatBgrx5551    DrmFormat = 0x35315842
	DrmFormatArgb1555    DrmFormat = 0x35315241
	DrmFormatAbgr1555    DrmFormat = 0x35314241
	DrmFormatRgba5551    DrmFormat = 0x35314152
	DrmFormatBgra5551    DrmFormat = 0x35314142
	DrmFormatRgb565      DrmFormat = 0x36314752
	DrmFormatBgr565      DrmFormat = 0x36314742
	DrmFormatRgb888      DrmFormat = 0x34324752
	DrmFormatBgr888      DrmFormat = 0x34324742
	DrmFormatXrgb8888    DrmFormat = 0x34325258
	DrmFormatXbgr8888    DrmFormat = 0x34324258
	DrmFormatRgbx8888    DrmFormat = 0x34325852
	DrmFormatBgrx8888    DrmFormat = 0x34325842
	DrmFormatArgb8888    DrmFormat = 0x34325241
	DrmFormatAbgr8888    DrmFormat = 0x34324241
	DrmFormatRgba8888    DrmFormat = 0x34324152
	DrmFormatBgra8888    DrmFormat = 0x34324142
	DrmFormatXrgb2101010 DrmFormat = 0x30335258
	DrmFormatXbgr2101010 DrmFormat = 0x30334258
	DrmFormatRgbx1010102 DrmFormat = 0x30335852
	DrmFormatBgrx1010102 DrmFormat = 0x30335842
	DrmFormatArgb2101010 DrmFormat = 0x30335241
	DrmFormatAbgr2101010 DrmFormat = 0x30334241
	DrmFormatRgba1010102 DrmFormat = 0x30334152
	DrmFormatBgra1010102 DrmFormat = 0x30334142
	DrmFormatYuyv        DrmFormat = 0x56595559
	DrmFormatYvyu        DrmFormat = 0x55595659
	DrmFormatUyvy        DrmFormat = 0x59565955
	DrmFormatVyuy        DrmFormat = 0x59555956
	DrmFormatAyuv        DrmFormat = 0x56555941
	DrmFormatXyuv8888    DrmFormat = 0x56555958
	DrmFormatNv12        DrmFormat = 0x3231564e
	DrmFormatNv21        DrmFormat = 0x3132564e
	DrmFormatNv16        DrmFormat = 0x3631564e
	DrmFormatNv61        DrmFormat = 0x3136564e
	DrmFormatYuv410      DrmFormat = 0x39565559
	DrmFormatYvu410      DrmFormat = 0x39555659
	DrmFormatYuv411      DrmFormat = 0x31315559
	DrmFormatYvu411      DrmFormat = 0x31315659
	DrmFormatYuv420      DrmFormat = 0x32315559
	DrmFormatYvu420      DrmFormat = 0x32315659
	DrmFormatYuv422      DrmFormat = 0x36315559
	DrmFormatYvu422      DrmFormat = 0x36315659
	DrmFormatYuv444      DrmFormat = 0x34325559
	DrmFormatYvu444      DrmFormat = 0x34325659
	DrmFormatAbgr16F     DrmFormat = 0x48344241
	DrmFormatXbgr16F     DrmFormat = 0x48344258
)

func (e DrmFormat) Name() string {
	switch e {
	case DrmFormatC8:
		return "c8"
	case DrmFormatRgb332:
		return "rgb332"
	case DrmFormatBgr233:
		return "bgr233"
	case DrmFormatXrgb4444:
		return "xrgb4444"
	case DrmFormatXbgr4444:
		return "xbgr4444"
	case DrmFormatRgbx4444:
		return "rgbx4444"
	case DrmFormatBgrx4444:
		return "bgrx4444"
	case DrmFormatArgb4444:
		return "argb4444"
	case DrmFormatAbgr4444:
		return "abgr4444"
	case DrmFormatRgba4444:
		return "rgba4444"
	case DrmFormatBgra4444:
		return "bgra4444"
	case DrmFormatXrgb1555:
		return "xrgb1555"
	case DrmFormatXbgr1555:
		return "xbgr1555"
	case DrmFormatRgbx5551:
		return "rgbx5551"
	case DrmFormatBgrx5551:
		return "bgrx5551"
	case DrmFormatArgb1555:
		return "argb1555"
	case DrmFormatAbgr1555:
		return "abgr1555"
	case DrmFormatRgba5551:
		return "rgba5551"
	case DrmFormatBgra5551:
		return "bgra5551"
	case DrmFormatRgb565:
		return "rgb565"
	case DrmFormatBgr565:
		return "bgr565"
	case DrmFormatRgb888:
		return "rgb888"
	case DrmFormatBgr888:
		return "bgr888"
	case DrmFormatXrgb8888:
		return "xrgb8888"
	case DrmFormatXbgr8888:
		return "xbgr8888"
	case DrmFormatRgbx8888:
		return "rgbx8888"
	case DrmFormatBgrx8888:
		return "bgrx8888"
	case DrmFormatArgb8888:
		return "argb8888"
	case DrmFormatAbgr8888:
		return "abgr8888"
	case DrmFormatRgba8888:
		return "rgba8888"
	case DrmFormatBgra8888:
		return "bgra8888"
	case DrmFormatXrgb2101010:
		return "xrgb2101010"
	case DrmFormatXbgr2101010:
		return "xbgr2101010"
	case DrmFormatRgbx1010102:
		return "rgbx1010102"
	case DrmFormatBgrx1010102:
		return "bgrx1010102"
	case DrmFormatArgb2101010:
		return "argb2101010"
	case DrmFormatAbgr2101010:
		return "abgr2101010"
	case DrmFormatRgba1010102:
		return "rgba1010102"
	case DrmFormatBgra1010102:
		return "bgra1010102"
	case DrmFormatYuyv:
		return "yuyv"
	case DrmFormatYvyu:
		return "yvyu"
	case DrmFormatUyvy:
		return "uyvy"
	case DrmFormatVyuy:
		return "vyuy"
	case DrmFormatAyuv:
		return "ayuv"
	case DrmFormatXyuv8888:
		return "xyuv8888"
	case DrmFormatNv12:
		return "nv12"
	case DrmFormatNv21:
		return "nv21"
	case DrmFormatNv16:
		return "nv16"
	case DrmFormatNv61:
		return "nv61"
	case DrmFormatYuv410:
		return "yuv410"
	case DrmFormatYvu410:
		return "yvu410"
	case DrmFormatYuv411:
		return "yuv411"
	case DrmFormatYvu411:
		return "yvu411"
	case DrmFormatYuv420:
		return "yuv420"
	case DrmFormatYvu420:
		return "yvu420"
	case DrmFormatYuv422:
		return "yuv422"
	case DrmFormatYvu422:
		return "yvu422"
	case DrmFormatYuv444:
		return "yuv444"
	case DrmFormatYvu444:
		return "yvu444"
	case DrmFormatAbgr16F:
		return "abgr16f"
	case DrmFormatXbgr16F:
		return "xbgr16f"
	default:
		return ""
	}
}

func (e DrmFormat) Value() string {
	switch e {
	case DrmFormatC8:
		return "0x20203843"
	case DrmFormatRgb332:
		return "0x38424752"
	case DrmFormatBgr233:
		return "0x38524742"
	case DrmFormatXrgb4444:
		return "0x32315258"
	case DrmFormatXbgr4444:
		return "0x32314258"
	case DrmFormatRgbx4444:
		return "0x32315852"
	case DrmFormatBgrx4444:
		return "0x32315842"
	case DrmFormatArgb4444:
		return "0x32315241"
	case DrmFormatAbgr4444:
		return "0x32314241"
	case DrmFormatRgba4444:
		return "0x32314152"
	case DrmFormatBgra4444:
		return "0x32314142"
	case DrmFormatXrgb1555:
		return "0x35315258"
	case DrmFormatXbgr1555:
		return "0x35314258"
	case DrmFormatRgbx5551:
		return "0x35315852"
	case DrmFormatBgrx5551:
		return "0x35315842"
	case DrmFormatArgb1555:
		return "0x35315241"
	case DrmFormatAbgr1555:
		return "0x35314241"
	case DrmFormatRgba5551:
		return "0x35314152"
	case DrmFormatBgra5551:
		return "0x35314142"
	case DrmFormatRgb565:
		return "0x36314752"
	case DrmFormatBgr565:
		return "0x36314742"
	case DrmFormatRgb888:
		return "0x34324752"
	case DrmFormatBgr888:
		return "0x34324742"
	case DrmFormatXrgb8888:
		return "0x34325258"
	case DrmFormatXbgr8888:
		return "0x34324258"
	case DrmFormatRgbx8888:
		return "0x34325852"
	case DrmFormatBgrx8888:
		return "0x34325842"
	case DrmFormatArgb8888:
		return "0x34325241"
	case DrmFormatAbgr8888:
		return "0x34324241"
	case DrmFormatRgba8888:
		return "0x34324152"
	case DrmFormatBgra8888:
		return "0x34324142"
	case DrmFormatXrgb2101010:
		return "0x30335258"
	case DrmFormatXbgr2101010:
		return "0x30334258"
	case DrmFormatRgbx1010102:
		return "0x30335852"
	case DrmFormatBgrx1010102:
		return "0x30335842"
	case DrmFormatArgb2101010:
		return "0x30335241"
	case DrmFormatAbgr2101010:
		return "0x30334241"
	case DrmFormatRgba1010102:
		return "0x30334152"
	case DrmFormatBgra1010102:
		return "0x30334142"
	case DrmFormatYuyv:
		return "0x56595559"
	case DrmFormatYvyu:
		return "0x55595659"
	case DrmFormatUyvy:
		return "0x59565955"
	case DrmFormatVyuy:
		return "0x59555956"
	case DrmFormatAyuv:
		return "0x56555941"
	case DrmFormatXyuv8888:
		return "0x56555958"
	case DrmFormatNv12:
		return "0x3231564e"
	case DrmFormatNv21:
		return "0x3132564e"
	case DrmFormatNv16:
		return "0x3631564e"
	case DrmFormatNv61:
		return "0x3136564e"
	case DrmFormatYuv410:
		return "0x39565559"
	case DrmFormatYvu410:
		return "0x39555659"
	case DrmFormatYuv411:
		return "0x31315559"
	case DrmFormatYvu411:
		return "0x31315659"
	case DrmFormatYuv420:
		return "0x32315559"
	case DrmFormatYvu420:
		return "0x32315659"
	case DrmFormatYuv422:
		return "0x36315559"
	case DrmFormatYvu422:
		return "0x36315659"
	case DrmFormatYuv444:
		return "0x34325559"
	case DrmFormatYvu444:
		return "0x34325659"
	case DrmFormatAbgr16F:
		return "0x48344241"
	case DrmFormatXbgr16F:
		return "0x48344258"
	default:
		return ""
	}
}

func (e DrmFormat) String() string {
	return e.Name() + "=" + e.Value()
}

type DrmCapability uint32

// DrmCapability : wl_drm capability bitmask
//
// Bitmask of capabilities.
const (
	// DrmCapabilityPrime : wl_drm prime available
	DrmCapabilityPrime DrmCapability = 1
)

func (e DrmCapability) Name() string {
	switch e {
	case DrmCapabilityPrime:
		return "prime"
	default:
		return ""
	}
}

func (e DrmCapability) Value() string {
	switch e {
	case DrmCapabilityPrime:
		return "1"
	default:
		return ""
	}
}

func (e DrmCapability) String() string {
	return e.Name() + "=" + e.Value()
}

// DrmDeviceEvent :
type DrmDeviceEvent struct {
	Name string
}
type DrmDeviceHandlerFunc func(DrmDeviceEvent)

// SetDeviceHandler : sets handler for DrmDeviceEvent
func (i *Drm) SetDeviceHandler(f DrmDeviceHandlerFunc) {
	i.deviceHandler = f
}

// DrmFormatEvent :
type DrmFormatEvent struct {
	Format uint32
}
type DrmFormatHandlerFunc func(DrmFormatEvent)

// SetFormatHandler : sets handler for DrmFormatEvent
func (i *Drm) SetFormatHandler(f DrmFormatHandlerFunc) {
	i.formatHandler = f
}

// DrmAuthenticatedEvent :
type DrmAuthenticatedEvent struct{}
type DrmAuthenticatedHandlerFunc func(DrmAuthenticatedEvent)

// SetAuthenticatedHandler : sets handler for DrmAuthenticatedEvent
func (i *Drm) SetAuthenticatedHandler(f DrmAuthenticatedHandlerFunc) {
	i.authenticatedHandler = f
}

// DrmCapabilitiesEvent :
type DrmCapabilitiesEvent struct {
	Value uint32
}
type DrmCapabilitiesHandlerFunc func(DrmCapabilitiesEvent)

// SetCapabilitiesHandler : sets handler for DrmCapabilitiesEvent
func (i *Drm) SetCapabilitiesHandler(f DrmCapabilitiesHandlerFunc) {
	i.capabilitiesHandler = f
}

func (i *Drm) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.deviceHandler == nil {
			return
		}
		var e DrmDeviceEvent
		l := 0
		nameLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Name = client.String(data[l : l+nameLen])
		l += nameLen

		i.deviceHandler(e)
	case 1:
		if i.formatHandler == nil {
			return
		}
		var e DrmFormatEvent
		l := 0
		e.Format = client.Uint32(data[l : l+4])
		l += 4

		i.formatHandler(e)
	case 2:
		if i.authenticatedHandler == nil {
			return
		}
		var e DrmAuthenticatedEvent

		i.authenticatedHandler(e)
	case 3:
		if i.capabilitiesHandler == nil {
			return
		}
		var e DrmCapabilitiesEvent
		l := 0
		e.Value = client.Uint32(data[l : l+4])
		l += 4

		i.capabilitiesHandler(e)
	}
}
