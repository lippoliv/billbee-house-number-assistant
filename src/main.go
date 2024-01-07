package main

import (
	"fmt"
	"github.com/lippoliv/billbee-house-number-assistant/billbee"
	"os"
	"strconv"
	"time"
)

func main() {
	api := billbee.NewApiClient(
		os.Getenv("BILLBEE_USER"),
		os.Getenv("BILLBEE_PASSWORD"),
		os.Getenv("BILLBEE_API_KEY"),
	)

	lastOrderId := int64(0)
	runInterval, err := strconv.ParseInt(os.Getenv("RUN_INTERVAL"), 10, 16)
	if err != nil {
		runInterval = 300
	}

	for {
		orders := api.GetLastOrders(lastOrderId + 1)
		fmt.Printf("Check %d orders\n", len(orders))
		for _, order := range orders {
			if order.Id > lastOrderId {
				lastOrderId = order.Id
			}

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

		time.Sleep(time.Duration(runInterval) * time.Second)
	}
}
