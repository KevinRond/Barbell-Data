package cli

import (
	"bufio"
	"context"
	"fmt"
	"go-stats/application/sync"
	"go-stats/domain/result"
	"os"
	"strings"
)

type FetchHandler struct {
	client sync.ResultsFetcher
}

func NewFetchHandler(client sync.ResultsFetcher) *FetchHandler {
	return &FetchHandler{client: client}
}

func (h *FetchHandler) HandleFetchLatestRankings() ([]result.Result, error) {
	return h.client.FetchLatestRankings(context.Background())
}

func (h *FetchHandler) HandleFetch(query sync.RankingsQuery) ([]result.Result, error) {
	return h.client.FetchRankings(context.Background(), query)
}

func (h *FetchHandler) HandleFetchAthleteResults(athleteID result.LifterID) ([]result.Result, error) {
	return h.client.FetchAthleteResults(context.Background(), athleteID)
}

func (h *FetchHandler) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Print("Enter your age: ")
		ageStr, _ := reader.ReadString('\n')
		ageStr = strings.TrimSpace(ageStr)

		fmt.Printf("Hello %s, you are %s years old.\n", name, ageStr)
	}
}
