package bytes

func Word(b []byte) int16 {
	return int16(b[1])<<8 | int16(b[0])
}

func Uword(b []byte) uint16 {
	return uint16(b[1])<<8 | uint16(b[0])
}

func Dword(b []byte) int32 {
	return int32(b[0]) + (int32(b[1]) << 8) + (int32(b[2]) << 16) + (int32(b[3]) << 24)
}

func UDword(b []byte) uint32 {
	return uint32(b[0]) + (uint32(b[1]) << 8) + (uint32(b[2]) << 16) + (uint32(b[3]) << 24)
}
