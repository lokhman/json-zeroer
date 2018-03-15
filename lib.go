package json_zeroer

import (
	"reflect"
	"unsafe"
)

//go:linkname _isEmptyValue encoding/json.isEmptyValue
func _isEmptyValue(reflect.Value) bool

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		val := v.Interface()
		if z, ok := val.(Zeroer); ok {
			return z.IsZero()
		}
		return reflect.Zero(v.Type()).Interface() == val
	}
	return false
}

func ptrToFunc(fn func(reflect.Value) bool) unsafe.Pointer {
	return unsafe.Pointer(&fn)
}

func init() {
	err := patch(**(**uintptr)(ptrToFunc(_isEmptyValue)), *(*uintptr)(ptrToFunc(isEmptyValue)))
	if err != nil {
		panic(err)
	}
}
