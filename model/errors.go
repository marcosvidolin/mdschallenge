package model

import "fmt"

type ProductNotFoundError struct {
	Country string
	Sku     string
}

func (p *ProductNotFoundError) Error() string {
	return fmt.Sprintf("there is no product with the given country [%s] and sku [%s]", p.Country, p.Sku)
}

type NegativeStokError struct {
	CurAmount int
	ReqAmount int
}

func (n *NegativeStokError) Error() string {
	return fmt.Sprintf("insufficient quantity in stock [current stock: %v | req. amount: %v]", n.CurAmount, n.ReqAmount)
}

type InternalServerError struct {
	Message string
}

func (i *InternalServerError) Error() string {
	return i.Message
}
