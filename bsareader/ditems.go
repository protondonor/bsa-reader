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
