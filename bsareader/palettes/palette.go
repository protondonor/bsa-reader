package palettes

import "image/color"

func ReadPalette(palBytes []byte) color.Palette {
	pal := color.Palette{}
	alpha := uint8(0)
	for i := 0; i < 256; i++ {
		rgba := color.NRGBA{
			R: palBytes[i*3],
			G: palBytes[i*3+1],
			B: palBytes[i*3+2],
			A: alpha,
		}
		pal = append(pal, rgba)
		if i == 0 {
			alpha = 255
		}
	}
	return pal
}

func ReadCol(colBytes []byte) color.Palette {
	return ReadPalette(colBytes[8:])
}
