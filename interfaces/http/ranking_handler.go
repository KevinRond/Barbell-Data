package httpapi

import (
	"net/http"

	"go-stats/application/ranking"
)

// RankingHandler is a driver: it calls into application/ranking.Service,
// the same application layer the cron job's SyncService lives beside.
type RankingHandler struct {
	rankingService *ranking.Service
}

func NewRankingHandler(rankingService *ranking.Service) *RankingHandler {
	return &RankingHandler{rankingService: rankingService}
}

// TODO: not wired into a running server yet — application/ranking.Service
// has no read-side store to query until infrastructure/sqlite exists.

func (h *RankingHandler) HandleRankings(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *RankingHandler) HandleProgression(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

func (h *RankingHandler) HandleStats(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
