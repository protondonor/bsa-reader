package bsareader

type DItems struct {
	Offsets []DungeonOffset
	Items   []DungeonInterior
}

type DungeonOffset struct {
	Offset       uint32
	DungeonObjId uint32
}

type DungeonInterior struct {
	LRE    LocationRecordElement
	Header DungeonHeader
	Blocks []DungeonBlock
}

type DungeonHeader struct {
	BlockCount uint16
}

type DungeonBlock struct {
	X               uint8
	Z               uint8
	BlockNumber     uint16
	IsStartingBlock bool
	BlockIndex      uint8
}

func (d DItems) GetBlocks(name string) []string {
	var blocks []DungeonBlock
	for i := 0; i < len(d.Items); i++ {
		if d.Items[i].LRE.Header.Name == name {
			blocks = d.Items[i].Blocks
		}
	}

	var result []string
	for i := 0; i < len(blocks); i++ {
		result = append(result, blocks[i].BlockInfo())
	}

	return result
}

func ReadDItems(bsa []byte) DItems {
	count := udword(bsa[0:4])
	var offsets []DungeonOffset
	var items []DungeonInterior
	for i := 0; i < int(count)*8; i += 8 {
		offset := DungeonOffset{
			Offset:       udword(bsa[i+4 : i+8]),
			DungeonObjId: udword(bsa[i : i+4]),
		}
		offsets = append(offsets, offset)

		lreOffset := offset.Offset + count*8 + 4

		lre := ReadLocationRecordElem(bsa[lreOffset:])
		dhOffset := lreOffset + uint32(lre.Len())
		bc := word(bsa[dhOffset+3 : dhOffset+5])

		var blocks []DungeonBlock
		bOffset := dhOffset + 10
		for j := 0; j < int(bc); j += 1 {
			thisBlock := bOffset + 4*uint32(j)
			bn, isb, bi := parseBlockInfo(word(bsa[thisBlock+2 : thisBlock+4]))

			block := DungeonBlock{
				X:               bsa[thisBlock],
				Z:               bsa[thisBlock+1],
				BlockNumber:     bn,
				IsStartingBlock: isb,
				BlockIndex:      bi,
			}

			blocks = append(blocks, block)
		}

		item := DungeonInterior{
			LRE:    lre,
			Header: DungeonHeader{BlockCount: bc},
			Blocks: blocks,
		}
		items = append(items, item)
	}

	return DItems{Offsets: offsets, Items: items}
}

func parseBlockInfo(bnsi uint16) (uint16, bool, uint8) {
	bi := uint8((bnsi & 0xF800) >> 11)
	isb := (bnsi & 0x0400) >> 10
	bn := bnsi & 0x03FF
	return bn, isb != 0, bi
}
