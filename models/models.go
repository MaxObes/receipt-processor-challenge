package models

// Receipt represents the structure of a receipt
type Receipt struct {
	retailer    string  `json:"retailer"`
	purchaseDate string `json:"purchaseDate"`
	purchaseTime string `json:"purchaseTime"`
	items       []Item  `json:"items"`
	total       string  `json:"total"`
}

// Item represents a single item in the receipt
type Item struct {
	shortDescription string  `json:"shortDescription"`
	price            string  `json:"price"`
}

// ReceiptResponse represents the response with the generated ID
type ReceiptResponse struct {
	ID string `json:"id"`
}

// PointsResponse represents the response with the points
type PointsResponse struct {
	Points int64 `json:"points"`
}