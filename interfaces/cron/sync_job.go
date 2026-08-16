package cron

import (
	"context"
	"log"
	"time"

	syncapp "go-stats/application/sync"
)

// SyncJob is a driver: it calls into application/sync.SyncService on a
// schedule, the same application layer the HTTP server sits beside.
//
// TODO: not started from main yet — there's no concrete
// domain/result.Repository to construct a SyncService with until
// infrastructure/sqlite exists.
type SyncJob struct {
	syncService *syncapp.SyncService
	interval    time.Duration
}

func NewSyncJob(syncService *syncapp.SyncService, interval time.Duration) *SyncJob {
	return &SyncJob{syncService: syncService, interval: interval}
}

// Run blocks, triggering SyncDaily on every tick until ctx is cancelled.
func (j *SyncJob) Run(ctx context.Context) {
	ticker := time.NewTicker(j.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := j.syncService.SyncDaily(ctx); err != nil {
				log.Printf("sync job: SyncDaily failed: %v", err)
			}
		}
	}
}
