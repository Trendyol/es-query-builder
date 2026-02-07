package modifier

// Modifier represents the different modifier functions for field_value_factor in function_score queries.
//
// Modifier is a string type used to specify the mathematical function applied to the field value
// before it is used to compute the score in a field_value_factor function.
//
// Example usage:
//
//	var mod Modifier = Log1p
//
//	// Use mod in a field_value_factor configuration
//
// Constants:
//   - None: Do not apply any multiplier to the field value.
//   - Log: Take the common logarithm of the field value.
//   - Log1p: Add 1 to the field value and take the common logarithm.
//   - Log2p: Add 2 to the field value and take the common logarithm.
//   - Ln: Take the natural logarithm of the field value.
//   - Ln1p: Add 1 to the field value and take the natural logarithm.
//   - Ln2p: Add 2 to the field value and take the natural logarithm.
//   - Square: Square the field value (multiply it with itself).
//   - Sqrt: Take the square root of the field value.
//   - Reciprocal: Reciprocate the field value (1/x).
type Modifier string

const (
	// None indicates that no modifier should be applied to the field value.
	None Modifier = "none"

	// Log indicates that the common logarithm should be applied to the field value.
	Log Modifier = "log"

	// Log1p indicates that 1 should be added to the field value before taking the common logarithm.
	Log1p Modifier = "log1p"

	// Log2p indicates that 2 should be added to the field value before taking the common logarithm.
	Log2p Modifier = "log2p"

	// Ln indicates that the natural logarithm should be applied to the field value.
	Ln Modifier = "ln"

	// Ln1p indicates that 1 should be added to the field value before taking the natural logarithm.
	Ln1p Modifier = "ln1p"

	// Ln2p indicates that 2 should be added to the field value before taking the natural logarithm.
	Ln2p Modifier = "ln2p"

	// Square indicates that the field value should be squared.
	Square Modifier = "square"

	// Sqrt indicates that the square root should be applied to the field value.
	Sqrt Modifier = "sqrt"

	// Reciprocal indicates that the reciprocal (1/x) should be applied to the field value.
	Reciprocal Modifier = "reciprocal"
)

func (modifier Modifier) String() string {
	return string(modifier)
}
