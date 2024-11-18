package models

// Receipt represents the structure of a receipt
type Receipt struct {
	Retailer    string  `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items       []Item  `json:"items"`
	Total       string  `json:"total"`
}

// Item represents a single item in the receipt
type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            string  `json:"price"`
}

// ReceiptResponse represents the response with the generated ID
type ReceiptResponse struct {
	ID string `json:"id"`
}

// PointsResponse represents the response with the points
type PointsResponse struct {
	Points int64 `json:"points"`
}