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

		fmt.Printf(
			"Order %d missing house number, street is '%s', line2 is '%s'\n",
			order.Id,
			order.ShippingAddress.Street,
			order.ShippingAddress.Line2,
		)
	}
}
