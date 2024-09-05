package es

import Operator "github.com/Trendyol/es-query-builder/es/enums/match/operator"

type matchType Object

// Match creates a new matchType object with the specified field and query.
//
// This function initializes a matchType object for a match query, where the key
// is the field name and query is the value to search for in that field. This is used
// to construct queries that match the specified value in the given field.
//
// Example usage:
//
//	m := Match("title", "es-query-builder")
//	// m now contains a matchType object that matches the query "es-query-builder" in the "title" field.
//
// Parameters:
//   - key: A string representing the field name for the match query.
//   - query: The value to be matched in the specified field. The type is generic.
//
// Returns:
//
//	A matchType object containing the specified match query.
func Match[T any](key string, query T) matchType {
	return matchType{
		"match": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

func (m matchType) putInTheField(key string, value any) matchType {
	if match, ok := m["match"].(Object); ok {
		for field := range match {
			if fieldObject, foOk := match[field].(Object); foOk {
				fieldObject[key] = value
			}
		}
	}
	return m
}

// Operator sets the "operator" field in the match query.
//
// This method configures the match query to use a specified operator (e.g., "AND" or "OR")
// for the matching process. It calls putInTheField to add or update the "operator" key
// in the match query object.
//
// Example usage:
//
//	m := Match("title", "es-query-builder").Operator("AND")
//	// m now has an "operator" field set to "AND" in the match query object.
//
// Parameters:
//   - operator: An Operator.Operator value representing the logical operator to be used in the match query.
//
// Returns:
//
//	The updated matchType object with the "operator" field set to the specified value.
func (m matchType) Operator(operator Operator.Operator) matchType {
	return m.putInTheField("operator", operator)
}

// Boost sets the "boost" field in the match query.
//
// This method configures the match query to use a specified boost factor, which influences
// the relevance scoring of the matched documents. It calls putInTheField to add or update
// the "boost" key in the match query object.
//
// Example usage:
//
//	m := Match("title", "es-query-builder").Boost(1.5)
//	// m now has a "boost" field set to 1.5 in the match query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match query.
//
// Returns:
//
//	The updated matchType object with the "boost" field set to the specified value.
func (m matchType) Boost(boost float64) matchType {
	return m.putInTheField("boost", boost)
}
