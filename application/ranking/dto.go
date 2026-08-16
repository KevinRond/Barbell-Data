package ranking

import "go-stats/domain/result"

// RankingRow is one row in a ranking-by-lift result: read-side data,
// not a domain entity, since the read path queries directly and never
// hydrates full Results.
type RankingRow struct {
	Rank        int
	LifterID    result.LifterID
	Name        string
	WeightClass result.WeightClass
	BodyWeight  float64
	Lift        result.LiftType
	Value       float64
}

// ProgressionRow is one point in a lifter's progression over time.
type ProgressionRow struct {
	MeetID result.MeetID
	Date   string
	Lift   result.LiftType
	Value  float64
}
