package sync

import (
	"context"

	"go-stats/domain/result"
)

// ResultsFetcher is the outbound port for pulling results from an external
// federation's API. Implementations (e.g. infrastructure/fqd) own the
// mapping from that federation's raw response shape into domain Results —
// nothing above this port ever sees the raw shape.

type RankingsQuery struct {
	Year        string
	Equipped    bool
	BenchOnly   bool
	Gender      string
	AgeCategory string
	WeightClass string
}

type ResultsFetcher interface {
	FetchLatestRankings(ctx context.Context) ([]result.Result, error)
	FetchRankings(ctx context.Context, query RankingsQuery) ([]result.Result, error)
	FetchAthleteResults(ctx context.Context, athleteID result.LifterID) ([]result.Result, error)
}
