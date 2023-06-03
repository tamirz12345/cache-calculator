package main

import (
	"cache-calculator/pkg/calculator"
	"cache-calculator/pkg/storage"
	"encoding/json"
	"log"
	"net/http"
)

var (
	cache *storage.CalculatorCache
)

func init() {
	cache = storage.NewCalculatorCache()
}

func main() {
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/plus", plusHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, "Multiply", calculator.Multiply)
}

func plusHandler(w http.ResponseWriter, r *http.Request) {
	handleRequest(w, r, "Plus", calculator.Plus)
}

func handleRequest(w http.ResponseWriter, r *http.Request, operation string, opFunc func([]float64) float64) {
	calculatorRequest, err := calculator.ParseHttpRequest(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	calculatorRequest.Operation = operation
	requestID := calculatorRequest.GenerateRequestId()
	result, ok := cache.Get(requestID)
	if ok {
		log.Printf("Request served from cache. User ID: %s", calculatorRequest.UserId)
	} else {
		result = opFunc(calculatorRequest.Parameters)
		cache.Set(requestID, result)
		log.Printf("Request served from calculation. User ID: %s", calculatorRequest.UserId)
	}
	response := &calculator.Response{
		Result: result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
