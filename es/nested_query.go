package es

import ScoreMode "github.com/Trendyol/es-query-builder/es/enums/score-mode"

type nestedType Object

// Nested creates a new es.nestedType object for a nested query.
//
// This function initializes a nested query object with the specified path and query.
// The path represents the field path for the nested query, and the nested query is
// constructed using the es.NewQuery function.
//
// Example usage:
//
//	nestedQuery := es.Nested("comments", es.Bool().Filter(...).MustNot(...))
//	// nestedQuery now contains an es.nestedType object with the specified path and query.
//
// Parameters:
//   - path: A string representing the path for the nested query.
//   - nestedQuery: The query to be applied within the nested query.
//
// Returns:
//
//	an es.nestedType object with the "nested" query and specified path.
func Nested[T any](path string, nestedQuery T) nestedType {
	o := NewQuery(nestedQuery)
	o["path"] = path
	return nestedType{
		"nested": o,
	}
}

// InnerHits sets the "inner_hits" field for the nested query.
//
// This method specifies the inner hits for the nested query, allowing you to control
// the returned inner documents for the nested query. It uses the putInNested method
// to add or update the "inner_hits" key in the nested query object.
//
// Example usage:
//
//	nested := es.Nested("comments", es.Term("text", "example"))
//	nested = nested.InnerHits(Object{"inner": "hits"})
//	// nested now has an "inner_hits" field with the specified Object in the nested query.
//
// Parameters:
//   - innerHits: An es.innerHitsType representing the inner hits configuration for the nested query.
//
// Returns:
//
//	The updated es.nestedType object with the "inner_hits" field set to the specified value.
func (n nestedType) InnerHits(innerHits innerHitsType) nestedType {
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
//	nested := es.Nested("comments", es.Term("text", "example"))
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

// Boost sets the "boost" field for the nested query.
//
// This method applies a boost factor to the nested query, influencing the relevance scoring
// of documents that match the query. It applies the boost to all nested fields in the query.
//
// Example usage:
//
//	n := es.Nested("address").Boost(2.0)
//	// n now has a "boost" field set to 2.0 in the nested query for the "address" field.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the nested query.
//
// Returns:
//
//	The updated es.nestedType object with the "boost" field set to the specified value.
func (n nestedType) Boost(boost float64) nestedType {
	return n.putInNested("boost", boost)
}

// IgnoreUnmapped sets the "ignore_unmapped" field for the nested query.
//
// This method specifies whether to ignore unmapped fields in the nested query.
// If set to true, the query will not fail if a field is not mapped in the index.
//
// Example usage:
//
//	n := es.Nested("address").IgnoreUnmapped(true)
//	// n now has an "ignore_unmapped" field set to true in the nested query for the "address" field.
//
// Parameters:
//   - ignoreUnmapped: A boolean value indicating whether to ignore unmapped fields.
//   - true: Ignore unmapped fields and prevent query failures.
//   - false: Do not ignore unmapped fields (default behavior).
//
// Returns:
//
//	The updated es.nestedType object with the "ignore_unmapped" field set to the specified value.
func (n nestedType) IgnoreUnmapped(ignoreUnmapped bool) nestedType {
	return n.putInNested("ignore_unmapped", ignoreUnmapped)
}

func (n nestedType) putInNested(key string, value any) nestedType {
	if nestedObject, ok := n["nested"].(Object); ok {
		nestedObject[key] = value
	}
	return n
}
