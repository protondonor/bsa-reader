package bsareader

import "strconv"

var blockPrefixes = [45]string{
	"TVRN", "GENR", "RESI", "WEAP", "ARMR",
	"ALCH", "BANK", "BOOK", "CLOT", "FURN",
	"GEMS", "LIBR", "PAWN", "TEMP", "PALA",
	"FARM", "DUNG", "CAST", "MANR", "SHRI",
	"RUIN", "SHCK", "GRVE", "FILL", "KRAV",
	"KDRA", "KOWL", "KMOO", "KCAN", "KFLA",
	"KHOR", "KROS", "KWHE", "KSCA", "KHAW",
	"MAGE", "THIE", "DARK", "FIGH", "CUST",
	"WALL", "MARK", "SHIP", "WITC",
}

var blockExtLetters = [12]string{
	"AA", "BA", "AL", "BL", "AM", "BM", "AS", "BS",
	"GA", "GL", "GM", "GS",
}

var templeNumbers = [8]string{
	"A0", "B0", "C0", "D0", "E0", "F0", "G0", "H0",
}

func (e ExteriorData) BlockInfo(index int) string {
	// special case for temples
	if e.BlockIndex[index] == 13 || e.BlockIndex[index] == 14 {
		letters := "AA"
		if e.BlockChar[index] > 7 {
			letters = "GA"
		}

		return "TEMP" + letters + templeNumbers[e.BlockNumber[index]&7]
	}

	// general case
	return blockPrefixes[e.BlockIndex[index]] +
		blockExtLetters[e.BlockChar[index]] +
		strconv.Itoa(int(e.BlockNumber[index]))
}
