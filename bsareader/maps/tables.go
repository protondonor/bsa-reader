package maps

import (
	"fmt"
	"github.com/rowanjacobs/bsa-reader/bsareader/bytes"
)

type Table struct {
	Rows []Row
}

type Row struct {
	MapId         MapId
	LatitudeType  LatitudeType
	LongitudeType LongitudeType
	Flavor        uint8
	Services      uint32
}

type MapId struct {
	LocationExterior uint32 // actually 20 bits
	LocationIndex    uint16 // actually 12 bits
}

// LatitudeType This struct makes up 32 bits as a whole.
type LatitudeType struct {
	Latitude   uint32
	Type       uint16
	Discovered bool
	Hidden     bool
}

// LongitudeType This struct makes up 32 bits as a whole.
type LongitudeType struct {
	Longitude uint32
	Width     uint32
	Height    uint32
}

func (r Row) String() string {
	return fmt.Sprintf(
		"Exterior:%05x\nLatitude:%d\nLongitude:%d\nDiscovered:%t\nServices:%b\n",
		r.MapId.LocationExterior,
		r.LatitudeType.Latitude,
		r.LongitudeType.Longitude,
		r.LatitudeType.Discovered,
		r.Services,
	)
}

func ReadTable(bsa []byte) Table {
	var rows []Row
	for i := 0; i < len(bsa); i += 17 {
		rows = append(rows, makeRow(bsa[i:i+17]))
	}
	return Table{Rows: rows}
}

func makeRow(bsa []byte) Row {
	mapId := makeMapId(bytes.Dword(bsa[0:4]))

	latTypeNum := bytes.Udword(bsa[4:8])
	latType := LatitudeType{
		Latitude:   latTypeNum & 0x1ffffff,
		Type:       uint16((latTypeNum >> 25) & 0x1f),
		Discovered: latTypeNum&0x40000000 != 0,
		Hidden:     latTypeNum&0x80000000 != 0,
	}

	longTypeNum := bytes.Udword(bsa[8:12])
	longType := LongitudeType{
		Longitude: longTypeNum & 0xffffff,
		Height:    (longTypeNum >> 24) & 0xf,
		Width:     longTypeNum >> 28,
	}

	return Row{
		MapId:         mapId,
		LatitudeType:  latType,
		LongitudeType: longType,
		Flavor:        bsa[12],
		Services:      bytes.Udword(bsa[13:17]),
	}
}

func makeMapId(mapIdNum int32) MapId {
	return MapId{
		LocationExterior: uint32(mapIdNum & 0x000fffff),
		LocationIndex:    uint16((uint32(mapIdNum) & 0xfff00000) >> 20),
	}
}
