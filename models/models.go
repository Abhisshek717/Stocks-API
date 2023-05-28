package models

// Stock represents a stock.
type Stock struct {
	// StockID is the unique identifier for the stock.
	StockID int64 `json:"id"`
	// Name is the name of the stock.
	Name string `json:"name"`
	// Company is the company that issued the stock.
	Company string `json:"email"`
	// Price is the current price of the stock.
	Price int64 `json:"price"`
}
