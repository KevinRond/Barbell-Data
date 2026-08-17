package main

import (
	"go-stats/infrastructure/fqd"
	"go-stats/interfaces/cli"
)

// TODO: once infrastructure/sqlite exists, wire up:
//   - domain/result.Repository (sqlite-backed)
//   - application/sync.SyncService, started via interfaces/cron.SyncJob
//   - application/ranking.Service, served via interfaces/http.RankingHandler

func main() {
	client := fqd.NewClient(fqd.NewApi("https://sheltered-inlet-15640.herokuapp.com/api"))

	fetchHandler := cli.NewFetchHandler(client)
	fetchHandler.Run()
}
