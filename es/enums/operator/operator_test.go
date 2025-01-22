package operator_test

import (
	"testing"

	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_OperatorString(t *testing.T) {
	tests := []struct {
		operator Operator.Operator
		result   string
	}{
		{Operator.And, "and"},
		{Operator.Or, "or"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.operator.String())
		})
	}
}
