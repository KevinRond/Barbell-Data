package fqd

import (
	"encoding/json"
	"strconv"
)

// This file holds FQD's exact JSON response shapes. Nothing outside this
// package should ever reference these types directly — client.go maps them
// into domain types at the port boundary.

type AthleteAPIResponse struct {
	AthleteId int    `json:"athleteId"`
	LastName  string `json:"last"`
	FirstName string `json:"first"`
}

type RankingsRequest struct {
	Year        string
	Equipped    bool
	BenchOnly   bool
	Gender      string
	AgeCategory string
	WeightClass string
}

// FlexibleFloat64 unmarshals JSON values for lifts since they can be
// represented as strings, numbers or empty values. For example:
//
//	A succesful lift will be represented as 127.5
//	A missed lift will be represented as -127.5
//	An attempt not taken will be "-"
//	The best lift is represented as a string "127.5"
//	If a lifter fails all attempts, the best lift is represented as an empty string ""
type FlexibleFloat64 float64

func (ff *FlexibleFloat64) UnmarshalJSON(data []byte) error {
	rawStr := string(data)
	if rawStr == `""` || rawStr == "null" || rawStr == `"-"` {
		*ff = 0
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if str == "" || str == "-" {
			*ff = 0
			return nil
		}
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		*ff = FlexibleFloat64(val)
		return nil
	}

	var val float64
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*ff = FlexibleFloat64(val)
	return nil
}

type RankingsAPIResponse struct {
	Gender         string          `json:"gender"`
	Name           string          `json:"name"`
	BodyWeight     float64         `json:"bw"`
	Type           string          `json:"type"`
	Division       string          `json:"division"`
	AgeCategory    string          `json:"ac"`
	WeightClass    string          `json:"wc"`
	FirstSquat     FlexibleFloat64 `json:"s1"`
	SecondSquat    FlexibleFloat64 `json:"s2"`
	ThirdSquat     FlexibleFloat64 `json:"s3"`
	BestSquat      FlexibleFloat64 `json:"squat"`
	FirstBench     FlexibleFloat64 `json:"b1"`
	SecondBench    FlexibleFloat64 `json:"b2"`
	ThirdBench     FlexibleFloat64 `json:"b3"`
	BestBench      FlexibleFloat64 `json:"bench"`
	FirstDeadlift  FlexibleFloat64 `json:"d1"`
	SecondDeadlift FlexibleFloat64 `json:"d2"`
	ThirdDeadlift  FlexibleFloat64 `json:"d3"`
	BestDeadlift   FlexibleFloat64 `json:"deadlift"`
	Total          FlexibleFloat64 `json:"total"`
	Gl             FlexibleFloat64 `json:"gl"`
	MeetId         int             `json:"meetId"`
	AthleteId      int             `json:"athleteId"`
	IsNovice       bool            `json:"isNovice"`
}

type ResultAPIResponse struct {
	Gender         string          `json:"gender"`
	Name           string          `json:"name"`
	BodyWeight     float64         `json:"bw"`
	Type           string          `json:"type"`
	Division       string          `json:"division"`
	AgeCategory    string          `json:"ac"`
	WeightClass    string          `json:"wc"`
	FirstSquat     FlexibleFloat64 `json:"s1"`
	SecondSquat    FlexibleFloat64 `json:"s2"`
	ThirdSquat     FlexibleFloat64 `json:"s3"`
	BestSquat      FlexibleFloat64 `json:"squat"`
	FirstBench     FlexibleFloat64 `json:"b1"`
	SecondBench    FlexibleFloat64 `json:"b2"`
	ThirdBench     FlexibleFloat64 `json:"b3"`
	BestBench      FlexibleFloat64 `json:"bench"`
	FirstDeadlift  FlexibleFloat64 `json:"d1"`
	SecondDeadlift FlexibleFloat64 `json:"d2"`
	ThirdDeadlift  FlexibleFloat64 `json:"d3"`
	BestDeadlift   FlexibleFloat64 `json:"deadlift"`
	Total          FlexibleFloat64 `json:"total"`
	Gl             FlexibleFloat64 `json:"gl"`
	MeetId         int             `json:"meetId"`
	AthleteId      int             `json:"athleteId"`
}
