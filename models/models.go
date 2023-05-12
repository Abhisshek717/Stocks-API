package models

type Stock struct {
	StockID    int64    `json:"id"`
	Name  string `json:"name"`
	Company string `json:"email"`
	Price int64 `json:"price"`
}