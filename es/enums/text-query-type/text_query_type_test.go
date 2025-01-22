package textquerytype_test

import (
	"testing"

	TextQueryType "github.com/Trendyol/es-query-builder/es/enums/text-query-type"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_TextQueryTypeString(t *testing.T) {
	tests := []struct {
		textQueryType TextQueryType.TextQueryType
		result        string
	}{
		{TextQueryType.Bestfields, "best_fields"},
		{TextQueryType.Mostfields, "most_fields"},
		{TextQueryType.Crossfields, "cross_fields"},
		{TextQueryType.Phrase, "phrase"},
		{TextQueryType.Phraseprefix, "phrase_prefix"},
		{TextQueryType.Boolprefix, "bool_prefix"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.textQueryType.String())
		})
	}
}
