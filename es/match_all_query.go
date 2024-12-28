package es

type matchAllType Object

// MatchAll creates a new matchAllType object that matches all documents.
//
// This function initializes a matchAllType object that is used to construct queries
// that match all documents. This can be useful for scenarios where you want to ensure
// that all documents are included in the results.
//
// Example usage:
//
//	ma := es.MatchAll()
//	// ma now contains a matchAllType object that matches all documents.
//
// Returns:
//
//	A matchAllType object with a match_all query that matches all documents.
func MatchAll() matchAllType {
	return matchAllType{
		"match_all": Object{},
	}
}

// Boost sets the "boost" field in the match_all query.
//
// This method configures the match_all query to use a specified boost factor, which influences
// the relevance scoring of the matched documents.
//
// Example usage:
//
//	ma := es.MatchAll().Boost(1.5)
//	// ma now has a "boost" field set to 1.5 in the match_all query object.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the match_all query.
//
// Returns:
//
//	The updated matchAllType object with the "boost" field set to the specified value.
func (m matchAllType) Boost(boost float64) matchAllType {
	return m.putInTheField("boost", boost)
}

func (m matchAllType) putInTheField(key string, value any) matchAllType {
	if matchAll, ok := m["match_all"].(Object); ok {
		matchAll[key] = value
	}
	return m
}
