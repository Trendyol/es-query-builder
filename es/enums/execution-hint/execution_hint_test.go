package executionhint_test

import (
	"testing"

	ExecutionHint "github.com/Trendyol/es-query-builder/es/enums/execution-hint"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ExecutionHintString(t *testing.T) {
	tests := []struct {
		executionHint ExecutionHint.ExecutionHint
		result        string
	}{
		{ExecutionHint.Map, "map"},
		{ExecutionHint.GlobalOrdinals, "global_ordinals"},
		{ExecutionHint.FieldData, "fielddata"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.executionHint.String())
		})
	}
}
