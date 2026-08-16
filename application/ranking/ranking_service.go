package ranking

import (
	"context"
	"errors"

	"go-stats/domain/result"
)

// ErrNotImplemented is returned until the read-side query mechanism
// (a direct SQL query path, per the CQRS split — this does not go through
// domain/result.Repository) is wired up in infrastructure.
var ErrNotImplemented = errors.New("ranking: not implemented")

// Service is the read-side, pure aggregation/reporting entry point.
// It intentionally does not depend on domain/result.Repository: rankings,
// progression, and stats are queries, not domain objects.
type Service struct {
	// TODO: a query dependency (e.g. *sql.DB or a query struct) once
	// infrastructure/sqlite exists. Left unset in this skeleton pass.
}

func NewService() *Service {
	return &Service{}
}

// RankingByLift returns the ranked lifters for a given lift and weight class.
func (s *Service) RankingByLift(ctx context.Context, lift result.LiftType, wc result.WeightClass) ([]RankingRow, error) {
	// TODO: query infrastructure/sqlite directly, ordered by Value desc.
	return nil, ErrNotImplemented
}

// Progression returns a lifter's results for a given lift over time.
func (s *Service) Progression(ctx context.Context, id result.LifterID, lift result.LiftType) ([]ProgressionRow, error) {
	// TODO: query infrastructure/sqlite directly, ordered by meet date.
	return nil, ErrNotImplemented
}

// Stats returns an aggregate stat (e.g. mean) for a given lift and weight class.
func (s *Service) Stats(ctx context.Context, lift result.LiftType, wc result.WeightClass) (mean float64, err error) {
	// TODO: query infrastructure/sqlite directly and aggregate.
	return 0, ErrNotImplemented
}
