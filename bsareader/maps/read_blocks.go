package maps

import (
	"fmt"
	"github.com/rowanjacobs/bsa-reader/bsareader"
)

func ReadBlocks(bsa []byte, name string, region string) []string {
	regionDItems := bsareader.ReadRecord(bsa, fmt.Sprintf("MAPDITEM.%s", region))
	dItems := ReadDItems(regionDItems.Contents)
	return dItems.GetBlocks(name)
}
