package main

import (
	"fmt"
	"github.com/remogatto/videocore/egl"
	gl "github.com/remogatto/videocore/opengles"
	"log"
	"runtime/debug"
	"unsafe"
)

var (
	attr = []int32{
		egl.RED_SIZE, 8,
		egl.GREEN_SIZE, 8,
		egl.BLUE_SIZE, 8,
		egl.ALPHA_SIZE, 8,
		egl.SURFACE_TYPE, egl.WINDOW_BIT,
		egl.NONE,
	}
	ctxAttr = []int32{
		egl.CONTEXT_CLIENT_VERSION, 2,
		egl.NONE,
	}
	display                     egl.Display
	config                      egl.Config
	context                     egl.Context
	surface                     egl.Surface
	numConfig                   int32
	Done                        = make(chan bool, 1)
	vid                         int32
	dstRect, srcRect            egl.VCRect
	nativeWindow                egl.EGLDispmanxWindow
	screen_width, screen_height uint32

	attrPos           uint32
	attrCol           uint32
	p                 uint32
	vertexArrayBuffer uint32
	colorArrayBuffer  uint32
	vertices          = [12]float32{
		-0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.0, 1.0,
		0.0, 0.5, 0.0, 1.0,
	}
	colors = [9]float32{
		1.0, 0.0, 0.0,
		0.0, 1.0, 0.0,
		0.0, 0.0, 1.0,
	}
)

func assert(cond bool) {
	if !cond {
		debug.PrintStack()
		panic("Assertion failed!")
	}
}

func check() {
	error := gl.GetError()
	if error != 0 {
		debug.PrintStack()
		panic(fmt.Sprintf("An error occurred! Code: 0x%x", error))
	}
}

func showLog(shader uint32) {
	// Prints the compile log for a shader
	var buf = make([]byte, 1024)
	var s = string(buf)
	gl.GetShaderInfoLog(shader, 1024, nil, &s)
	log.Printf("%d:shader:\n%s\n", shader, s)
}

func showProgramLog(shader uint32) {
	// Prints the information log for a program object
	var logString string
	gl.GetProgramInfoLog(shader, 1024, nil, &logString)
	log.Printf("%d:program:\n%s\n", shader, logString)
}

func initOGL() {
	display = egl.GetDisplay(egl.DEFAULT_DISPLAY)
	if ok := egl.Initialize(display, nil, nil); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.ChooseConfig(display, attr, &config, 1, &numConfig); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(display, config, egl.NATIVE_VISUAL_ID, &vid); !ok {
		egl.LogError(egl.GetError())
	}
	egl.BindAPI(egl.OPENGL_ES_API)
	context = egl.CreateContext(display, config, egl.NO_CONTEXT, &ctxAttr[0])

	screen_width, screen_height = egl.GraphicsGetDisplaySize(0)
	log.Printf("Display size W: %d H: %d\n", screen_width, screen_height)

	dstRect.X = 0
	dstRect.Y = 0
	dstRect.Width = int32(screen_width)
	dstRect.Height = int32(screen_height)

	srcRect.X = 0
	srcRect.Y = 0
	srcRect.Width = int32(screen_width << 16)
	srcRect.Height = int32(screen_height << 16)

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
	check()

	nativeWindow.Element = dispman_element
	nativeWindow.Width = int(screen_width)
	nativeWindow.Height = int(screen_height)
	egl.VCDispmanxUpdateSubmitSync(dispman_update)

	check()

	surface = egl.CreateWindowSurface(
		display,
		config,
		egl.NativeWindowType(unsafe.Pointer(&nativeWindow)),
		nil)
	assert(surface != egl.NO_SURFACE)

	// connect the context to the surface
	result := egl.MakeCurrent(display, surface, surface, context)
	assert(result)

	var val int32
	if ok := egl.QuerySurface(display, &val, egl.WIDTH, surface); !ok {
		egl.LogError(egl.GetError())
	}

	if ok := egl.QuerySurface(display, &val, egl.HEIGHT, surface); !ok {
		egl.LogError(egl.GetError())
	}
	if ok := egl.GetConfigAttrib(display, config, egl.SURFACE_TYPE, &val); !ok {
		egl.LogError(egl.GetError())
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Viewport(0, 0, gl.Sizei(screen_width), gl.Sizei(screen_height))
}

func initShaders() {
	p = Program(FragmentShader(fsh), VertexShader(vsh))
	gl.UseProgram(p)

	attrPos = uint32(gl.GetAttribLocation(p, "pos"))
	attrCol = uint32(gl.GetAttribLocation(p, "color"))

	gl.GenBuffers(1, &vertexArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(vertices))*4, unsafe.Pointer(&vertices[0]), gl.STATIC_DRAW)
	gl.GenBuffers(1, &colorArrayBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorArrayBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, gl.SizeiPtr(len(colors))*4, unsafe.Pointer(&colors[0]), gl.STATIC_DRAW)
	gl.EnableVertexAttribArray(attrPos)
	gl.EnableVertexAttribArray(attrCol)
}

func draw() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexArrayBuffer)
	gl.VertexAttribPointer(attrPos, 4, gl.FLOAT, false, 0, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, colorArrayBuffer)
	gl.VertexAttribPointer(attrCol, 3, gl.FLOAT, false, 0, nil)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.Flush()
        gl.Finish()
	egl.SwapBuffers(display, surface)
}

func cleanup() {
	egl.DestroySurface(display, surface)
	egl.DestroyContext(display, context)
	egl.Terminate(display)
}

func main() {
	egl.BCMHostInit()
	initOGL()
	initShaders()
	defer cleanup()
	for {
		select {
		case <-Done:
			return
		default:
			draw()
		}
	}
}
