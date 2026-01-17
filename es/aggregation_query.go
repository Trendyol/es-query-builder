package es

type aggsType Object

// NewAggs creates an aggregation object containing multiple aggregations.
//
// This function is used to combine multiple aggregations under a single "aggs" field.
//
// Example usage:
//
//	aggs := es.NewAggs(
//	    es.Agg("max_price", es.MaxAgg("price")),
//	    es.Agg("avg_price", es.AvgAgg("price")),
//	)
//
// Parameters:
//   - aggs: A variadic list of aggregation definitions.
//
// Returns:
//
//	An Object containing the "aggs" field with the merged aggregation definitions.
func NewAggs(aggs ...aggsType) Object {
	return Object{
		"aggs": reduceAggs(aggs...),
	}
}

// Agg creates a named aggregation entry.
//
// This function allows defining a named aggregation, associating a name with an
// aggregation definition.
//
// Example usage:
//
//	maxPriceAgg := es.Agg("max_price", es.MaxAgg("price"))
//
// Parameters:
//   - name: The name of the aggregation.
//   - agg: The aggregation definition, which must be a map-like type.
//
// Returns:
//
//	An aggsType object representing the named aggregation.
func Agg[T ~map[string]any](name string, agg T) aggsType {
	return aggsType{
		name: agg,
	}
}

// Query adds a query clause to an Elasticsearch request body.
//
// This function modifies the Object to include a "query" field if the provided
// query clause is valid.
//
// Example usage:
//
//	query := es.NewAggs(...).Query(es.MatchQuery("title", "golang"))
//
// Parameters:
//   - queryClause: The query clause to be added.
//
// Returns:
//
//	A modified Object containing the query definition.
func (o Object) Query(queryClause any) Object {
	if field, ok := correctType(queryClause); ok {
		o["query"] = field
	}
	return o
}

func reduceAggs(aggs ...aggsType) Object {
	aggregates := Object{}
	for _, agg := range aggs {
		for key, value := range agg {
			aggregates[key] = value
			break
		}
	}
	return aggregates
}
