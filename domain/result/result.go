package result

type LifterID int

type MeetID int

type WeightClass string

type LiftType string

const (
	Squat    LiftType = "squat"
	Bench    LiftType = "bench"
	Deadlift LiftType = "deadlift"
)

type Result struct {
	LifterID    LifterID
	MeetID      MeetID
	Name        string
	Gender      string
	Division    string
	AgeCategory string
	WeightClass WeightClass
	BodyWeight  float64
	IsNovice    bool

	Squat1       float64
	Squat2       float64
	Squat3       float64
	BestSquat    float64
	Bench1       float64
	Bench2       float64
	Bench3       float64
	BestBench    float64
	Deadlift1    float64
	Deadlift2    float64
	Deadlift3    float64
	BestDeadlift float64

	Total float64
	GL    float64
}
