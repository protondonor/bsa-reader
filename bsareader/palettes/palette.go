package palettes

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Palette [256]Color

func ReadPalette(palBytes []byte) Palette {
	pal := Palette{}
	for i := 0; i < 256; i++ {
		color := Color{
			Red:   palBytes[i*3],
			Green: palBytes[i*3+1],
			Blue:  palBytes[i*3+2],
		}
		pal[i] = color
	}
	return pal
}

func ReadCol(colBytes []byte) Palette {
	pal := Palette{}
	offset := 8
	for i := 0; i < 256; i++ {
		color := Color{
			Red:   colBytes[offset+i*3],
			Green: colBytes[offset+i*3+1],
			Blue:  colBytes[offset+i*3+2],
		}
		pal[i] = color
	}
	return pal
}
