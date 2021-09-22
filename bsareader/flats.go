package bsareader

import (
	"fmt"
	"strconv"
	"strings"
)

type Flat struct {
	Texture     TextureIndex
	Description string
	Gender      Gender
	FaceIndex   int
}

type TextureIndex struct {
	File  int
	Index int
}

type Gender struct {
	Gender  int
	Obscene bool
}

func (g Gender) String() string {
	gs := ""
	switch g.Gender {
	case 1:
		gs = "Male"
	case 2:
		gs = "Female"
	}

	if g.Obscene {
		gs += " (obscene)"
	}

	return gs
}

func (f Flat) String() string {
	return fmt.Sprintf(
		"%s: TEXTURE.%d[%d], gender %s, face %d",
		f.Description,
		f.Texture.File,
		f.Texture.Index,
		f.Gender,
		f.FaceIndex,
		)
}

func ReadFlats(cfg []byte) []Flat {
	cfgLines := strings.Split(string(cfg), "\r\n")

	var flats []Flat
	curFlat := Flat{}
	j := 0
	for i := 0; i < len(cfgLines); i++ {
		switch j {
		case 6:
			flats = append(flats, curFlat)
			j = 0
			continue
		case 0:
			ti := strings.Split(cfgLines[i], " ")
			if len(ti) < 2 {
				// we're probably at the end of the file, which has a few newlines in a row
				break
			}
			file, err1 := strconv.Atoi(ti[0])
			index, err2 := strconv.Atoi(ti[1])
			if err1 != nil || err2 != nil {
				fmt.Printf("Error reading FLATS.CFG at line %d; expected integer values but got %q\n", i, cfgLines[i])
			}
			curFlat.Texture = TextureIndex{
				File:  file,
				Index: index,
			}
			j++
		case 1:
			curFlat.Description = cfgLines[i]
			j++
		case 2:
			gender := Gender{}
			g := cfgLines[i]
			if cfgLines[i][0] == '?' {
				gender.Obscene = true
				g = g[1:]
			}
			var err error
			gender.Gender, err = strconv.Atoi(g)
			if err != nil {
				fmt.Printf("Error reading FLATS.CFG at line %d; expected integer value but got %q\n", i, cfgLines[i])
			}
			curFlat.Gender = gender
			j++
		case 5:
			fi, err := strconv.Atoi(strings.TrimSpace(cfgLines[i]))
			if err != nil {
				fmt.Printf("Error reading FLATS.CFG at line %d; expected integer value but got %q\n", i, cfgLines[i])
			}
			curFlat.FaceIndex = fi
			j++
		default:
			j++
		}
	}
	return flats
}
