package ass

import (
	// "encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	csvmap "github.com/recursionpharma/go-csv-map"
)

func Parse() {
	fmt.Println("hellow")
	x, err := os.Open("E:\\Development\\Sample\\FullData.csv")
	if err != nil {
		log.Fatal("screw you")
	}
	c := csvmap.NewReader(x)
	cols, errh := c.ReadHeader()
	if errh != nil {
		log.Fatal("screw you")
	}
	c.Columns = cols
	data, errcsv := c.ReadAll()
	if errcsv != nil {
		log.Fatal("screw you again")
	}
	for _, cont := range data {

		speed, _ := strconv.Atoi(cont["Speed"])
		if speed > 90 {
			fmt.Println(cont["Name"])
		}
	}
}
