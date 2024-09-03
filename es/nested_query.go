package es

import ScoreMode "github.com/Trendyol/es-query-builder/es/enums/nested/score-mode"

type nestedType Object

// Nested creates a new nestedType object for a nested query.
//
// This function initializes a nested query object with the specified path and query.
// The path represents the field path for the nested query, and the nested query is
// constructed using the es.NewQuery function.
//
// Example usage:
//
//	nestedQuery := Nested("comments", es.Bool().Filter(...).MustNot(...))
//	// nestedQuery now contains a nestedType object with the specified path and query.
//
// Parameters:
//   - path: A string representing the path for the nested query.
//   - nestedQuery: The query to be applied within the nested query.
//
// Returns:
//
//	A nestedType object with the "nested" query and specified path.
func Nested[T any](path string, nestedQuery T) nestedType {
	o := NewQuery(nestedQuery)
	o["path"] = path
	return nestedType{
		"nested": o,
	}
}

func (n nestedType) putInNested(key string, value any) nestedType {
	if nested, exists := n["nested"]; exists {
		nested.(Object)[key] = value
	}
	return n
}

// InnerHits sets the "inner_hits" field for the nested query.
//
// This method specifies the inner hits for the nested query, allowing you to control
// the returned inner documents for the nested query. It uses the putInNested method
// to add or update the "inner_hits" key in the nested query object.
//
// Example usage:
//
//	nested := Nested("comments", es.Term("text", "example"))
//	nested = nested.InnerHits(Object{"inner": "hits"})
//	// nested now has an "inner_hits" field with the specified Object in the nested query.
//
// Parameters:
//   - innerHits: An es.Object representing the inner hits configuration for the nested query.
//
// Returns:
//
//	The updated nestedType object with the "inner_hits" field set to the specified value.
func (n nestedType) InnerHits(innerHits Object) nestedType {
	return n.putInNested("inner_hits", innerHits)
}

// ScoreMode sets the "score_mode" field for the nested query.
//
// This method configures how the scores for the nested documents should be calculated
// and combined. It uses the putInNested method to add or update the "score_mode"
// key in the nested query object.
//
// Example usage:
//
//	nested := Nested("comments", es.Term("text", "example"))
//	nested = nested.ScoreMode(ScoreMode.Sum)
//	// nested now has a "score_mode" field with the specified ScoreMode value in the nested query.
//
// Parameters:
//   - scoreMode: A ScoreMode.ScoreMode value representing the scoring mode for the nested query.
//
// Returns:
//
//	The updated nestedType object with the "score_mode" field set to the specified value.
func (n nestedType) ScoreMode(scoreMode ScoreMode.ScoreMode) nestedType {
	return n.putInNested("score_mode", scoreMode)
}
