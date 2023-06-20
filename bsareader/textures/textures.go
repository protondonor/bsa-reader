package textures

import "github.com/rowanjacobs/bsa-reader/bsareader/bytes"

type TextureFile struct {
	Header         TextureFileHeader
	/*
	There is actually an entire RecordHeader structure.
	(See https://en.uesp.net/wiki/Daggerfall_Mod:Image_formats/Texture )
	However, the only actual value of known import in it is the RecordPosition.
	 */
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
	Scale       int16
	pointer     int32
}

func ReadTextures(data []byte) TextureFile {
	head := TextureFileHeader{
		Count: bytes.Word(data[0:2]),
		Name:  string(data[2:26]),
	}

	var recPtrs []int32
	cursor := 26

	for i := 0; i < int(head.Count); i++ {
		ptr := bytes.Dword(data[cursor+2 : cursor+6])
		recPtrs = append(recPtrs, ptr)

		cursor += 20
	}

	var trs []TextureRecord
	curTr := TextureRecord{}

	for i := 0; i < len(recPtrs); i++ {
		j := recPtrs[i]
		curTr = TextureRecord{
			OffsetX:     bytes.Word(data[j : j+2]),
			OffsetY:     bytes.Word(data[j+2 : j+4]),
			Width:       bytes.Word(data[j+4 : j+6]),
			Height:      bytes.Word(data[j+6 : j+8]),
			Compression: bytes.UWord(data[j+8 : j+10]),
			RecordSize:  bytes.UDword(data[j+10 : j+14]),
			DataOffset:  bytes.UDword(data[j+14 : j+18]),
			IsNormal:    data[j+18] != 0,
			FrameCount:  bytes.UWord(data[j+20 : j+22]),
			Scale:       bytes.Word(data[j+24 : j+26]),
			pointer:     j,
		}
		trs = append(trs, curTr)
	}

	txtFile := TextureFile{Header: head, RecordPointers: recPtrs, TextureRecords: trs}
	return txtFile
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

func (t TextureRecord) Decompress(data []byte) [][]byte {
	// todo: check compression type and act accordingly
	img := [][]byte{}
	cursor := int(t.DataOffset) + int(t.pointer)
	for i := 0; i < int(t.Height); i++ {
		row := data[cursor : cursor+int(t.Width)]
		img = append(img, row)

		cursor += 256 - int(t.Width)
	}

	return img
}

// todo: match images with palettes
// todo: convert images into bitmaps
