package api

import (
    "net/http"
    "encoding/json"
    "server-beaconed/database"
    "fmt"
)
//creating sample response structure
type HelloResponse struct {
    Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(database.Dbi)
    response := HelloResponse{
        Message: "Hello, World!",
    }
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Write(jsonResponse)
}


//GetHistoricalData
type HistoricalPrice struct {
	Id     string  	`json:"id"`
	Date   string  `json:"date"`
	Symbol string  `json:"symbol"`
	Price  string `json:"price"`
}

func GetHistoricalData(w http.ResponseWriter, r *http.Request) {
	db := database.Dbi
	symbol := r.URL.Query().Get("symbol")
	fromDate := r.URL.Query().Get("from_date")
	toDate := r.URL.Query().Get("to_date")
	fmt.Println(symbol)
	fmt.Println(fromDate)
	fmt.Println(toDate)
	// Construct SQL query with an exact match on the symbol
	query := "SELECT id,date, price, symbol FROM historical_prices WHERE symbol = ? AND date BETWEEN ? AND ?"
	rows, err := db.Query(query, symbol, fromDate, toDate)
	if err != nil {
		http.Error(w, "Internal Server Error (DB Query)", http.StatusInternalServerError)
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	var historicalData []HistoricalPrice
	for rows.Next() {
		var hp HistoricalPrice
		err := rows.Scan(&hp.Id, &hp.Date, &hp.Price, &hp.Symbol)
		if err != nil {
			http.Error(w, "Internal Server Error (Scan)", http.StatusInternalServerError)
			fmt.Println("Error scanning row:", err)
			return
		}
		historicalData = append(historicalData, hp)
	}

	responseJSON, err := json.Marshal(historicalData)
	if err != nil {
		http.Error(w, "Internal Server Error (JSON Marshal)", http.StatusInternalServerError)
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write(responseJSON)
}

func FirstTenEntriesHandler(w http.ResponseWriter, r *http.Request) {
	db := database.Dbi
	query := "SELECT id,date,price,symbol FROM historical_prices LIMIT 10"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Internal Server Error (DB Query)", http.StatusInternalServerError)
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	var historicalData []HistoricalPrice
	for rows.Next() {
		var hp HistoricalPrice
		err := rows.Scan(&hp.Id, &hp.Date, &hp.Price, &hp.Symbol)
		if err != nil {
			http.Error(w, "Internal Server Error (Scan)", http.StatusInternalServerError)
			fmt.Println("Error scanning row:", err)
			return
		}
		historicalData = append(historicalData, hp)
	}

	responseJSON, err := json.Marshal(historicalData)
	if err != nil {
		http.Error(w, "Internal Server Error (JSON Marshal)", http.StatusInternalServerError)
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write(responseJSON)
	
}
