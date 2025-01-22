package operator

// Operator sets the "operator" field in the matchType object.
//
// This method specifies the logical operator used to combine multiple match conditions.
// It updates the matchType object to include the provided operator for query matching.
//
// Example usage:
//
//	match := es.Match("field", "value").Operator(Operator.And)
//	// match now has the "operator" field set to "and" in the es.matchType object.
//
// Parameters:
//   - operator: An Operator value representing the logical operator to use for combining match conditions.
//     It can be either Operator.Or or Operator.And.
//
// Returns:
//
//	The updated es.matchType object with the "operator" field set to the specified operator.
type Operator string

const (
	// Or indicates that conditions should be combined with a logical OR.
	Or Operator = "or"

	// And indicates that conditions should be combined with a logical AND.
	And Operator = "and"
)

func (operator Operator) String() string {
	return string(operator)
}
