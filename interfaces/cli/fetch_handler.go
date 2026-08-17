package cli

import (
	"bufio"
	"context"
	"fmt"
	"go-stats/application/sync"
	"go-stats/domain/result"
	"os"
	"strconv"
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

func (h *FetchHandler) HandleFetchRankings(query sync.RankingsQuery) ([]result.Result, error) {
	return h.client.FetchRankings(context.Background(), query)
}

func (h *FetchHandler) HandleFetchAthleteResults(athleteID result.LifterID) ([]result.Result, error) {
	return h.client.FetchAthleteResults(context.Background(), athleteID)
}

func (h *FetchHandler) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Barbell Data CLI")
		fmt.Println("Query FQD rankings. What do you want to do: ")
		fmt.Println("1. Fetch rankings.")
		fmt.Println("2. Fetch athlete results.")
		fmt.Println("3. Exit.")

		choice, ok, err := promptInt(reader, "Enter your choice: ")
		if err != nil {
			fmt.Println("Error reading choice:", err)
			return
		}
		if !ok {
			fmt.Println("Invalid input! Please enter a valid integer.")
			return
		}

		switch choice {
		case 1:
			year, err := promptString(reader, "Enter Year or \"all\" for all time rankings: ")
			if err != nil {
				fmt.Println("Error reading year:", err)
				return
			}
			isValidYear, err := validateYear(year)
			if err != nil {
				fmt.Println("Error validating year:", err)
				return
			}
			if !isValidYear {
				fmt.Println("Invalid year! Please enter a valid year between 2011 and 2026, or \"all\".")
				return
			}

			equipped, err := promptBool(reader, "Equipped? (y/n): ")
			if err != nil {
				fmt.Println("Error reading equipped status:", err)
				return
			}

			benchOnly, err := promptBool(reader, "Bench Only? (y/n): ")
			if err != nil {
				fmt.Println("Error reading bench only status:", err)
				return
			}

			ageCategory, err := promptAgeCategory(reader)
			if err != nil {
				fmt.Println("Error reading age category:", err)
				return
			}

			gender, err := promptString(reader, "Enter gender (m/f): ")
			if err != nil {
				fmt.Println("Error reading gender:", err)
				return
			}
			isValidGender, err := validateGender(gender)
			if err != nil {
				fmt.Println("Error validating gender:", err)
				return
			}
			if !isValidGender {
				fmt.Println("Invalid gender! Please enter a valid gender.")
				return
			}

			weightClass, err := promptWeightClass(reader, gender)
			if err != nil {
				fmt.Println("Error reading weight class:", err)
				return
			}
			if weightClass == "All weight classes" {
				weightClass = "all"
			}

			var query = sync.RankingsQuery{
				Year:        year,
				Equipped:    equipped,
				BenchOnly:   benchOnly,
				Gender:      gender,
				AgeCategory: ageCategory,
				WeightClass: weightClass,
			}

			results, err := h.HandleFetchRankings(query)
			if err != nil {
				fmt.Println("Error fetching rankings", err)
				return
			}

			limit, ok, err := promptInt(reader, "How many rankings do you want to see? (0 for all): ")
			if err != nil {
				fmt.Println("Error reading limit:", err)
				return
			}
			if !ok {
				fmt.Println("Invalid input! Please enter a valid integer.")
				return
			}
			if limit > 0 && limit < len(results) {
				results = results[:limit]
			}

			fmt.Println("Rankings: ")
			for i, r := range results {
				fmt.Printf("Result %d: %+v\n", i+1, r)
			}

		case 2:
			athleteID, ok, err := promptInt(reader, "Enter athlete ID: ")
			if err != nil {
				fmt.Println("Error reading athlete ID:", err)
				return
			}
			if !ok {
				fmt.Println("Invalid input! Please enter a valid integer.")
				return
			}

			results, err := h.HandleFetchAthleteResults(result.LifterID(athleteID))
			if err != nil {
				fmt.Println("Error fetching athlete results for ID:", err)
				return
			}
			fmt.Println("Athlete Results:")
			for _, r := range results {
				fmt.Printf("Result: %+v\n", r)
			}

		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}

func promptString(reader *bufio.Reader, label string) (string, error) {
	fmt.Print(label)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("reading input: %w", err)
	}
	return strings.TrimSpace(input), nil
}

func promptInt(reader *bufio.Reader, label string) (val int, ok bool, err error) {
	input, err := promptString(reader, label)
	if err != nil {
		return 0, false, err
	}

	val, err = strconv.Atoi(input)
	if err != nil {
		return 0, false, fmt.Errorf("invalid integer %q: %w", input, err)
	}
	return val, true, nil
}

func promptBool(reader *bufio.Reader, prompt string) (bool, error) {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	input = strings.TrimSpace(strings.ToLower(input))

	switch input {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid choice %q: please enter 'y' or 'n'", input)
	}
}

func validateYear(year string) (bool, error) {
	year = strings.TrimSpace(strings.ToLower(year))

	if year == "all" {
		return true, nil
	}

	yearNum, err := strconv.Atoi(year)
	if err != nil {
		return false, fmt.Errorf("invalid year format: %w", err)
	}

	if yearNum >= 2011 && yearNum <= 2026 {
		return true, nil
	}

	return false, fmt.Errorf("year %d out of bounds (2011-2026)", yearNum)
}

func getAgeCategories() []string {
	return []string{
		"all",
		"sj",
		"j",
		"m1",
		"m2",
		"m3",
		"m4",
	}
}

func promptAgeCategory(reader *bufio.Reader) (string, error) {
	categories := getAgeCategories()

	fmt.Println("\nSelect age category:")
	for i, cat := range categories {
		fmt.Printf("%d. %s\n", i+1, cat)
	}

	choice, ok, err := promptInt(reader, "\nEnter choice number: ")
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("could not process integer selection")
	}

	if choice < 1 || choice > len(categories) {
		return "", fmt.Errorf("choice %d is out of bounds (1-%d)", choice, len(categories))
	}

	return categories[choice-1], nil
}

func validateGender(gender string) (bool, error) {
	gender = strings.TrimSpace(strings.ToLower(gender))

	switch gender {
	case "m", "f":
		return true, nil
	default:
		return false, fmt.Errorf("invalid gender %q: must be 'm' (male) or 'f' (female)", gender)
	}
}

func getWeightClasses(gender string) []string {
	gender = strings.TrimSpace(strings.ToLower(gender))

	if gender == "f" || gender == "female" {
		return []string{
			"All weight classes",
			"+84kg",
			"-84kg",
			"-76kg",
			"-69kg",
			"-63kg",
			"-57kg",
			"-52kg",
			"-47kg",
			"-43kg",
		}
	}

	return []string{
		"All weight classes",
		"+120kg",
		"-120kg",
		"-105kg",
		"-93kg",
		"-83kg",
		"-74kg",
		"-66kg",
		"-59kg",
		"-53kg",
	}
}

func promptWeightClass(reader *bufio.Reader, gender string) (string, error) {
	weightClasses := getWeightClasses(gender)

	fmt.Println("\nSelect weight class:")
	for i, class := range weightClasses {
		fmt.Printf("%d. %s\n", i+1, class)
	}

	choice, ok, err := promptInt(reader, "\nEnter choice number: ")
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("could not process integer selection")
	}

	if choice < 1 || choice > len(weightClasses) {
		return "", fmt.Errorf("choice %d is out of bounds (1-%d)", choice, len(weightClasses))
	}

	return weightClasses[choice-1], nil
}
