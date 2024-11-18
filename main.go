package main

import (
	"net/http"
	"github.com/MaxObes/receipt-processor-challenge/handler"
	"log"
)

func main() {
	//api endpoints
	http.HandleFunc("/receipts/process", processReceipt)
	http.HandleFunc("/receipts/", getPoints)

	// Start the server
	var port string = "8080"
	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}