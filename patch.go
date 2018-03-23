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

func patch(p, p0 uintptr) error {
	return write(p, jmp(p0))
}
