package highlightertype_test

import (
	"testing"

	HighlighterType "github.com/Trendyol/es-query-builder/es/enums/highlighter-type"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_HighlighterTypeString(t *testing.T) {
	tests := []struct {
		highlighterType HighlighterType.HighlighterType
		result          string
	}{
		{HighlighterType.Unified, "unified"},
		{HighlighterType.Plain, "plain"},
		{HighlighterType.Fvh, "fvh"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.highlighterType.String())
		})
	}
}
