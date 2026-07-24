package es

type matchNoneType Object

// MatchNone creates a new es.matchNoneType object that matches no documents.
//
// This function initializes an es.matchNoneType object that is used to construct queries
// that match no documents. This can be useful for scenarios where you want to explicitly
// return an empty result set.
//
// Example usage:
//
//	mn := es.MatchNone()
//	// mn now contains an es.matchNoneType object that matches no documents.
//
// Returns:
//
//	An es.matchNoneType object with a match_none query that matches no documents.
func MatchNone() matchNoneType {
	return matchNoneType{
		"match_none": Object{},
	}
}

// Boost sets the "boost" field in the match_none query.
//
// This method configures the match_none query to use a specified boost factor, which influences
// the relevance scoring of the matched documents.
//
// Example usage:
//
//	mn := es.MatchNone().Boost(1.5)
//	// mn now has a "boost" field set to 1.5 in the match_none query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match_none query.
//
// Returns:
//
//	The updated es.matchNoneType object with the "boost" field set to the specified value.
func (m matchNoneType) Boost(boost float64) matchNoneType {
	return m.putInTheField("boost", boost)
}

func (m matchNoneType) putInTheField(key string, value any) matchNoneType {
	return genericPutInTheField(m, "match_none", key, value)
}
