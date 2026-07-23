package validationmethod

// ValidationMethod represents how geo coordinates are validated.
//
// Example usage:
//
//	es.GeoDistance("pin.location", 40.0, -70.0, "12km").
//	    ValidationMethod(ValidationMethod.Strict)
type ValidationMethod string

const (
	// Strict rejects invalid coordinates (default).
	Strict ValidationMethod = "STRICT"

	// IgnoreMalformed accepts invalid coordinates and ignores them.
	IgnoreMalformed ValidationMethod = "IGNORE_MALFORMED"

	// Coerce attempts to correct invalid coordinates.
	Coerce ValidationMethod = "COERCE"
)

func (validationMethod ValidationMethod) String() string {
	return string(validationMethod)
}
