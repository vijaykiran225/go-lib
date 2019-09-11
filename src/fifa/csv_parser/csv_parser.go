package ass

import (
	// "encoding/csv"
	"fmt"
	"log"
	"os"

	csvmap "github.com/recursionpharma/go-csv-map"
)

/*
Parse shit sddf
*/
func Parse(path string) []map[string]string {
	fmt.Println("hellow")
	x, err := os.Open(path)
	if err != nil {
		log.Fatal("screw you")
	}
	c := csvmap.NewReader(x)
	cols, errh := c.ReadHeader()
	if errh != nil {
		log.Fatal("screw you once more")
	}
	c.Columns = cols
	data, errcsv := c.ReadAll()
	if errcsv != nil {
		log.Fatal("screw you again")
	}
	return data
}
