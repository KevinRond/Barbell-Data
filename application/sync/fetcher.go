package sync

import (
	"context"

	"go-stats/domain/result"
)

// ResultsFetcher is the outbound port for pulling results from an external
// federation's API. Implementations (e.g. infrastructure/fqd) own the
// mapping from that federation's raw response shape into domain Results —
// nothing above this port ever sees the raw shape.
type ResultsFetcher interface {
	FetchLatest(ctx context.Context) ([]result.Result, error)
}
