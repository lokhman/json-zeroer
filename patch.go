package json_zeroer

import (
	"reflect"
	"unsafe"
)

func direct(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

/* func patch(src, dest uintptr) ([]byte, error) {
	jb := jmp(dest)
	raw := direct(src, len(jb))
	bkp := make([]byte, len(raw))
	copy(bkp, raw)
	return bkp, write(src, jb)
} */

func patch(p, p0 uintptr) error {
	return write(p, jmp(p0))
}
