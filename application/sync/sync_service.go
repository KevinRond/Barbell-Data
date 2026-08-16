package sync

import (
	"context"
	"fmt"

	"go-stats/domain/result"
)

// SyncService drives the daily write path: pull the latest results from
// the fetcher and persist them through the domain repository. It is called
// by drivers (the cron job, or manually) — it is not itself a layer.
type SyncService struct {
	repo    result.Repository
	fetcher ResultsFetcher
}

func NewSyncService(repo result.Repository, fetcher ResultsFetcher) *SyncService {
	return &SyncService{repo: repo, fetcher: fetcher}
}

func (s *SyncService) SyncDaily(ctx context.Context) error {
	results, err := s.fetcher.FetchLatest(ctx)
	if err != nil {
		return fmt.Errorf("fetching latest results: %w", err)
	}

	for i := range results {
		if err := s.repo.Save(ctx, &results[i]); err != nil {
			return fmt.Errorf("saving result for lifter %v: %w", results[i].LifterID, err)
		}
	}

	return nil
}
