package maps

import "github.com/rowanjacobs/bsa-reader/bsareader/bytes"

type Names struct {
	LocationCount uint32
	Names         []string
}

func readName(bsa []byte) string {
	var curname []byte
	for j := 0; j < 32; j++ {
		if bsa[j] == 0 {
			return string(curname)
		}
		curname = append(curname, bsa[j])
	}
	return string(curname)
}

func ReadNames(bsa []byte) Names {
	lc := bytes.Udword(bsa[0:4])
	var names []string

	for i := 4; i < len(bsa); i += 32 {
		names = append(names, readName(bsa[i:i+32]))
	}

	return Names{
		LocationCount: lc,
		Names:         names,
	}
}
