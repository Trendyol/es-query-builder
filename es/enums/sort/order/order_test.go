package order_test

import (
	Order "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/test/assert"
	"testing"
)

func TestOrderString(t *testing.T) {
	tests := []struct {
		mode   Order.Order
		result string
	}{
		{Order.Asc, "asc"},
		{Order.Desc, "desc"},
		{Order.Default, "_default"},
	}

	for _, test := range tests {
		t.Run(test.result, func(t *testing.T) {
			assert.Equal(t, test.result, test.mode.String())
		})
	}
}
