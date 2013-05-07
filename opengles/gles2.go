package gles2

//#cgo linux LDFLAGS: -lGLESv2  -lEGL  
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/gl2platform.h>
import "C"
import "unsafe"

func ActiveTexture(
	texture Enum) {
	C.glActiveTexture(
		C.GLenum(texture))
}
func AttachShader(
	program uint32, shader uint32) {
	C.glAttachShader(
		C.GLuint(program),
		C.GLuint(shader))
}
func BindAttribLocation(
	program uint32, index uint32, name string) {
	s := glString(name)
	C.glBindAttribLocation(
		C.GLuint(program),
		C.GLuint(index),
		s)

}
func BindBuffer(
	target Enum, buffer uint32) {
	C.glBindBuffer(
		C.GLenum(target),
		C.GLuint(buffer))
}
func BindFramebuffer(
	target Enum, framebuffer uint32) {
	C.glBindFramebuffer(
		C.GLenum(target),
		C.GLuint(framebuffer))
}
func BindRenderbuffer(
	target Enum, renderbuffer uint32) {
	C.glBindRenderbuffer(
		C.GLenum(target),
		C.GLuint(renderbuffer))

}
func BindTexture(
	target Enum, texture uint32) {
	C.glBindTexture(
		C.GLenum(target),
		C.GLuint(texture))

}
func BlendColor(
	red Clampf, green Clampf,
	blue Clampf, alpha Clampf) {
	C.glBlendColor(
		C.GLclampf(red),
		C.GLclampf(green),
		C.GLclampf(blue),
		C.GLclampf(alpha))

}
func BlendEquation(
	mode Enum) {
	C.glBlendEquation(
		C.GLenum(mode))

}
func BlendEquationSeparate(
	modeRGB Enum, modeAlpha Enum) {
	C.glBlendEquationSeparate(
		C.GLenum(modeRGB),
		C.GLenum(modeAlpha))

}
func BlendFunc(
	sfactor Enum, dfactor Enum) {
	C.glBlendFunc(
		C.GLenum(sfactor),
		C.GLenum(dfactor))

}
func BlendFuncSeparate(
	srcRGB Enum, dstRGB Enum,
	srcAlpha Enum, dstAlpha Enum) {
	C.glBlendFuncSeparate(
		C.GLenum(srcRGB),
		C.GLenum(dstRGB),
		C.GLenum(srcAlpha),
		C.GLenum(dstAlpha))

}
func BufferData(
	target Enum, size SizeiPtr,
	data Void, usage Enum) {
	C.glBufferData(
		C.GLenum(target),
		C.GLsizeiptr(size),
		unsafe.Pointer(data),
		C.GLenum(usage))

}
func BufferSubData(
	target Enum, offset IntPtr,
	size SizeiPtr, data Void) {
	C.glBufferSubData(
		C.GLenum(target),
		C.GLintptr(offset),
		C.GLsizeiptr(size),
		unsafe.Pointer(data))

}
func Clear(
	mask Bitfield) {
	C.glClear(
		C.GLbitfield(mask))

}
func ClearColor(
	red Clampf, green Clampf,
	blue Clampf, alpha Clampf) {
	C.glClearColor(
		C.GLclampf(red),
		C.GLclampf(green),
		C.GLclampf(blue),
		C.GLclampf(alpha))

}
func ClearDepthf(
	depth Clampf) {
	C.glClearDepthf(
		C.GLclampf(depth))

}
func ClearStencil(
	s int32) {
	C.glClearStencil(
		C.GLint(s))

}
func ColorMask(
	red bool, green bool,
	blue bool, alpha bool) {
	C.glColorMask(
		glBoolean(red),
		glBoolean(green),
		glBoolean(blue),
		glBoolean(alpha))

}
func CompileShader(
	shader uint32) {
	C.glCompileShader(
		C.GLuint(shader))

}
func CompressedTexImage2D(
	target Enum, level int32, internalformat Enum,
	width Sizei, height Sizei, border int32,
	imageSize Sizei, data Void) {
	C.glCompressedTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLenum(internalformat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLint(border),
		C.GLsizei(imageSize),
		unsafe.Pointer(data))

}
func CompressedTexSubImage2D(
	target Enum, level int32,
	xoffset int32, yoffset int32, width Sizei, height Sizei,
	format Enum, imageSize Sizei, data Void) {
	C.glCompressedTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset),
		C.GLint(yoffset),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLsizei(imageSize),
		unsafe.Pointer(data))

}
func CopyTexImage2D(
	target Enum, level int32, internalformat Enum,
	x int32, y int32, width Sizei, height Sizei, border int32) {
	C.glCopyTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLenum(internalformat),
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLint(border))

}
func CopyTexSubImage2D(
	target Enum, level int32, xoffset int32,
	yoffset int32, x int32, y int32, width Sizei, height Sizei) {
	C.glCopyTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset),
		C.GLint(yoffset),
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height))

}
func CullFace(
	mode Enum) {
	C.glCullFace(
		C.GLenum(mode))

}
func DeleteBuffers(
	n Sizei, buffers *uint32) {
	C.glDeleteBuffers(
		C.GLsizei(n),
		(*C.GLuint)(buffers))

}
func DeleteFramebuffers(
	n Sizei, framebuffers *uint32) {
	C.glDeleteFramebuffers(
		C.GLsizei(n),
		(*C.GLuint)(framebuffers))

}
func DeleteProgram(
	program uint32) {
	C.glDeleteProgram(
		C.GLuint(program))

}
func DeleteRenderbuffers(
	n Sizei, renderbuffers *uint32) {
	C.glDeleteRenderbuffers(
		C.GLsizei(n),
		(*C.GLuint)(renderbuffers))

}
func DeleteShader(
	shader uint32) {
	C.glDeleteShader(
		C.GLuint(shader))

}
func DeleteTextures(
	n Sizei, textures *uint32) {
	C.glDeleteTextures(
		C.GLsizei(n),
		(*C.GLuint)(textures))

}
func DepthFunc(
	func_ Enum) {
	C.glDepthFunc(
		C.GLenum(func_))

}
func DepthMask(
	flag bool) {
	C.glDepthMask(
		glBoolean(flag))

}
func DepthRangef(
	zNear Clampf, zFar Clampf) {
	C.glDepthRangef(
		C.GLclampf(zNear),
		C.GLclampf(zFar))

}
func DetachShader(
	program uint32, shader uint32) {
	C.glDetachShader(
		C.GLuint(program),
		C.GLuint(shader))

}
func Disable(
	cap Enum) {
	C.glDisable(
		C.GLenum(cap))

}
func DisableVertexAttribArray(
	index uint32) {
	C.glDisableVertexAttribArray(
		C.GLuint(index))

}
func DrawArrays(
	mode Enum, first int32, count Sizei) {
	C.glDrawArrays(
		C.GLenum(mode),
		C.GLint(first),
		C.GLsizei(count))

}
func DrawElements(
	mode Enum, count Sizei,
	type_ Enum, indices Void) {
	C.glDrawElements(
		C.GLenum(mode),
		C.GLsizei(count),
		C.GLenum(type_),
		unsafe.Pointer(indices))
}
func Enable(
	cap Enum) {
	C.glEnable(C.GLenum(cap))

}
func EnableVertexAttribArray(
	index uint32) {
	C.glEnableVertexAttribArray(
		C.GLuint(index))
}
func Finish() {
	C.glFinish()
}
func Flush() {
	C.glFlush()
}
func FramebufferRenderbuffer(
	target Enum, attachment Enum,
	renderbuffertarget Enum, renderbuffer uint32) {
	C.glFramebufferRenderbuffer(
		C.GLenum(target),
		C.GLenum(attachment),
		C.GLenum(renderbuffertarget),
		C.GLuint(renderbuffer))
}
func FramebufferTexture2D(
	target Enum, attachment Enum,
	textarget Enum, texture uint32, level int32) {
	C.glFramebufferTexture2D(
		C.GLenum(target),
		C.GLenum(attachment),
		C.GLenum(textarget),
		C.GLuint(texture),
		C.GLint(level))
}
func FrontFace(
	mode Enum) {
	C.glFrontFace(
		C.GLenum(mode))
}
func GenBuffers(
	n Sizei, buffers *uint32) {
	C.glGenBuffers(
		C.GLsizei(n),
		(*C.GLuint)(buffers))
}
func GenerateMipmap(
	target Enum) {

}
func GenFramebuffers(
	n Sizei, framebuffers *uint32) {
	C.glGenFramebuffers(
		C.GLsizei(n),
		(*C.GLuint)(framebuffers))
}
func GenRenderbuffers(
	n Sizei, renderbuffers *uint32) {
	C.glGenRenderbuffers(
		C.GLsizei(n),
		(*C.GLuint)(renderbuffers))
}
func GenTextures(
	n Sizei, textures *uint32) {
	C.glGenTextures(
		C.GLsizei(n),
		(*C.GLuint)(textures))
}
func GetActiveAttrib(
	program uint32, index uint32, bufsize Sizei,
	length *Sizei, size *int32, type_ *Enum, name *string) {
	s := glString(*name)
	C.glGetActiveAttrib(
		C.GLuint(program),
		C.GLuint(index),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		(*C.GLint)(size),
		(*C.GLenum)(type_),
		s)
	name = goString(s)
}
func GetActiveUniform(
	program uint32, index uint32, bufsize Sizei,
	length *Sizei, size *int32, type_ *Enum, name *string) {
	s := glString(*name)
	C.glGetActiveUniform(
		C.GLuint(program),
		C.GLuint(index),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		(*C.GLint)(size),
		(*C.GLenum)(type_),
		s)
	name = goString(s)
}
func GetAttachedShaders(
	program uint32, maxcount Sizei, count *Sizei, shaders *uint32) {
	C.glGetAttachedShaders(
		C.GLuint(program),
		C.GLsizei(maxcount),
		(*C.GLsizei)(count),
		(*C.GLuint)(shaders))
}
func GetBooleanv(
	pname Enum, params *bool) {
	p := glBoolean(*params)
	C.glGetBooleanv(
		C.GLenum(pname),
		&p)
	params = goBoolean(p)

}
func GetBufferParameteriv(
	target Enum, pname Enum, params *int32) {
	C.glGetBufferParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetFloatv(
	pname Enum, params *float32) {
	C.glGetFloatv(
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func GetFramebufferAttachmentParameteriv(
	target Enum, attachment Enum, pname Enum, params *int32) {
	C.glGetFramebufferAttachmentParameteriv(
		C.GLenum(target),
		C.GLenum(attachment),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetIntegerv(
	pname Enum, params *int32) {
	C.glGetIntegerv(
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetProgramiv(
	program uint32, pname Enum, params *int32) {
	C.glGetProgramiv(
		C.GLuint(program),
		C.GLenum(pname),
		(*C.GLint)(params))

}
func GetProgramInfoLog(
	program uint32, bufsize Sizei,
	length *Sizei, infolog *string) {
	s := glString(*infolog)
	C.glGetProgramInfoLog(
		C.GLuint(program),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		s)
	infolog = goString(s)
}
func GetRenderbufferParameteriv(
	target Enum, pname Enum, params *int32) {
	C.glGetRenderbufferParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetShaderiv(
	shader uint32, pname Enum, params *int32) {
	C.glGetShaderiv(
		C.GLuint(shader),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetShaderInfoLog(
	shader uint32, bufsize Sizei,
	length *Sizei, infolog *string) {
	s := glString(*infolog)
	C.glGetShaderInfoLog(
		C.GLuint(shader),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		s)
	infolog = goString(s)
}
func GetShaderPrecisionFormat(
	shadertype Enum, precisiontype Enum,
	range_ *int32, precision *int32) {
	C.glGetShaderPrecisionFormat(
		C.GLenum(shadertype),
		C.GLenum(precisiontype),
		(*C.GLint)(range_),
		(*C.GLint)(precision))
}
func GetShaderSource(
	shader uint32, bufsize Sizei,
	length *Sizei, source *string) {
	s := glString(*source)
	C.glGetShaderSource(
		C.GLuint(shader),
		C.GLsizei(bufsize),
		(*C.GLsizei)(length),
		s)
	source = goString(s)
}
func GetTexParameterfv(
	target Enum, pname Enum, params *float32) {
	C.glGetTexParameterfv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func GetTexParameteriv(
	target Enum, pname Enum, params *int32) {
	C.glGetTexParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetUniformfv(
	program uint32, location int32, params *float32) {
	C.glGetUniformfv(
		C.GLuint(program),
		C.GLint(location),
		(*C.GLfloat)(params))
}
func GetUniformiv(
	program uint32, location int32, params *int32) {
	C.glGetUniformiv(
		C.GLuint(program),
		C.GLint(location),
		(*C.GLint)(params))
}
func GetVertexAttribfv(
	index uint32, pname Enum, params *float32) {
	C.glGetVertexAttribfv(
		C.GLuint(index),
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func GetVertexAttribiv(
	index uint32, pname Enum, params *int32) {
	C.glGetVertexAttribiv(
		C.GLuint(index),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func GetVertexAttribPointerv(
	index uint32, pname Enum, pointer *Void) {
	C.glGetVertexAttribPointerv(
		C.GLuint(index),
		C.GLenum(pname),
		(*unsafe.Pointer)(*pointer))
}
func Hint(
	target Enum, mode Enum) {
	C.glHint(
		C.GLenum(target),
		C.GLenum(mode))
}
func LineWidth(
	width float32) {
	C.glLineWidth(
		C.GLfloat(width))
}
func LinkProgram(
	program uint32) {
	C.glLinkProgram(
		C.GLuint(program))
}
func PixelStorei(
	pname Enum, param int32) {
	C.glPixelStorei(
		C.GLenum(pname),
		C.GLint(param))
}
func PolygonOffset(
	factor float32, units float32) {
	C.glPolygonOffset(
		C.GLfloat(factor),
		C.GLfloat(units))
}
func ReadPixels(
	x int32, y int32, width Sizei, height Sizei,
	format Enum, type_ Enum, pixels Void) {
	C.glReadPixels(
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(type_),
		unsafe.Pointer(pixels))
}
func ReleaseShaderCompiler() {
	C.glReleaseShaderCompiler()
}
func RenderbufferStorage(
	target Enum, internalformat Enum,
	width Sizei, height Sizei) {
	C.glRenderbufferStorage(
		C.GLenum(target),
		C.GLenum(internalformat),
		C.GLsizei(width),
		C.GLsizei(height))
}
func SampleCoverage(
	value Clampf, invert bool) {
	C.glSampleCoverage(
		C.GLclampf(value),
		glBoolean(invert))
}
func Scissor(
	x int32, y int32, width Sizei, height Sizei) {
	C.glScissor(
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height))
}
func ShaderBinary(
	n Sizei, shaders *uint32,
	binaryformat Enum, binary Void, length Sizei) {
	C.glShaderBinary(
		C.GLsizei(n),
		(*C.GLuint)(shaders),
		C.GLenum(binaryformat),
		unsafe.Pointer(binary),
		C.GLsizei(length))
}
func ShaderSource(
	shader uint32, count Sizei,
	string_ *string, length *int32) {
	s := glString(*string_)
	C.glShaderSource(
		C.GLuint(shader),
		C.GLsizei(count),
		&s,
		(*C.GLint)(length))
	string_ = goString(s)
}
func StencilFunc(
	func_ Enum, ref int32, mask uint32) {
	C.glStencilFunc(
		C.GLenum(func_),
		C.GLint(ref),
		C.GLuint(mask))
}
func StencilFuncSeparate(
	face Enum, func_ Enum, ref int32, mask uint32) {
	C.glStencilFuncSeparate(
		C.GLenum(face),
		C.GLenum(func_),
		C.GLint(ref),
		C.GLuint(mask))
}
func StencilMask(
	mask uint32) {
	C.glStencilMask(
		C.GLuint(mask))
}
func StencilMaskSeparate(
	face Enum, mask uint32) {
	C.glStencilMaskSeparate(
		C.GLenum(face),
		C.GLuint(mask))
}
func StencilOp(
	fail Enum, zfail Enum, zpass Enum) {
	C.glStencilOp(
		C.GLenum(fail),
		C.GLenum(zfail),
		C.GLenum(zpass))
}
func StencilOpSeparate(
	face Enum, fail Enum,
	zfail Enum, zpass Enum) {
	C.glStencilOpSeparate(
		C.GLenum(face),
		C.GLenum(fail),
		C.GLenum(zfail),
		C.GLenum(zpass))
}
func TexImage2D(
	target Enum, level int32, internalformat int32,
	width Sizei, height Sizei, border int32, format Enum,
	type_ Enum, pixels Void) {
	C.glTexImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(internalformat),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLint(border),
		C.GLenum(format),
		C.GLenum(type_),
		unsafe.Pointer(pixels))
}
func TexParameterf(
	target Enum, pname Enum, param float32) {
	C.glTexParameterf(
		C.GLenum(target),
		C.GLenum(pname),
		C.GLfloat(param))
}
func TexParameterfv(
	target Enum, pname Enum, params *float32) {
	C.glTexParameterfv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLfloat)(params))
}
func TexParameteri(
	target Enum, pname Enum, param int32) {
	C.glTexParameteri(
		C.GLenum(target),
		C.GLenum(pname),
		C.GLint(param))
}
func TexParameteriv(
	target Enum, pname Enum, params *int32) {
	C.glTexParameteriv(
		C.GLenum(target),
		C.GLenum(pname),
		(*C.GLint)(params))
}
func TexSubImage2D(
	target Enum, level int32, xoffset int32, yoffset int32,
	width Sizei, height Sizei, format Enum, type_ Enum, pixels Void) {
	C.glTexSubImage2D(
		C.GLenum(target),
		C.GLint(level),
		C.GLint(xoffset),
		C.GLint(yoffset),
		C.GLsizei(width),
		C.GLsizei(height),
		C.GLenum(format),
		C.GLenum(type_),
		unsafe.Pointer(pixels))
}
func Uniform1f(
	location int32, x float32) {
	C.glUniform1f(
		C.GLint(location),
		C.GLfloat(x))
}
func Uniform1fv(
	location int32, count Sizei, v *float32) {
	C.glUniform1fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform1i(
	location int32, x int32) {
	C.glUniform1i(
		C.GLint(location),
		C.GLint(x))
}
func Uniform1iv(
	location int32, count Sizei, v *int32) {
	C.glUniform1iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func Uniform2f(
	location int32, x float32, y float32) {
	C.glUniform2f(
		C.GLint(location),
		C.GLfloat(x),
		C.GLfloat(y))
}
func Uniform2fv(
	location int32, count Sizei, v *float32) {
	C.glUniform2fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform2i(
	location int32, x int32, y int32) {
	C.glUniform2i(
		C.GLint(location),
		C.GLint(x),
		C.GLint(y))
}
func Uniform2iv(
	location int32, count Sizei, v *int32) {
	C.glUniform2iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func Uniform3f(
	location int32, x float32, y float32, z float32) {
	C.glUniform3f(
		C.GLint(location),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z))
}
func Uniform3fv(
	location int32, count Sizei, v *float32) {
	C.glUniform3fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform3i(
	location int32, x int32, y int32, z int32) {
	C.glUniform3i(
		C.GLint(location),
		C.GLint(x),
		C.GLint(y),
		C.GLint(z))
}
func Uniform3iv(
	location int32, count Sizei, v *int32) {
	C.glUniform3iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func Uniform4f(
	location int32, x float32, y float32, z float32, w float32) {
	C.glUniform4f(
		C.GLint(location),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z),
		C.GLfloat(w))
}
func Uniform4fv(
	location int32, count Sizei, v *float32) {
	C.glUniform4fv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLfloat)(v))
}
func Uniform4i(
	location int32, x int32, y int32, z int32, w int32) {
	C.glUniform4i(
		C.GLint(location),
		C.GLint(x),
		C.GLint(y),
		C.GLint(z),
		C.GLint(w))
}
func Uniform4iv(
	location int32, count Sizei, v *int32) {
	C.glUniform4iv(
		C.GLint(location),
		C.GLsizei(count),
		(*C.GLint)(v))
}
func UniformMatrix2fv(
	location int32, count Sizei,
	transpose bool, value *float32) {
	C.glUniformMatrix2fv(
		C.GLint(location),
		C.GLsizei(count),
		glBoolean(transpose),
		(*C.GLfloat)(value))
}
func UniformMatrix3fv(
	location int32, count Sizei,
	transpose bool, value *float32) {
	C.glUniformMatrix3fv(
		C.GLint(location),
		C.GLsizei(count),
		glBoolean(transpose),
		(*C.GLfloat)(value))
}
func UniformMatrix4fv(
	location int32, count Sizei,
	transpose bool, value *float32) {
	C.glUniformMatrix4fv(
		C.GLint(location),
		C.GLsizei(count),
		glBoolean(transpose),
		(*C.GLfloat)(value))
}
func UseProgram(
	program uint32) {
	C.glUseProgram(
		C.GLuint(program))
}
func ValidateProgram(
	program uint32) {
	C.glValidateProgram(
		C.GLuint(program))
}
func VertexAttrib1f(
	indx uint32, x float32) {
	C.glVertexAttrib1f(
		C.GLuint(indx),
		C.GLfloat(x))
}
func VertexAttrib1fv(
	indx uint32, values *float32) {
	C.glVertexAttrib1fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttrib2f(
	indx uint32, x float32, y float32) {
	C.glVertexAttrib2f(
		C.GLuint(indx),
		C.GLfloat(x),
		C.GLfloat(y))
}
func VertexAttrib2fv(
	indx uint32, values *float32) {
	C.glVertexAttrib2fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttrib3f(
	indx uint32, x float32, y float32, z float32) {
	C.glVertexAttrib3f(
		C.GLuint(indx),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z))
}
func VertexAttrib3fv(
	indx uint32, values *float32) {
	C.glVertexAttrib3fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttrib4f(
	indx uint32, x float32, y float32, z float32, w float32) {
	C.glVertexAttrib4f(
		C.GLuint(indx),
		C.GLfloat(x),
		C.GLfloat(y),
		C.GLfloat(z),
		C.GLfloat(w))
}
func VertexAttrib4fv(
	indx uint32, values *float32) {
	C.glVertexAttrib4fv(
		C.GLuint(indx),
		(*C.GLfloat)(values))
}
func VertexAttribPointer(
	indx uint32, size int32, type_ Enum,
	normalized bool, stride Sizei, ptr Void) {
	C.glVertexAttribPointer(
		C.GLuint(indx),
		C.GLint(size),
		C.GLenum(type_),
		glBoolean(normalized),
		C.GLsizei(stride),
		unsafe.Pointer(ptr))
}
func Viewport(
	x int32, y int32, width Sizei, height Sizei) {
	C.glViewport(
		C.GLint(x),
		C.GLint(y),
		C.GLsizei(width),
		C.GLsizei(height))
}
func IsBuffer(buffer uint32) bool {
	return *goBoolean(C.glIsBuffer(
		C.GLuint(buffer)))
}
func IsEnabled(cap Enum) bool {
	return *goBoolean(C.glIsEnabled(
		C.GLenum(cap)))
}
func IsFramebuffer(framebuffer uint32) bool {
	return *goBoolean(C.glIsFramebuffer(
		C.GLuint(framebuffer)))
}
func IsProgram(program uint32) bool {
	return *goBoolean(C.glIsProgram(
		C.GLuint(program)))
}
func IsRenderbuffer(renderbuffer uint32) bool {
	return *goBoolean(C.glIsRenderbuffer(
		C.GLuint(renderbuffer)))
}
func IsShader(shader uint32) bool {
	return *goBoolean(C.glIsShader(
		C.GLuint(shader)))
}
func IsTexture(texture uint32) bool {
	return *goBoolean(C.glIsTexture(
		C.GLuint(texture)))
}
func GetError() Enum {
	return Enum(C.glGetError())
}
func CheckFramebufferStatus(target Enum) Enum {
	return Enum(C.glCheckFramebufferStatus(
		C.GLenum(target)))
}
func CreateProgram() uint32 {
	return uint32(C.glCreateProgram())
}
func CreateShader(type_ Enum) uint32 {
	return uint32(C.glCreateShader(
		C.GLenum(type_)))
}
func GetAttribLocation(program uint32, name *string) uintptr {
	s := glString(*name)
	return uintptr(C.glGetAttribLocation(
		C.GLuint(program),
		s))
}
func GetUniformLocation(program uint32, name string) uintptr {
	s := glString(name)
	return uintptr(C.glGetUniformLocation(
		C.GLuint(program),
		s))
}
func GetString(name Enum) string {
	return string(byte(*C.glGetString(
		C.GLenum(name))))
}
