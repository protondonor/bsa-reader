package bsareader

type BsaHeader struct {
	RecordCount uint16
	BsaType     uint16
}

const (
	NameRecord   = 0x100
	NumberRecord = 0x200
)

func word(slice []byte, index int) uint16 {
	return uint16(slice[index])<<8 | uint16(slice[index+1])
}

func ReadHeader(bsa []byte) BsaHeader {
	return BsaHeader{
		RecordCount: word(bsa, 0),
		BsaType:     word(bsa, 2),
	}
}
