package bsareader

import (
	"fmt"
	"os"
)

// GetDaggerfallPath
// IDK if it's cool for me to redistribute MAPS.BSA with this repo
// so this is what we're going to do instead.
func GetDaggerfallPath() string {
	path := os.Getenv("DAGGERFALL_PATH")
	if path != "" {
		return path
	}
	return "/home/rowan/games/abandon/dfall/arena2"
}

func ReadBlocks(bsa []byte, name string, region string) []string {
	regionDItems := ReadRecord(bsa, fmt.Sprintf("MAPDITEM.%s", region))
	dItems := ReadDItems(regionDItems.Contents)
	return dItems.GetBlocks(name)
}
