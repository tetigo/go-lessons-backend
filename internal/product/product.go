package product

import (
	"github.com/Rhymond/go-money"
)

type Product struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	PriceVATExcluded Amount `json:"priceVatExcluded"`
	VAT              Amount `json:"vat"`
	TotalPrice       Amount `json:"totalPrice"`
	Image            string `json:"image"`
}

type Amount struct {
	Money   *money.Money `json:"money"`
	Display string       `json:"display"`
}
