package es

type disMaxType Object

// DisMax creates a new es.disMaxType object that returns documents matching
// one or more wrapped queries, using the highest score from any matching clause.
//
// Example usage:
//
//	d := es.DisMax(
//	    es.Term("title", "quick"),
//	    es.Term("body", "brown"),
//	)
//
// Parameters:
//   - queries: A variadic list of query clauses to wrap.
//
// Returns:
//
//	An es.disMaxType object containing the specified dis_max query.
func DisMax(queries ...any) disMaxType {
	q := make(Array, 0, len(queries))
	for i := 0; i < len(queries); i++ {
		if field, ok := correctType(queries[i]); ok {
			q = append(q, field)
		}
	}
	return disMaxType{
		"dis_max": Object{
			"queries": q,
		},
	}
}

// TieBreaker sets the "tie_breaker" parameter in an es.disMaxType query.
//
// The tie breaker multiplies the score of non-maximum matching clauses and
// adds them to the final score. Default is 0.0 (only the best clause scores).
//
// Example usage:
//
//	d := es.DisMax(es.Term("title", "quick")).TieBreaker(0.7)
//
// Parameters:
//   - tieBreaker: A float between 0 and 1.
//
// Returns:
//
//	The updated es.disMaxType object with the "tie_breaker" parameter set.
func (d disMaxType) TieBreaker(tieBreaker float64) disMaxType {
	return d.putInTheField("tie_breaker", tieBreaker)
}

// Boost sets the "boost" parameter in an es.disMaxType query.
//
// Example usage:
//
//	d := es.DisMax(es.Term("title", "quick")).Boost(1.2)
//
// Parameters:
//   - boost: A float64 value representing the boost factor.
//
// Returns:
//
//	The updated es.disMaxType object with the "boost" parameter set.
func (d disMaxType) Boost(boost float64) disMaxType {
	return d.putInTheField("boost", boost)
}

// Name sets the "_name" parameter in an es.disMaxType query.
//
// Example usage:
//
//	d := es.DisMax(es.Term("title", "quick")).Name("best_fields")
//
// Parameters:
//   - name: A custom name for the query.
//
// Returns:
//
//	The updated es.disMaxType object with the "_name" parameter set.
func (d disMaxType) Name(name string) disMaxType {
	return d.putInTheField("_name", name)
}

func (d disMaxType) putInTheField(key string, value any) disMaxType {
	return genericPutInTheField(d, "dis_max", key, value)
}
