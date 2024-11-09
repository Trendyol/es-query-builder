package zerotermsquery_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/test/assert"

	ZeroTermsQuery "github.com/Trendyol/es-query-builder/es/enums/match/zero-terms-query"
)

func Test_ZeroTermsQueryString(t *testing.T) {
	tests := []struct {
		zeroTermsQuery ZeroTermsQuery.ZeroTermsQuery
		result         string
	}{
		{ZeroTermsQuery.All, "all"},
		{ZeroTermsQuery.None, "none"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.zeroTermsQuery.String())
		})
	}
}
