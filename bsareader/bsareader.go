package bsareader

import (
	"fmt"
	"github.com/rowanjacobs/bsa-reader/bsareader/bytes"
)

// Header https://en.uesp.net/wiki/Daggerfall:BSA_file_formats#BSA_Header
type Header struct {
	RecordCount uint16
	Type        byte
}

type Record struct {
	Name     string
	Size     int32
	Contents []byte
}

const (
	NameRecord   = 1
	NumberRecord = 2
)

// ReadHeader Reads the first 4 bytes of a byte slice as a BSA header.
func ReadHeader(bsa []byte) Header {
	return Header{
		RecordCount: bytes.Word(bsa[0:2]),
		Type:        bsa[3],
	}
}

// GetFooterOffset
// Given the information in a BSA header, returns the offset,
// measured from the end of the file, to the beginning of the
// footer.
func GetFooterOffset(recordCount uint16, bsaType byte) int {
	if bsaType == NameRecord {
		return 18 * int(recordCount)
	}
	return 8 * int(recordCount)
}

// ReadFooter
// Parses a BSA footer into records. If bsaType = 1, it will
// parse NameRecords; otherwise, it will parse NumberRecords.
// https://en.uesp.net/wiki/Daggerfall:BSA_file_formats#BsaFooter
func ReadFooter(footer []byte, bsaType byte) []Record {
	records := []Record{}

	if bsaType == NameRecord {
		for i := 0; i+17 <= len(footer); i += 18 {
			records = append(records, Record{
				Name: string(footer[i : i+12]),
				Size: bytes.Dword(footer[i+14 : i+18]),
			})
		}
	} else {
		for i := 0; i+7 <= len(footer); i += 8 {
			records = append(records, Record{
				Name: fmt.Sprintf("%d", bytes.Word(footer[i:i+2])),
				Size: bytes.Dword(footer[i+4 : i+8]),
			})
		}
	}

	return records
}

func List(bsa []byte) []Record {
	header := ReadHeader(bsa)
	offset := GetFooterOffset(header.RecordCount, header.Type)

	return ReadFooter(bsa[len(bsa)-offset:], header.Type)
}

func Read(bsa []byte) []Record {
	header := ReadHeader(bsa)
	offset := GetFooterOffset(header.RecordCount, header.Type)

	records := ReadFooter(bsa[len(bsa)-offset:], header.Type)

	cursor := 4
	for i := 0; i < int(header.RecordCount); i++ {
		records[i].Contents = bsa[cursor : cursor+int(records[i].Size)]
		cursor += int(records[i].Size)
	}

	return records
}

func ReadRecord(bsa []byte, name string) Record {
	header := ReadHeader(bsa)
	offset := GetFooterOffset(header.RecordCount, header.Type)

	records := ReadFooter(bsa[len(bsa)-offset:], header.Type)

	cursor := 4
	for i := 0; i < int(header.RecordCount); i++ {
		if records[i].Name == name {
			records[i].Contents = bsa[cursor : cursor+int(records[i].Size)]
			return records[i]
		}
		cursor += int(records[i].Size)
	}

	return Record{}
}
