package types

type Cart struct {
	Id       uint `json:"cart_id"`
	UserId   uint `json:"user_id"`
	Products []Product
}
