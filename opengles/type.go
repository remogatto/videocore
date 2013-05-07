package gles2

//#cgo linux LDFLAGS: -lGLESv2  -lEGL  
//#include <stdlib.h>
//#include <GLES2/gl2.h>
//#include <GLES2/gl2ext.h>
//#include <GLES2/gl2platform.h>
import "C"
import "unsafe"

type (
	Void     unsafe.Pointer
	Enum     uint32
	Bitfield uint32
	Sizei    int32
	Clampf   float32
	Fixed    uintptr
	IntPtr   int
	SizeiPtr int
)

func glString(s string) *C.GLchar {
	return (*C.GLchar)(C.CString(s))
}
func goString(s *C.GLchar) *string {
	gs := C.GoString((*C.char)(s))
	return &gs
}
func goBoolean(n C.GLboolean) *bool {
	b := n == 1
	return &b
}
func glBoolean(n bool) C.GLboolean {
	var b int
	if n == true {
		b = 1
	}
	return C.GLboolean(b)
}
