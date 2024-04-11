package scoremode

type ScoreMode string

const (
	Avg  ScoreMode = "avg"
	Max  ScoreMode = "max"
	Min  ScoreMode = "min"
	None ScoreMode = "none"
	Sum  ScoreMode = "sum"
)

func (scoreMode ScoreMode) String() string {
	return string(scoreMode)
}
