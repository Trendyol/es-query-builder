package multivaluesmode_test

import (
	"testing"

	MultiValuesMode "github.com/Trendyol/es-query-builder/es/enums/multi-values-mode"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_MultiValuesModeString(t *testing.T) {
	tests := []struct {
		mode   MultiValuesMode.MultiValuesMode
		result string
	}{
		{MultiValuesMode.Min, "min"},
		{MultiValuesMode.Max, "max"},
		{MultiValuesMode.Avg, "avg"},
		{MultiValuesMode.Sum, "sum"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.mode.String())
		})
	}
}
