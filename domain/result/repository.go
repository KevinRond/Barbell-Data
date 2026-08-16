package result

import "context"

type Repository interface {
	Save(ctx context.Context, r *Result) error
	ByLifter(ctx context.Context, id LifterID) ([]*Result, error)
}
