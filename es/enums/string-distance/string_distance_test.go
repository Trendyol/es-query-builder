package stringdistance_test

import (
	"testing"

	StringDistance "github.com/Trendyol/es-query-builder/es/enums/string-distance"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_StringDistanceString(t *testing.T) {
	tests := []struct {
		stringDistance StringDistance.StringDistance
		result         string
	}{
		{StringDistance.Internal, "internal"},
		{StringDistance.DamerauLevenshtein, "damerau_levenshtein"},
		{StringDistance.Levenshtein, "levenshtein"},
		{StringDistance.JaroWinkler, "jaro_winkler"},
		{StringDistance.Ngram, "ngram"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.stringDistance.String())
		})
	}
}
