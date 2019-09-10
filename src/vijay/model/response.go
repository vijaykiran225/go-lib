package model

type Players struct {
	Name           string `json: Name`
	Nationality    string `json: Nationality`
	Club           string `json: Club`
	Preffered_Foot string `json: Foot`
	Curve          string `json: Curve`
	Crossing       string `json: Crossing`
	Finishing      string `json: Finishing`
	Speed          string `json: Speed`
}
