package es

type rangeType Object

// Range creates a new rangeType object with the specified field.
//
// This function initializes a rangeType object for specifying range queries. The key represents
// the field name, and the rangeType object is used to define the range conditions for that field.
//
// Example usage:
//
//	r := Range("age")
//	// r now contains a rangeType object with the specified field "age" for range queries.
//
// Parameters:
//   - key: A string representing the field name for the range query.
//
// Returns:
//
//	A rangeType object with the specified field ready for defining range conditions.
func Range(key string) rangeType {
	return rangeType{
		key: Object{},
	}
}

// LesserThan sets the "lt" (less than) field for the range query.
//
// This method specifies that the range query should match values that are less than
// the provided value. It removes any existing "lte" (less than or equal to) field to ensure
// that only one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").LesserThan(20)
//	// r now has an "lt" field set to 20 in the range query for the "age" field.
//
// Parameters:
//   - lt: The value that the field should be less than.
//
// Returns:
//
//	The updated rangeType object with the "lt" field set to the specified value.
func (r rangeType) LesserThan(lt any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["lt"] = lt
			delete(rangeObject, "lte")
		}
	}
	return r
}

// LesserThanOrEqual sets the "lte" (less than or equal to) field for the range query.
//
// This method specifies that the range query should match values that are less than or equal
// to the provided value. It removes any existing "lt" (less than) field to ensure that only
// one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").LesserThanOrEqual(20)
//	// r now has an "lte" field set to 20 in the range query for the "age" field.
//
// Parameters:
//   - lte: The value that the field should be less than or equal to.
//
// Returns:
//
//	The updated rangeType object with the "lte" field set to the specified value.
func (r rangeType) LesserThanOrEqual(lte any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["lte"] = lte
			delete(rangeObject, "lt")
		}
	}
	return r
}

// GreaterThan sets the "gt" (greater than) field for the range query.
//
// This method specifies that the range query should match values that are greater than
// the provided value. It removes any existing "gte" (greater than or equal to) field
// to ensure that only one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").GreaterThan(50)
//	// r now has a "gt" field set to 50 in the range query for the "age" field.
//
// Parameters:
//   - gt: The value that the field should be greater than.
//
// Returns:
//
//	The updated rangeType object with the "gt" field set to the specified value.
func (r rangeType) GreaterThan(gt any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["gt"] = gt
			delete(rangeObject, "gte")
		}
	}
	return r
}

// GreaterThanOrEqual sets the "gte" (greater than or equal to) field for the range query.
//
// This method specifies that the range query should match values that are greater than or equal
// to the provided value. It removes any existing "gt" (greater than) field to ensure that only
// one type of range condition is applied.
//
// Example usage:
//
//	r := Range("age").GreaterThanOrEqual(50)
//	// r now has a "gte" field set to 50 in the range query for the "age" field.
//
// Parameters:
//   - gte: The value that the field should be greater than or equal to.
//
// Returns:
//
//	The updated rangeType object with the "gte" field set to the specified value.
func (r rangeType) GreaterThanOrEqual(gte any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["gte"] = gte
			delete(rangeObject, "gt")
		}
	}
	return r
}

// Format sets the "format" field for the range query.
//
// This method specifies a format for the range query values, which can be useful for
// controlling how date or numeric values are interpreted. It applies the format to
// all fields in the range query object.
//
// Example usage:
//
//	r := Range("date").Format("yyyy-MM-dd")
//	// r now has a "format" field set to "yyyy-MM-dd" in the range query for the "date" field.
//
// Parameters:
//   - format: A string representing the format to be applied to the range query values.
//
// Returns:
//
//	The updated rangeType object with the "format" field set to the specified value.
func (r rangeType) Format(format string) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["format"] = format
		}
	}
	return r
}

// Boost sets the "boost" field for the range query.
//
// This method applies a boost factor to the range query, influencing the relevance scoring
// of documents that match the query. It applies the boost to all fields in the range query object.
//
// Example usage:
//
//	r := Range("age").Boost(1.5)
//	// r now has a "boost" field set to 1.5 in the range query for the "age" field.
//
// Parameters:
//   - boost: A float64 value representing the boost factor to be applied to the range query.
//
// Returns:
//
//	The updated rangeType object with the "boost" field set to the specified value.
func (r rangeType) Boost(boost float64) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["boost"] = boost
		}
	}
	return r
}
