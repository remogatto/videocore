package main

import (
	gl "github.com/remogatto/videocore/opengles"
	"log"
)

var (
	vsh = `
        attribute vec4 pos;
        attribute vec3 color;
        varying vec3 v_color;
        void main() {
          gl_Position = pos;
          v_color = color;
        }`

	fsh = `
        varying vec3 v_color;
        void main() {
          gl_FragColor = vec4(v_color, 1.0);
        }`
)

func FragmentShader(s string) uint32 {
	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(shader, 1, &s, nil)
	gl.CompileShader(shader)
	var stat int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &stat)
	if stat == 0 {
		showLog(shader)
		log.Fatalln("Fragment shader compile error ", stat)
	}
	return shader

}

func VertexShader(s string) uint32 {
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(shader, 1, &s, nil)
	gl.CompileShader(shader)
	var stat int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &stat)
	if stat == 0 {
		showLog(shader)
		log.Fatalln("Vertex shader compiler error ", stat)
	}
	return shader
}

func Program(fsh, vsh uint32) uint32 {
	p := gl.CreateProgram()
	gl.AttachShader(p, fsh)
	gl.AttachShader(p, vsh)
	gl.LinkProgram(p)

	var stat int32
	gl.GetProgramiv(p, gl.LINK_STATUS, &stat)

	if stat == 0 {
		var s = make([]byte, 1000)
		var length gl.Sizei
		_log := string(s)
		gl.GetProgramInfoLog(p, 1000, &length, &_log)
		log.Fatalf("Error: linking:\n%s\n", _log)
	}
	return p
}
