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

type AddressPatch struct {
	Street      string `json:"Street"`
	HouseNumber string `json:"HouseNumber"`
	Line2       string `json:"AddressAddition"`
}

func (address Address) HasHouseNumber() bool {
	return len(strings.Trim(address.HouseNumber, " ")) != 0
}

func (address Address) FixHouseNumber() Address {
	// House number is at the start of street text
	re := regexp.MustCompile(`^ *(\d+\S+)(\D*?)$`)
	for _, match := range re.FindAllStringSubmatch(address.Street, -1) {
		address.Street = strings.Trim(match[2], " ")
		address.HouseNumber = match[1]

		return address
	}

	// House number is at the end of street text
	re = regexp.MustCompile(`(\D*?)(\d+.*?)$`)
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

func (address Address) toPatch() AddressPatch {
	return AddressPatch{
		Street:      address.Street,
		HouseNumber: address.HouseNumber,
		Line2:       address.Line2,
	}
}
