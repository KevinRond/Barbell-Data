package fqd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go-stats/application/sync"
	"go-stats/domain/result"
)

type Api struct {
	baseUrl          string
	athletesEndpoint string
	resultsEndpoint  string
	meetsEndpoint    string
	rankingsEndpoint string
}

func NewApi(baseUrl string) Api {
	return Api{
		baseUrl:          baseUrl,
		athletesEndpoint: "/athletes",
		resultsEndpoint:  "/results",
		meetsEndpoint:    "/meets",
		rankingsEndpoint: "/rankings",
	}
}

// Client is the adapter implementing application/sync.ResultsFetcher against
// FQD's public API. It is FQD's raw response shape's only consumer — every
// method that returns domain types maps at this boundary.
type Client struct {
	api Api
}

func NewClient(api Api) *Client {
	return &Client{api: api}
}

func (c *Client) getResponseBody(ctx context.Context, endpoint string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.api.baseUrl+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("requesting %s: %w", endpoint, err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("%s returned status %d", endpoint, resp.StatusCode)
	}

	return resp.Body, nil
}

func fetchJSON[T any](ctx context.Context, c *Client, endpoint string) (T, error) {
	var out T
	body, err := c.getResponseBody(ctx, endpoint)
	if err != nil {
		return out, err
	}
	defer body.Close()
	if err := json.NewDecoder(body).Decode(&out); err != nil {
		return out, fmt.Errorf("decoding %s: %w", endpoint, err)
	}
	return out, nil
}

func (c *Client) GetAthleteResults(ctx context.Context, athleteId int) ([]ResultAPIResponse, error) {
	endpoint := c.api.athletesEndpoint + fmt.Sprintf("/%d/results", athleteId)
	return fetchJSON[[]ResultAPIResponse](ctx, c, endpoint)
}

func (c *Client) GetRankings(ctx context.Context, req RankingsRequest) ([]ResultAPIResponse, error) {
	var liftingType string
	if req.BenchOnly {
		liftingType = "bp"
	} else {
		liftingType = "pl"
	}

	var division string
	if req.Equipped {
		division = "eq"
	} else {
		division = "cl"
	}

	params := url.Values{}
	params.Add("year", req.Year)
	params.Add("type", liftingType)
	params.Add("division", division)
	params.Add("gender", req.Gender)
	params.Add("ac", req.AgeCategory)
	params.Add("wc", req.WeightClass) // Safely escapes the "-" character

	endpoint := c.api.rankingsEndpoint + "?" + params.Encode()
	return fetchJSON[[]ResultAPIResponse](ctx, c, endpoint)
}

// defaultLatestQuery is a placeholder for what "latest results" means until
// the real sync scope (which weight classes / genders / years to pull) is
// decided. TODO: iterate the full set of weight classes and genders rather
// than a single hardcoded query.
var defaultLatestQuery = RankingsRequest{
	Year:        "2026",
	Equipped:    false,
	BenchOnly:   false,
	Gender:      "m",
	AgeCategory: "all",
	WeightClass: "-83kg",
}

func (c *Client) FetchLatestRankings(ctx context.Context) ([]result.Result, error) {
	rankings, err := c.GetRankings(ctx, defaultLatestQuery)
	if err != nil {
		return nil, err
	}

	results := make([]result.Result, len(rankings))
	for i, r := range rankings {
		results[i] = mapToResult(r)
	}

	return results, nil
}

func (c *Client) FetchRankings(ctx context.Context, q sync.RankingsQuery) ([]result.Result, error) {
	rankings, err := c.GetRankings(ctx, RankingsRequest{
		Year:        q.Year,
		Equipped:    q.Equipped,
		BenchOnly:   q.BenchOnly,
		Gender:      q.Gender,
		AgeCategory: q.AgeCategory,
		WeightClass: q.WeightClass,
	})
	if err != nil {
		return nil, err
	}
	results := make([]result.Result, len(rankings))
	for i, r := range rankings {
		results[i] = mapToResult(r)
	}
	return results, nil
}

func (c *Client) FetchAthleteResults(ctx context.Context, athleteID result.LifterID) ([]result.Result, error) {
	apiResults, err := c.GetAthleteResults(ctx, int(athleteID))
	if err != nil {
		return nil, err
	}
	results := make([]result.Result, len(apiResults))
	for i, r := range apiResults {
		results[i] = mapToResult(r)
	}
	return results, nil
}

func mapToResult(r ResultAPIResponse) result.Result {
	return result.Result{
		LifterID:     result.LifterID(r.AthleteId),
		MeetID:       result.MeetID(r.MeetId),
		Name:         r.Name,
		Gender:       r.Gender,
		Division:     r.Division,
		AgeCategory:  r.AgeCategory,
		WeightClass:  result.WeightClass(r.WeightClass),
		BodyWeight:   r.BodyWeight,
		IsNovice:     r.IsNovice,
		Squat1:       float64(r.FirstSquat),
		Squat2:       float64(r.SecondSquat),
		Squat3:       float64(r.ThirdSquat),
		BestSquat:    float64(r.BestSquat),
		Bench1:       float64(r.FirstBench),
		Bench2:       float64(r.SecondBench),
		Bench3:       float64(r.ThirdBench),
		BestBench:    float64(r.BestBench),
		Deadlift1:    float64(r.FirstDeadlift),
		Deadlift2:    float64(r.SecondDeadlift),
		Deadlift3:    float64(r.ThirdDeadlift),
		BestDeadlift: float64(r.BestDeadlift),
		Total:        float64(r.Total),
		GL:           float64(r.Gl),
	}
}
