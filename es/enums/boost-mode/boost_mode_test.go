package boostmode_test

import (
	"testing"

	BoostMode "github.com/Trendyol/es-query-builder/es/enums/boost-mode"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_BoostModeString(t *testing.T) {
	tests := []struct {
		mode   BoostMode.BoostMode
		result string
	}{
		{BoostMode.Multiply, "multiply"},
		{BoostMode.Replace, "replace"},
		{BoostMode.Sum, "sum"},
		{BoostMode.Avg, "avg"},
		{BoostMode.Max, "max"},
		{BoostMode.Min, "min"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.mode.String())
		})
	}
}
