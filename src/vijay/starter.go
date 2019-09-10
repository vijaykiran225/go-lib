package main

import (
	"log"
	"net/http"
	parser "vijay/csv_parser"
)

func readApi(w http.ResponseWriter, r *http.Request) {
	parser.Parse()
}

func main() {
	http.HandleFunc("/call", readApi)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
