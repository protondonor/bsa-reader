package bsareader

type Door struct {
	BuildingDataIndex uint16
	// omitting two values with no use:
	// NullValue uint8
	// Mask uint8
	// and omitting two values with unknown purpose:
	// Unknown1 uint8
	// Unknown2 uint8
}

type LocationRecordElement struct {
	Doors        []Door
	ObjectHeader ObjectHeader
	Header       LocationRecordElementHeader
}

// ObjectHeader The real record is 71 bytes long
// but most of those bytes are trash.
type ObjectHeader struct {
	Latitude   int32
	Longitude  int32
	ObjectId   uint32
	IsExterior bool
	ParentId   uint32
}

type LocationRecordElementHeader struct {
	Name          string
	Width         uint8
	Height        uint8
	LocationType  uint8
	BuildingCount uint16
}

func ReadLocationRecordElem(bsa []byte) LocationRecordElement {
	doorCount := uint32(dword(bsa[0:4]))
	var doors []Door
	for i := 0; uint32(i) < doorCount; i++ {
		d := 4 + i*6 // door record start
		door := Door{
			BuildingDataIndex: word(bsa[d : d+2]),
		}
		doors = append(doors, door)
	}

	return LocationRecordElement{
		Doors:        doors,
		ObjectHeader: readObjectHeader(bsa[4:75]),
		Header:       readLREHeader(bsa[75:123]),
	}
}

func readLREHeader(bsa []byte) LocationRecordElementHeader {
	return LocationRecordElementHeader{
		Name:          readName(bsa[0:32]),
		Width:         bsa[32],
		Height:        bsa[33],
		LocationType:  bsa[34],
		BuildingCount: word(bsa[41:43]),
	}
}

func readObjectHeader(bsa []byte) ObjectHeader {
	return ObjectHeader{
		Latitude:   dword(bsa[7:11]),
		Longitude:  dword(bsa[15:19]),
		IsExterior: word(bsa[19:21]) == 0x8000,
		ObjectId:   uint32(dword(bsa[31:35])),
		ParentId:   uint32(dword(bsa[39:43])),
	}
}
