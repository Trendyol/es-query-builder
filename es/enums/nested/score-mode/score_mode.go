package scoremode

// ScoreMode represents the different scoring modes for nested queries.
//
// ScoreMode is a string type used to specify how scores should be calculated and
// combined for nested queries in search queries. It provides various options for
// aggregating scores of nested documents.
//
// Example usage:
//
//	var mode ScoreMode = Sum
//
//	// Use mode in a nested query configuration
//
// Constants:
//   - Avg: Average score of the nested documents.
//   - Max: Maximum score of the nested documents.
//   - Min: Minimum score of the nested documents.
//   - None: No scoring for the nested documents.
//   - Sum: Sum of the scores of the nested documents.
type ScoreMode string

const (
	// Avg indicates that the average score of nested documents should be used.
	Avg ScoreMode = "avg"

	// Max indicates that the maximum score among nested documents should be used.
	Max ScoreMode = "max"

	// Min indicates that the minimum score among nested documents should be used.
	Min ScoreMode = "min"

	// None indicates that no scoring should be applied to nested documents.
	None ScoreMode = "none"

	// Sum indicates that the sum of the scores of nested documents should be used.
	Sum ScoreMode = "sum"
)

func (scoreMode ScoreMode) String() string {
	return string(scoreMode)
}
