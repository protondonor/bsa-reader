package textures

type TextureFile struct {
	Header         TextureFileHeader
	RecordPointers []int32
	TextureRecords []TextureRecord
}

type TextureFileHeader struct {
	Count int16
	Name  string
}

type TextureRecord struct {
	OffsetX     int16
	OffsetY     int16
	Width       int16
	Height      int16
	Compression uint16
	RecordSize  uint32
	DataOffset  uint32
	IsNormal    bool
	FrameCount  uint16
	XScale      int16
	YScale      int16
	Data        []byte
}

func (t TextureRecord) CompressionType() string {
	switch t.Compression {
	case 0x1108:
		return "RecordRle"
	case 0x0108:
		return "ImageRle"
	case 2:
		return "RleCompressed"
	default:
		return "Uncompressed"
	}
}
