package main

import (
	"github.com/remogatto/videocore/egl"
	gl "github.com/remogatto/videocore/opengles"
	"log"
	"fmt"
	"unsafe"
	"runtime/debug"
)

type CubeState struct {
	screen_width  uint32
	screen_height uint32
	// OpenGL|ES objects
	display egl.Display
	surface egl.Surface
	context egl.Context
	verbose  uint
}

var (
	state CubeState
)

func check() {
	error := gl.GetError()
	if error != 0 {
		debug.PrintStack()
		panic(fmt.Sprintf("An error occurred! Code: 0x%x", error))
	}
}

func initOGL(state *CubeState) {
	var (
		config           egl.Config
		numConfig        int32
		dstRect, srcRect egl.VCRect
		nativeWindow     egl.EGLDispmanxWindow
	)
	attr := []int32{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.ALPHA_SIZE, 8,
		egl.SURFACE_TYPE, egl.WINDOW_BIT,
		egl.NONE,
	}
	egl.BCMHostInit()
	state.display = egl.GetDisplay(0)
	egl.Initialize(state.display, nil, nil)
	egl.ChooseConfig(state.display, attr, &config, 1, &numConfig)
	egl.BindAPI(egl.OPENGL_ES_API)

	context := egl.CreateContext(state.display, config, egl.NO_CONTEXT, nil)
	state.screen_width, state.screen_height = egl.GraphicsGetDisplaySize(0)
	log.Printf("Display size W: %d H: %d\n", state.screen_width, state.screen_height)

	dstRect.X = 0
	dstRect.Y = 0
	dstRect.Width = int32(state.screen_width)
	dstRect.Height = int32(state.screen_height)

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.Width = int32(state.screen_width << 16)
	srcRect.Height = int32(state.screen_height << 16)

	dispman_display := egl.VCDispmanxDisplayOpen(0 /* LCD */)
	dispman_update := egl.VCDispmanxUpdateStart(0)

	dispman_element := egl.VCDispmanxElementAdd(
		dispman_update,
		dispman_display,
		0, /*layer */
		&dstRect,
		0, /*src */
		&srcRect,
		egl.DISPMANX_PROTECTION_NONE,
		nil, /*alpha */
		nil, /*clamp */
		0 /*transform */)

	log.Printf("DispmanDisplay 0x%x DispmanUpdate 0x%x DispmanElement 0x%x\n", dispman_display, dispman_update, dispman_element)

	nativeWindow.Element = dispman_element
	nativeWindow.Width = int(state.screen_width)
	nativeWindow.Height = int(state.screen_height)
	egl.VCDispmanxUpdateSubmitSync(dispman_update)

	surface := egl.CreateWindowSurface(
		state.display,
		config,
		egl.NativeWindowType(unsafe.Pointer(&nativeWindow)),
		nil)

	log.Printf("Surface 0x%x\n", surface)

	// preserve the buffers on swap
	egl.SurfaceAttrib(state.display, surface, egl.SWAP_BEHAVIOR, egl.BUFFER_PRESERVED)

	// connect the context to the surface
	result := egl.MakeCurrent(state.display, surface, surface, context)
	log.Printf("MakeCurrent result %d\n", result)
 
	var fb uint32
	var rb uint32
	gl.GenFramebuffers(1, &fb)
	gl.BindFramebuffer(gl.FRAMEBUFFER, fb)

	gl.GenRenderbuffers(1, &rb)
	gl.BindRenderbuffer(gl.RENDERBUFFER, rb)

	gl.RenderbufferStorage(gl.RENDERBUFFER, gl.RGBA4, gl.Sizei(state.screen_width), gl.Sizei(state.screen_height))

	// Set background color and clear buffers
	gl.ClearColor(0.15, 0.25, 0.35, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.BindFramebuffer(gl.FRAMEBUFFER, fb)

	gl.FramebufferRenderbuffer(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.RENDERBUFFER, rb)

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport (0, 0, gl.Sizei(state.screen_width), gl.Sizei(state.screen_height))
}

//==============================================================================

func main() {
	state.verbose = 1
	// Start OGLES
	initOGL(&state)
	for {
		if error := gl.GetError(); error != 0 {
			panic(fmt.Sprintf("An error occured with code: 0x%x\n", error))
		}
		egl.SwapBuffers(state.display, state.surface)
	}
}
