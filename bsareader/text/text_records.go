package text

import (
	"github.com/rowanjacobs/bsa-reader/bsareader/bytes"
)

type TextRecordDatabase struct {
	Length  uint16
	Headers []TextRecordHeader
}

type TextRecordHeader struct {
	ID     uint16 // TODO This is the thing we should be selecting them by, not position in the []Headers slice
	Offset uint32
	Text   string
}

func ReadTextRecord(textRsc []byte) TextRecordDatabase {
	trd := TextRecordDatabase{}
	// docs say this is LE but if you read it as LE and look at what's beyond the limit of what you would read as BE
	// you get very strange records with repeated IDs and huge offsets
	trd.Length = bytes.UWordBE(textRsc[0:2])
	//for i := 0; i < 8; i++ {
	//	println(textRsc[i])
	//}

	byteOffset := 2
	var headers []TextRecordHeader

	for i := 0; i < int(trd.Length); i++ {
		trhOffset := byteOffset + i*6
		header := TextRecordHeader{
			// I believe this is LE because the ID of trd.Headers[1] for TEXT.RSC winds up being 1, and so on for a bit
			ID:     bytes.UWord(textRsc[trhOffset : trhOffset+2]),
			// we know this has to be LE because the BE results are ridiculously huge for like all of them
			Offset: bytes.UDword(textRsc[trhOffset+2 : trhOffset+6]),
		}
		if int(header.Offset) >= len(textRsc) {
			// there appear to be a large number of records, past record 1407, that simply don't point anywhere.
			headers = append(headers, header)
			continue
		}

		recordEnd := int(header.Offset)
		for ; recordEnd < len(textRsc); recordEnd++ {
			// TODO: implement subrecords and subrecord separators
			if textRsc[recordEnd] == 0xfe {
				break
			}
		}
		header.Text = string(textRsc[header.Offset:recordEnd])

		headers = append(headers, header)
	}

	trd.Headers = headers

	return trd
}
