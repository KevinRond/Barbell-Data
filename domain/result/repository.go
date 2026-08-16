package result

import "context"

// Repository is the write-side, domain-owned contract for persisting
// Results. Implementations live in infrastructure (e.g. infrastructure/sqlite)
// and must preserve domain invariants on Save.
type Repository interface {
	Save(ctx context.Context, r *Result) error
	ByLifter(ctx context.Context, id LifterID) ([]*Result, error)
}
