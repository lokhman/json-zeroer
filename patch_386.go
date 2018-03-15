package json_zeroer

func jmp(p uintptr) []byte {
	return []byte{
		0xBA,
		byte(p),
		byte(p >> 8),
		byte(p >> 16),
		byte(p >> 24),
		0xFF, 0x22,
	}
}
