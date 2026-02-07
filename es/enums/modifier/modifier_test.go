package modifier_test

import (
	"testing"

	Modifier "github.com/Trendyol/es-query-builder/es/enums/modifier"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_ModifierString(t *testing.T) {
	tests := []struct {
		modifier Modifier.Modifier
		result   string
	}{
		{Modifier.None, "none"},
		{Modifier.Log, "log"},
		{Modifier.Log1p, "log1p"},
		{Modifier.Log2p, "log2p"},
		{Modifier.Ln, "ln"},
		{Modifier.Ln1p, "ln1p"},
		{Modifier.Ln2p, "ln2p"},
		{Modifier.Square, "square"},
		{Modifier.Sqrt, "sqrt"},
		{Modifier.Reciprocal, "reciprocal"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.modifier.String())
		})
	}
}
