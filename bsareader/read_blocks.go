package bsareader

import "fmt"

func ReadBlocks(bsa []byte, name string, region string) []string {
	regionDItems := ReadRecord(bsa, fmt.Sprintf("MAPDITEM.%s", region))
	dItems := ReadDItems(regionDItems.Contents)
	return dItems.GetBlocks(name)
}
