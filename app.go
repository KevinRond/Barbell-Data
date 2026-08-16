package main

import (
	"context"
	"fmt"

	"go-stats/infrastructure/fqd"
)

// TODO: once infrastructure/sqlite exists, wire up:
//   - domain/result.Repository (sqlite-backed)
//   - application/sync.SyncService, started via interfaces/cron.SyncJob
//   - application/ranking.Service, served via interfaces/http.RankingHandler
//
// For now this just demonstrates the FQD adapter, same as the original
// single-file version of this program.
func main() {
	client := fqd.NewClient(fqd.NewApi("https://sheltered-inlet-15640.herokuapp.com/api"))

	results, err := client.FetchLatest(context.Background())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, r := range results {
		fmt.Printf("Result: %+v\n", r)
	}
}
