package types

type Order struct {
	OrderID         int
	ShippingAddress string
	TrackingNumber  string
	OrderDate       string
	Status          string
	TotalSum        int
	Products        []Product
}
