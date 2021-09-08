package bsareader

type Names struct {
	LocationCount uint32
	Names         []string
}

func ReadNames(bsa []byte) Names {
	lc := uint32(dword(bsa[0:4]))
	var names []string

	for i := 4; i < len(bsa); i += 32 {
		var curname []byte
		for j := 0; j < 32; j++ {
			if bsa[i+j] == 0 {
				names = append(names, string(curname))
				break
			}
			curname = append(curname, bsa[i+j])
		}
	}

	return Names{
		LocationCount: lc,
		Names:         names,
	}
}
