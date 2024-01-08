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

func TestAddress_FixHouseNumber(t *testing.T) {
	data := []struct {
		address  Address
		expected Address
	}{
		{
			address: Address{
				Street: "Str 1a",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1a",
			},
		},
		{
			address: Address{
				Street: "Str 1",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1",
			},
		},
		{
			address: Address{
				Street: "Str1a",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1a",
			},
		},
		{
			address: Address{
				Street: "Str1",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1",
			},
		},
		{
			address: Address{
				Street: "Str",
				Line2:  "1a",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1a",
			},
		},
		{
			address: Address{
				Street: "Str",
				Line2:  "1",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1",
			},
		},
		{
			address: Address{
				Street: "Str 1",
				Line2:  "2",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1",
				Line2:       "2",
			},
		},
		{
			address: Address{
				Street: "Str 1a",
				Line2:  "2b",
			},
			expected: Address{
				Street:      "Str",
				HouseNumber: "1a",
				Line2:       "2b",
			},
		},
	}

	for i, test := range data {
		// Given

		// When
		test.address = test.address.FixHouseNumber()

		// Then
		assert.Equal(
			t,
			test.expected,
			test.address,
			fmt.Sprintf("#%d failed", i),
		)
	}
}

func TestAddress_toPatch(t *testing.T) {
	data := []struct {
		address  Address
		expected AddressPatch
	}{
		{
			address: Address{
				Street:      "Str",
				HouseNumber: "1a",
			},
			expected: AddressPatch{
				Street:      "Str",
				HouseNumber: "1a",
			},
		},
		{
			address: Address{
				Street:      "Str",
				HouseNumber: "1a",
				Line2:       "2",
			},
			expected: AddressPatch{
				Street:      "Str",
				HouseNumber: "1a",
				Line2:       "2",
			},
		},
		{
			address: Address{
				Street:      "Str",
				HouseNumber: "1a",
				Line2:       "",
			},
			expected: AddressPatch{
				Street:      "Str",
				HouseNumber: "1a",
				Line2:       "",
			},
		},
	}

	for i, test := range data {
		// Given

		// When
		patch := test.address.toPatch()

		// Then
		assert.Equal(
			t,
			test.expected,
			patch,
			fmt.Sprintf("#%d failed", i),
		)
	}
}
