package json_zeroer

import "syscall"

func write(p uintptr, data []byte) error {
	page := direct(p & ^(uintptr(syscall.Getpagesize()-1)), syscall.Getpagesize())
	err := syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	if err != nil {
		return err
	}
	copy(direct(p, len(data)), data[:])

	err = syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_EXEC)
	if err != nil {
		return err
	}
	return nil
}
