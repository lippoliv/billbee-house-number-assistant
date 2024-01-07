package main

import (
	"fmt"
	"github.com/lippoliv/billbee-house-number-assistant/billbee"
	"os"
	"time"
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
			"Order %d, Address %d missing housenumber\n",
			order.Id,
			order.ShippingAddress.Id,
		)

		fixedAddress := order.ShippingAddress.FixHouseNumber()
		api.UpdateAddress(fixedAddress)

		fmt.Printf(
			"Order %d, Address %d was fixed\n",
			order.Id,
			fixedAddress.Id,
		)

		// API rate limit
		time.Sleep(1 * time.Second)
	}
}
