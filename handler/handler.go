package handler

import (
	"encoding/json"
	"net/http"
	"github.com/MaxObes/receipt-processor-challenge/models"
	"github.com/google/uuid"
	"math"
	"strconv"
	"time"
	"fmt"
	"regexp"
	"strings"
)

// local data option
// initialize a map with a string key (id of receipt)
// value is a receipt object
var receipts = make(map[string]models.Receipt)

// calculates the total points for a receipt based on the rules
func calculatePoints(receipt models.Receipt) int64 {
	points := int64(0)

	// Rule 1: 1 point for each alphanumeric characters in retailer name
	points += int64(countAlphanumeric(receipt.Retailer))

	// Rule 2: 50 points for round dollar amounts with no cents
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil {
		if total == math.Floor(total) {
			points += 50
		}
	} else {
		fmt.Println("Error:", err)
	}

	// Rule 3: 25 points for totals that are multiples of 0.25
	if total > 0 && math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt
	points += 5 * int64(math.Floor(float64(len(receipt.Items))/2))

	// Rule 5: Points for item descriptions that are multiples of 3 in length
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err == nil {
				itemPoints := int64(math.Ceil(price * 0.2))
				points += itemPoints
			}
		}
	}

	// Rule 6: Points if the day of the purchase is odd
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: Points if the purchase time is between 2:00pm and 4:00pm (exclusive)
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil {
		if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
			points += 10
		}
	}

	return points
}

// getPoints handles retrieving points for a receipt by ID
func GetPoints(w http.ResponseWriter, r *http.Request) {
	receiptID := strings.TrimPrefix(r.URL.Path, "/receipts/")
	receiptID = strings.TrimSuffix(receiptID, "/points")

	receipt, exists := receipts[receiptID]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Calculate points based on receipt total (for simplicity)
	points := calculatePoints(receipt)

	// Return the points
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.PointsResponse{Points: points})
}

// counts the number of alphanumeric characters in a string
func countAlphanumeric(s string) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(re.FindAllString(s, -1)) // -1 find all matches
}

// processReceipt handles receipt processing and generates an ID
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var receipt models.Receipt 
	err := json.NewDecoder(r.Body).Decode(&receipt)
	
	if err != nil {
		http.Error(w, "Invalid receipt data", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt
	receiptID, _ := generateReceiptID()

	// Store the receipt in memory
	receipts[receiptID] = receipt

	// Return the receipt ID in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.ReceiptResponse{ID: receiptID})
}

// generateReceiptID generates a random UUID ID
func generateReceiptID() (string, error) {
	// Generate a new UUID
	id := uuid.New()

	// Return the UUID as a string
	return id.String(), nil
}