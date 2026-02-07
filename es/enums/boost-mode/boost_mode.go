package boostmode

// BoostMode represents the different boost modes for function_score queries.
//
// BoostMode is a string type used to specify how the computed score from the
// function_score query is combined with the query score. It provides various
// options for combining the function score with the original query score.
//
// Example usage:
//
//	var mode BoostMode = Multiply
//
//	// Use mode in a function_score query configuration
//
// Constants:
//   - Multiply: Multiply the query score with the function score (default).
//   - Replace: Replace the query score with the function score.
//   - Sum: Add the query score and the function score.
//   - Avg: Average the query score and the function score.
//   - Max: Use the maximum of the query score and the function score.
//   - Min: Use the minimum of the query score and the function score.
type BoostMode string

const (
	// Multiply indicates that the query score should be multiplied with the function score.
	Multiply BoostMode = "multiply"

	// Replace indicates that the query score should be replaced with the function score.
	Replace BoostMode = "replace"

	// Sum indicates that the query score and the function score should be added.
	Sum BoostMode = "sum"

	// Avg indicates that the query score and the function score should be averaged.
	Avg BoostMode = "avg"

	// Max indicates that the maximum of the query score and the function score should be used.
	Max BoostMode = "max"

	// Min indicates that the minimum of the query score and the function score should be used.
	Min BoostMode = "min"
)

func (boostMode BoostMode) String() string {
	return string(boostMode)
}
