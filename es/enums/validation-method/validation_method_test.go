package validationmethod_test

import (
	"testing"

	ValidationMethod "github.com/Trendyol/es-query-builder/es/enums/validation-method"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ValidationMethodString(t *testing.T) {
	tests := []struct {
		validationMethod ValidationMethod.ValidationMethod
		result           string
	}{
		{ValidationMethod.Strict, "STRICT"},
		{ValidationMethod.IgnoreMalformed, "IGNORE_MALFORMED"},
		{ValidationMethod.Coerce, "COERCE"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.validationMethod.String())
		})
	}
}
