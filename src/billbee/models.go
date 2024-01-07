package billbee

import (
	"regexp"
	"strings"
)

type OrdersResponse struct {
	Data []Order
}

type Order struct {
	Id              int64   `json:"BillBeeOrderId"`
	ShippingAddress Address `json:"ShippingAddress"`
}

type Address struct {
	Id          int64  `json:"BillbeeId"`
	Street      string `json:"Street"`
	HouseNumber string `json:"HouseNumber"`
	Line2       string `json:"Line2"`
}

func (address Address) HasHouseNumber() bool {
	return len(strings.Trim(address.HouseNumber, " ")) != 0
}

func (address Address) FixHouseNumber() Address {
	// House number is in street text
	re := regexp.MustCompile(`(\D*?)(\d+.*?)$`)
	for _, match := range re.FindAllStringSubmatch(address.Street, -1) {
		address.Street = strings.Trim(match[1], " ")
		address.HouseNumber = match[2]

		return address
	}

	// House number is in line2 text
	for _, match := range re.FindAllStringSubmatch(address.Line2, -1) {
		address.Line2 = strings.Trim(match[1], " ")
		address.HouseNumber = match[2]

		return address
	}

	return address
}
