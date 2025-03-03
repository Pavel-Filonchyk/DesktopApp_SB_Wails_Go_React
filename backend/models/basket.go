package models

type BasketItem struct {
	Amount      int     `json:"amount"`
	Cost        int     `json:"cost"`
	Discriptions string `json:"discriptions"`
	ID          string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
}