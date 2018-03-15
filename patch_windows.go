package json_zeroer

import (
	"syscall"
	"unsafe"
)

var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

func virtualProtect(lpAddress uintptr, dwSize int, flNewProtect uint32, lpflOldProtect unsafe.Pointer) error {
	ret, _, _ := procVirtualProtect.Call(lpAddress, uintptr(dwSize), uintptr(flNewProtect), uintptr(lpflOldProtect))
	if ret == 0 {
		return syscall.GetLastError()
	}
	return nil
}

func write(p uintptr, data []byte) error {
	f := direct(p, len(data))

	var perms uint32
	err := virtualProtect(p, len(data), 0x40, unsafe.Pointer(&perms))
	if err != nil {
		return err
	}
	copy(f, data[:])

	var tmp uint32
	err = virtualProtect(p, len(data), perms, unsafe.Pointer(&tmp))
	if err != nil {
		return err
	}
	return nil
}
