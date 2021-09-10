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
	NameSeed  uint16
	FactionId uint16
	ObjectId  uint32
	Type      uint8
	Quality   uint8
}

type ExteriorData struct {
	Name        string
	MapId       MapId
	Width       uint8
	Height      uint8
	Port        bool
	BlockIndex  []uint8
	BlockNumber []uint8
	BlockChar   []uint8
	DungeonName string
	Encounters  uint8
	BlockCount  uint8
	Blocks      []DungeonBlock
	Services    uint32
}

func ReadPItems(bsa []byte, count int) PItems {
	var offsets []uint32
	for i := 0; i < count; i += 4 {
		offsets = append(offsets, udword(bsa[i:i+4]))
	}
	var locexts []LocationExterior
	for i := 0; i < len(offsets); i++ {
		offset := 4*uint32(count) + offsets[i]
		lre := ReadLocationRecordElem(bsa[offset:])

		bdStart := offset + uint32(lre.Len())
		var buildings []BuildingData
		for j := 0; j < int(lre.Header.BuildingCount); j++ {
			s := bdStart + 26*uint32(j)
			building := BuildingData{
				NameSeed:  word(bsa[s : s+2]),
				FactionId: word(bsa[s+18 : s+20]),
				ObjectId:  udword(bsa[s+20 : s+24]),
				Type:      bsa[s+24],
				Quality:   bsa[s+25],
			}
			buildings = append(buildings, building)
		}

		edStart := bdStart + uint32(lre.Header.BuildingCount)*26
		extData := ExteriorData{
			Name:        readName(bsa[edStart : edStart+32]),
			MapId:       makeMapId(dword(bsa[edStart+32 : edStart+36])),
			Width:       bsa[edStart+40],
			Height:      bsa[edStart+41],
			Port:        bsa[edStart+47] != 0,
			BlockIndex:  bsa[edStart+49 : edStart+113],
			BlockNumber: bsa[edStart+113 : edStart+177],
			BlockChar:   bsa[edStart+177 : edStart+241],
			DungeonName: readName(bsa[edStart+241 : edStart+273]),
			Encounters:  bsa[edStart+273],
			BlockCount:  bsa[edStart+274],
			Services:    udword(bsa[edStart+412 : edStart+416]),
		}

		locexts = append(locexts, LocationExterior{
			LRE:       lre,
			Buildings: buildings,
			Exterior:  extData,
		})
	}
	return PItems{Offsets: offsets, Items: locexts}
}
