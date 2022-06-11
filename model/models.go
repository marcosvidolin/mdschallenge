package model

// Product defines a product
type Product struct {
	Country string `json:"country"`
	Sku     string `json:"sku"`
	Name    string `json:"name"`
	Qtd     int    `json:"qtd`
}

// UpdateQtd decrease the stock given an amount.
// The amount value can be positive or negative
// throws an error if the amount in stock becomes negative with the operation
func (p *Product) UpdateQtd(amount int) (int, error) {
	a := p.Qtd + amount
	if a < 0 {
		return p.Qtd, &NegativeStokError{CurAmount: p.Qtd, ReqAmount: amount}
	}
	p.Qtd = a
	return p.Qtd, nil
}
