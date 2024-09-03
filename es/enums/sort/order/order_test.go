package order_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/test/assert"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"
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
