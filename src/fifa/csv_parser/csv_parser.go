package ass

import (
	// "encoding/csv"

	"log"
	"os"

	csvmap "github.com/recursionpharma/go-csv-map"
)

/*
Parse shit sddf
*/
func Parse(path string) []map[string]string {
	x, err := os.Open(path)
	if err != nil {
		log.Fatal("some failure while opening file", err)
	}
	c := csvmap.NewReader(x)
	cols, errh := c.ReadHeader()
	if errh != nil {
		log.Fatal("some failure while reading csv header", err)
	}
	c.Columns = cols
	data, errcsv := c.ReadAll()
	if errcsv != nil {
		log.Fatal("some failure while reading csv file", err)
	}
	return data
}
