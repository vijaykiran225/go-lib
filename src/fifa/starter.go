package main

import (
	"encoding/json"
	parser "fifa/csv_parser"
	"fmt"
	"log"
	"net/http"
	"strconv"

	model "fifa/model"

	"github.com/gorilla/mux"
)

func readAPI(w http.ResponseWriter, r *http.Request) {

	players := []model.Players{}

	vars := mux.Vars(r)
	fmt.Println(vars)
	temp := parser.Parse("E:\\Development\\Sample\\FullData.csv")

	key := vars["attr"]
	target := vars["target"]
	targetVal, _ := strconv.Atoi(target)
	for _, cont := range temp {

		someAttribute, _ := strconv.Atoi(cont[key])
		if someAttribute > targetVal {
			pl := model.Players{}
			pl.Name = cont["Name"]
			pl.Nationality = cont["Nationality"]
			pl.Club = cont["Club"]
			pl.Preffered_Foot = cont["Preffered_Foot"]
			pl.Curve = cont["Curve"]
			pl.Crossing = cont["Crossing"]
			pl.Finishing = cont["Finishing"]
			pl.Speed = cont["Speed"]
			players = append(players, pl)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/call/{attr}/{target}", readAPI)
	fmt.Println("started")
	log.Fatal(http.ListenAndServe(":8081", r))
}
