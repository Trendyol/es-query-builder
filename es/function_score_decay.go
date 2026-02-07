package es

import MultiValuesMode "github.com/Trendyol/es-query-builder/es/enums/multi-values-mode"

type decayFunctionType Object

// Decay creates a new es.decayFunctionType object with the specified field.
//
// This function initializes a decay function configuration for use with decay functions
// (gauss, linear, exp) in a function_score query. The field parameter represents the name
// of the field on which the decay function operates.
//
// Example usage:
//
//	d := es.Decay("date").Origin("now").Scale("10d")
//	// d now contains a decay configuration for the "date" field.
//
// Parameters:
//   - field: A string representing the field name for the decay function.
//
// Returns:
//
//	An es.decayFunctionType object with the specified field.
func Decay(field string) decayFunctionType {
	return decayFunctionType{
		field: Object{},
	}
}

// Origin sets the "origin" parameter in the decay function field configuration.
//
// This method specifies the point of origin used to calculate distance. The origin must
// be provided for numeric and geo fields. For date fields, the default is "now".
//
// Example usage:
//
//	d := es.Decay("date").Origin("now")
//	// d now includes an "origin" parameter set to "now" for the "date" field.
//
// Parameters:
//   - origin: A value representing the origin point. The type depends on the field type.
//
// Returns:
//
//	The updated es.decayFunctionType object with the "origin" parameter set.
func (d decayFunctionType) Origin(origin any) decayFunctionType {
	return d.putInTheField("origin", origin)
}

// Scale sets the "scale" parameter in the decay function field configuration.
//
// This method specifies the distance from the origin at which the computed score equals
// the decay parameter. For geo fields, the scale can be defined as a number with a unit
// (e.g., "1km", "12mi"). For date fields, it can be defined as a duration (e.g., "10d").
//
// Example usage:
//
//	d := es.Decay("date").Origin("now").Scale("10d")
//	// d now includes a "scale" parameter set to "10d" for the "date" field.
//
// Parameters:
//   - scale: A value representing the scale distance.
//
// Returns:
//
//	The updated es.decayFunctionType object with the "scale" parameter set.
func (d decayFunctionType) Scale(scale any) decayFunctionType {
	return d.putInTheField("scale", scale)
}

// Offset sets the "offset" parameter in the decay function field configuration.
//
// This method specifies a distance from the origin within which no decay is applied.
// Documents within this offset distance from the origin receive a score of 1.0.
//
// Example usage:
//
//	d := es.Decay("date").Origin("now").Scale("10d").Offset("5d")
//	// d now includes an "offset" parameter set to "5d" for the "date" field.
//
// Parameters:
//   - offset: A value representing the offset distance.
//
// Returns:
//
//	The updated es.decayFunctionType object with the "offset" parameter set.
func (d decayFunctionType) Offset(offset any) decayFunctionType {
	return d.putInTheField("offset", offset)
}

// DecayValue sets the "decay" parameter in the decay function field configuration.
//
// This method specifies the score at the scale distance from the origin. The default
// value is 0.5 if not specified.
//
// Example usage:
//
//	d := es.Decay("date").Origin("now").Scale("10d").DecayValue(0.5)
//	// d now includes a "decay" parameter set to 0.5 for the "date" field.
//
// Parameters:
//   - decay: A float64 value representing the decay score at the scale distance.
//
// Returns:
//
//	The updated es.decayFunctionType object with the "decay" parameter set.
func (d decayFunctionType) DecayValue(decay float64) decayFunctionType {
	return d.putInTheField("decay", decay)
}

// MultiValueMode sets the "multi_value_mode" parameter in the decay function configuration.
//
// This method specifies how the decay function should handle fields that contain multiple values.
// It determines which value is chosen for the distance calculation.
//
// Example usage:
//
//	d := es.Decay("location").Origin("0,0").Scale("5km").MultiValueMode(MultiValuesMode.Min)
//	// d now includes a "multi_value_mode" parameter set to "min".
//
// Parameters:
//   - multiValueMode: A MultiValuesMode.MultiValuesMode value representing the mode.
//
// Returns:
//
//	The updated es.decayFunctionType object with the "multi_value_mode" parameter set.
func (d decayFunctionType) MultiValueMode(multiValueMode MultiValuesMode.MultiValuesMode) decayFunctionType {
	d["multi_value_mode"] = multiValueMode
	return d
}

func (d decayFunctionType) putInTheField(key string, value any) decayFunctionType {
	return genericPutInTheFieldOfFirstObject(d, key, value)
}
