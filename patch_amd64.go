package json_zeroer

func jmp(p uintptr) []byte {
	return []byte{
		0x48, 0xBA,
		byte(p),
		byte(p >> 8),
		byte(p >> 16),
		byte(p >> 24),
		byte(p >> 32),
		byte(p >> 40),
		byte(p >> 48),
		byte(p >> 56),
		0xFF, 0x22,
	}
}
