package result

// LifterID identifies an athlete in FQD's system.
type LifterID int

// MeetID identifies a competition in FQD's system.
type MeetID int

// WeightClass is a competition weight class, e.g. "-83kg".
type WeightClass string

// LiftType identifies which of the three competition lifts a ranking or
// progression query is for.
type LiftType string

const (
	Squat    LiftType = "squat"
	Bench    LiftType = "bench"
	Deadlift LiftType = "deadlift"
)

// Result is a lifter's full performance at a single meet: the domain
// entity persisted by Repository and read back via ByLifter.
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
