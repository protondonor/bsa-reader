package bsareader

type PItems struct {
	Items   []LocationExterior
	Offsets []uint32
}

type LocationExterior struct {
	LRE       LocationRecordElement
	Buildings []BuildingData
	Exterior  ExteriorData
}

type BuildingData struct {
}

type ExteriorData struct {
}

func ReadPItems(bsa []byte, count int) PItems {
	var offsets []uint32
	for i := 0; i < count; i += 4 {
		offsets = append(offsets, uint32(dword(bsa[i:i+4])))
	}
	var locexts []LocationExterior
	for i := 0; i < len(offsets); i++ {
		offset := 4*uint32(count) + offsets[i]
		lre := ReadLocationRecordElem(bsa[offset:])
		locexts = append(locexts, LocationExterior{LRE: lre})
	}
	return PItems{Offsets: offsets, Items: locexts}
}
