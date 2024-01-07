package billbee

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddress_HasHouseNumber(t *testing.T) {
	data := []struct {
		address  Address
		expected bool
	}{
		{
			address:  Address{},
			expected: false,
		},
		{
			address: Address{
				HouseNumber: "",
			},
			expected: false,
		},
		{
			address: Address{
				HouseNumber: " ",
			},
			expected: false,
		},
		{
			address: Address{
				HouseNumber: "1",
			},
			expected: true,
		},
		{
			address: Address{
				HouseNumber: "1a",
			},
			expected: true,
		},
	}

	for i, test := range data {
		// Given

		// When

		// Then
		assert.Equal(
			t,
			test.expected,
			test.address.HasHouseNumber(),
			fmt.Sprintf("#%d failed", i),
		)
	}
}
