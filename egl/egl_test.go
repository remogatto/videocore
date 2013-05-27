package egl

import (
	"unsafe"
	"testing"
	"github.com/remogatto/prettytest"
)

type testSuite struct{ prettytest.Suite }

func (t *testSuite) TestInitialize() {
	var (
		config Config
		numConfig int32
		dstRect, srcRect VCRect
		nativeWindow EGLDispmanxWindow
	)
	attr := []int32{
		RED_SIZE, 8,
		GREEN_SIZE, 8,
		BLUE_SIZE, 8,
		ALPHA_SIZE, 8,
		SURFACE_TYPE, WINDOW_BIT,
		NONE,
	}
	BCMHostInit()
	display := GetDisplay(0)
	t.True(Initialize(display, nil, nil))
	t.True(BindAPI(OPENVG_API))
	t.True(ChooseConfig(display, attr, &config, 1, &numConfig))
	context := CreateContext(display, config, NO_CONTEXT, nil)
	t.True(context != NO_CONTEXT)
	w, h := GraphicsGetDisplaySize(0)
//	t.True(w > 0 && h > 0)

	dstRect.X = 0
	dstRect.Y = 0
	dstRect.Width = int32(w)
	dstRect.Height = int32(h)

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.Width = int32(w << 16)
	srcRect.Height = int32(h << 16)

	dispman_display := VCDispmanxDisplayOpen(0 /* LCD */ )
	dispman_update := VCDispmanxUpdateStart(0)

	dispman_element := VCDispmanxElementAdd(
		dispman_update, 
		dispman_display, 
		0 /*layer */ , 
		&dstRect, 
		0 /*src */,
		&srcRect, 
		DISPMANX_PROTECTION_NONE,
		nil /*alpha */ , 
		nil /*clamp */ ,
		0 /*transform */ );

	nativeWindow.Element = dispman_element
	nativeWindow.Width = int(w)
	nativeWindow.Height = int(h)
	VCDispmanxUpdateSubmitSync(dispman_update)

	surface := CreateWindowSurface(
		display, 
		config, 
		NativeWindowType(unsafe.Pointer(&nativeWindow)), 
		nil)
	t.True(surface != NO_SURFACE)

	// preserve the buffers on swap
	result := SurfaceAttrib(display, surface, SWAP_BEHAVIOR, BUFFER_PRESERVED)
	t.True(result != false)

	// connect the context to the surface
	result = MakeCurrent(display, surface, surface, context)
	t.True(result != false)
}

func TestEGL(t *testing.T) {
	prettytest.Run(
		t,
		new(testSuite),
	)
}




