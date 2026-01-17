package es

type constantScoreType Object

// ConstantScore creates a new es.constantScoreType object to apply a constant score to matching documents.
//
// This function initializes an es.constantScoreType object that wraps a given filter query,
// ensuring that all matching documents receive the same score instead of being influenced
// by relevance calculations. The filter parameter represents the query used to determine
// which documents match.
//
// Example usage:
//
//	cs := es.ConstantScore(es.Term("status", "active"))
//	// cs now contains an es.constantScoreType object that applies a constant score to documents
//	// matching the "status" field with the value "active".
//
// Parameters:
//   - filter: An object representing the filter query. It must be a valid query type.
//
// Returns:
//
//	An es.constantScoreType object that applies a constant score to the specified filter query.
func ConstantScore(filter any) constantScoreType {
	field, ok := correctType(filter)
	if !ok {
		return nil
	}
	return constantScoreType{
		"constant_score": Object{
			"filter": field,
		},
	}
}

// Name sets the "_name" parameter in an es.constantScoreType query.
//
// This method assigns a custom name to the constant score query, which can be useful
// for debugging and identifying specific queries in the search response. The assigned
// name appears in the query profile and allows easier tracking of query execution.
//
// Example usage:
//
//	cs := es.ConstantScore(es.Term("status", "active")).Name("status_filter")
//	// cs now includes a "_name" parameter set to "status_filter".
//
// Parameters:
//   - name: A string representing the custom name for the constant score query.
//
// Returns:
//
//	The updated es.constantScoreType object with the "_name" parameter set.
func (cs constantScoreType) Name(name string) constantScoreType {
	return cs.putInTheField("_name", name)
}

// Boost sets the "boost" parameter in an es.constantScoreType query.
//
// This method allows you to specify a boost factor for the constant score query,
// which influences the relevance score of matching documents. A higher boost value
// increases the importance of the query in the overall score, ensuring that documents
// satisfying the filter receive a proportionally higher score.
//
// Example usage:
//
//	cs := es.ConstantScore(es.Term("status", "active")).Boost(2.0)
//	// cs now includes a "boost" parameter set to 2.0.
//
// Parameters:
//   - boost: A float64 value representing the boost factor for the constant
//     score query.
//
// Returns:
//
//	The updated es.constantScoreType object with the "boost" parameter set.
func (cs constantScoreType) Boost(boost float64) constantScoreType {
	return cs.putInTheField("boost", boost)
}

func (cs constantScoreType) putInTheField(key string, value any) constantScoreType {
	return genericPutInTheField(cs, "constant_score", key, value)
}
