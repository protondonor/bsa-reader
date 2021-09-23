package maps

import (
	"fmt"
	"strconv"
)

var extBlockPrefixes = [45]string{
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

var dungeonBlockPrefixes = [6]string{
	"N", "W", "L", "S", "B", "M", // L is unused
}

var extBlockLetters = [12]string{
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
	return extBlockPrefixes[e.BlockIndex[index]] +
		extBlockLetters[e.BlockChar[index]] +
		strconv.Itoa(int(e.BlockNumber[index]))
}

func (d DungeonBlock) BlockInfo() string {
	return fmt.Sprintf("%s%07d.RDB",
		dungeonBlockPrefixes[d.BlockIndex],
		d.BlockNumber)
}
