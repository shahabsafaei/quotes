package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes []Quote

func loadQuotes(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("❌ Failed to open quotes file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&quotes)
	if err != nil {
		log.Fatalf("❌ Failed to decode quotes: %v", err)
	}
	log.Printf("✅ Loaded %d quotes", len(quotes))
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[rand.Intn(len(quotes))]
	json.NewEncoder(w).Encode(randomQuote)
}

func main() {
	loadQuotes("quotes.json")

	http.HandleFunc("/", quoteHandler)

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
