package main

import (
	"fmt"
	"github.com/lippoliv/billbee-house-number-assistant/billbee"
	"os"
)

func main() {
	api := billbee.NewApiClient(
		os.Getenv("BILLBEE_USER"),
		os.Getenv("BILLBEE_PASSWORD"),
		os.Getenv("BILLBEE_API_KEY"),
	)

	orders := api.GetLastOrders(0)
	for _, order := range orders {
		if order.ShippingAddress.HasHouseNumber() {
			continue
		}

		fixedAddress := order.ShippingAddress.FixHouseNumber()
		fmt.Printf(
			"Order %d was missing house number, street is '%s', line2 is '%s', housenumber is '%s'\n",
			order.Id,
			fixedAddress.Street,
			fixedAddress.Line2,
			fixedAddress.HouseNumber,
		)
	}
}
