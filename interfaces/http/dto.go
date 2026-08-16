package httpapi

// This file holds request/response JSON structs for OUR API. It is kept
// separate from application/ranking's DTOs so our wire contract can evolve
// independently of internal read-side shapes.

type RankingsQuery struct {
	Lift        string `json:"lift"`
	WeightClass string `json:"weightClass"`
}

type RankingRowResponse struct {
	Rank        int     `json:"rank"`
	LifterID    int     `json:"lifterId"`
	Name        string  `json:"name"`
	WeightClass string  `json:"weightClass"`
	BodyWeight  float64 `json:"bodyWeight"`
	Value       float64 `json:"value"`
}

type ProgressionRowResponse struct {
	MeetID int     `json:"meetId"`
	Date   string  `json:"date"`
	Value  float64 `json:"value"`
}
