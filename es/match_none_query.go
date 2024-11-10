package es

import Operator "github.com/Trendyol/es-query-builder/es/enums/match/operator"

type matchNoneType Object

// MatchNone creates a new matchNoneType object with the specified field and query.
//
// This function initializes a matchNoneType object for a match_none query, where the key
// represents the field name and query is the value to be matched. This is used to construct
// queries that explicitly match no documents for the specified value in the given field.
//
// Example usage:
//
//	mn := es.MatchNone("title", "es-query-builder")
//	// mn now contains a matchNoneType object that matches no documents for the "title" field with the query "es-query-builder".
//
// Parameters:
//   - key: A string representing the field name for the match_none query.
//   - query: The value to be used in the match_none query. The type is generic.
//
// Returns:
//
//	A matchNoneType object containing the specified match_none query.
func MatchNone[T any](key string, query T) matchNoneType {
	return matchNoneType{
		"match_none": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

func (m matchNoneType) putInTheField(key string, value any) matchNoneType {
	if matchNone, ok := m["match_none"].(Object); ok {
		for _, fieldObj := range matchNone {
			if fieldObject, foOk := fieldObj.(Object); foOk {
				fieldObject[key] = value
				break
			}
		}
	}
	return m
}

// Operator sets the "operator" field in the match_none query.
//
// This method configures the match_none query to use a specified operator (e.g., "AND" or "OR")
// for the matching process. It calls putInTheField to add or update the "operator" key
// in the match_none query object.
//
// Example usage:
//
//	mn := es.MatchNone("title", "es-query-builder").Operator("AND")
//	// mn now has an "operator" field set to "AND" in the match_none query object.
//
// Parameters:
//   - operator: An Operator.Operator value representing the logical operator to be used in the match_none query.
//
// Returns:
//
//	The updated matchNoneType object with the "operator" field set to the specified value.
func (m matchNoneType) Operator(operator Operator.Operator) matchNoneType {
	return m.putInTheField("operator", operator)
}

// Boost sets the "boost" field in the match_none query.
//
// This method configures the match_none query to use a specified boost factor, which influences
// the relevance scoring of the matched documents. It calls putInTheField to add or update
// the "boost" key in the match_none query object.
//
// Example usage:
//
//	mn := es.MatchNone("title", "es-query-builder").Boost(1.5)
//	// mn now has a "boost" field set to 1.5 in the match_none query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match_none query.
//
// Returns:
//
//	The updated matchNoneType object with the "boost" field set to the specified value.
func (m matchNoneType) Boost(boost float64) matchNoneType {
	return m.putInTheField("boost", boost)
}
