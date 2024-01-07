package billbee

type OrdersResponse struct {
	Data []Order
}

type Order struct {
	Id              int64   `json:"BillBeeOrderId"`
	ShippingAddress Address `json:"ShippingAddress"`
}

type Address struct {
	Id           int64  `json:"BillbeeId"`
	Company      string `json:"Company"`
	NameAddition string `json:"NameAddition"`
	Street       string `json:"Street"`
	HouseNumber  string `json:"HouseNumber"`
	Line2        string `json:"Line2"`
}
