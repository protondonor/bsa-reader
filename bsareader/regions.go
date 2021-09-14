package bsareader

import (
	"fmt"
	"strconv"
	"strings"
)

var Regions = []string{
	"Alik'r Desert",
	"Dragontail Mountains",
	"Glenpoint Foothills",
	"Daggerfall Bluffs",
	"Yeorth Burrowlands",
	"Dwynnen",
	"Ravennian Forest",
	"Devilrock",
	"Malekna Forest",
	"Isle of Balfiera",
	"Bantha",
	"Dak'fron",
	"Islands in the Western Iliac Bay",
	"Tamarilyn Point",
	"Lainlyn Cliffs",
	"Wrothgarian Mountains",
	"Daggerfall",
	"Glenpoint",
	"Betony",
	"Sentinel",
	"Anticlere",
	"Lainlyn",
	"Wayrest",
	"Gen Tem High Rock village",
	"Gen Rai Hammerfell village",
	"Orsinium Area",
	"Skeffington Wood",
	"Hammerfell bay coast",
	"Hammerfell sea coast",
	"High Rock bay coast",
	"High Rock sea coast",
	"Northmoor",
	"Menevia",
	"Alcaire",
	"Koegria",
	"Bhoriane",
	"Kambria",
	"Phrygias",
	"Urvaius",
	"Ykalon",
	"Daenia",
	"Shalgora",
	"Abibon-Gora",
	"Kairou",
	"Pothago",
	"Myrkwasa",
	"Ayasofya",
	"Tigonus",
	"Kozanset",
	"Satakalaam",
	"Totambu",
	"Mournoth",
	"Ephesus",
	"Santaki",
	"Antiphyllos",
	"Bergama",
	"Gavaudon",
	"Tulune",
	"Glenumbra Moors",
	"Ilessan Hills",
	"Cybiades",
	"Vraseth",
	"Haarvenu",
	"Thrafey",
	"Lyrezi",
	"Montalion",
	"Khulari",
	"Garlythi",
	"Anthotis",
	"Selenu",
}

func ParseRegion(region string) string {
	for i := 0; i < len(Regions); i++ {
		if strings.TrimSpace(strings.ToLower(region)) == strings.ToLower(Regions[i]) {
			return fmt.Sprintf("%03d", i+1)
		}
	}

	// if it's not a region string, maybe it's a region number?
	regionNum, err := strconv.Atoi(region)
	if err != nil {
		return region
	}
	return fmt.Sprintf("%03d", regionNum)
}