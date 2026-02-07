package fragmenter_test

import (
	"testing"

	Fragmenter "github.com/Trendyol/es-query-builder/es/enums/fragmenter"

	"github.com/Trendyol/es-query-builder/test/assert"
)

func Test_FragmenterString(t *testing.T) {
	tests := []struct {
		fragmenter Fragmenter.Fragmenter
		result     string
	}{
		{Fragmenter.Simple, "simple"},
		{Fragmenter.Span, "span"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.fragmenter.String())
		})
	}
}
