package es

type boostingType Object

// Boosting creates a new es.boostingType object that demotes documents matching
// the negative query while keeping documents that match the positive query.
//
// Example usage:
//
//	b := es.Boosting(
//	    es.Match("title", "apple"),
//	    es.Term("category", "outdated"),
//	    0.2,
//	)
//
// Parameters:
//   - positive: The query used to select matching documents.
//   - negative: The query whose matching documents have their score reduced.
//   - negativeBoost: A multiplier applied to scores of documents matching the negative query.
//
// Returns:
//
//	An es.boostingType object containing the specified boosting query, or nil when
//	either query is invalid.
func Boosting(positive, negative any, negativeBoost float64) boostingType {
	positiveField, positiveOk := correctType(positive)
	if !positiveOk {
		return nil
	}
	negativeField, negativeOk := correctType(negative)
	if !negativeOk {
		return nil
	}
	return boostingType{
		"boosting": Object{
			"positive":       positiveField,
			"negative":       negativeField,
			"negative_boost": negativeBoost,
		},
	}
}

// Boost sets the "boost" parameter in an es.boostingType query.
//
// Example usage:
//
//	b := es.Boosting(es.MatchAll(), es.Term("status", "archived"), 0.2).Boost(1.5)
//
// Parameters:
//   - boost: A float64 value representing the boost factor.
//
// Returns:
//
//	The updated es.boostingType object with the "boost" parameter set.
func (b boostingType) Boost(boost float64) boostingType {
	return b.putInTheField("boost", boost)
}

// Name sets the "_name" parameter in an es.boostingType query.
//
// Example usage:
//
//	b := es.Boosting(es.MatchAll(), es.Term("status", "archived"), 0.2).Name("demote_archived")
//
// Parameters:
//   - name: A custom name for the query.
//
// Returns:
//
//	The updated es.boostingType object with the "_name" parameter set.
func (b boostingType) Name(name string) boostingType {
	return b.putInTheField("_name", name)
}

func (b boostingType) putInTheField(key string, value any) boostingType {
	return genericPutInTheField(b, "boosting", key, value)
}
