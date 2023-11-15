package api

import (
    "net/http"
    "encoding/json"
    "server-beaconed/database"
    "fmt"
	"github.com/dgrijalva/jwt-go"
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

type LoginResponse struct {
	Message string `json:message`
	Email string `json:"username"`
	Token    string `json:"token"`
}

type User struct {
    ID         	string    `json:"user_id"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := database.Dbi
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	query := "SELECT user_id FROM users WHERE email = ? AND password_hash = ?"
	row := db.QueryRow(query, email, password)
	var user User
	err := row.Scan(&user.ID)

	if err != nil {
        http.Error(w, "User not found.", http.StatusUnauthorized)
		fmt.Println("No user")
    	return
}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: 0, 
	})

	tokenString, err := token.SignedString([]byte("truebeaconbyharsh"))
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		Message:  "Login successful",
		Email: email,
		Token:    tokenString,
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


type RegisterUser struct {
	UserID       string `json:"user_id"`
	UserType     string `json:"user_type"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	Broker       string `json:"broker"`
	Password     string `json:"password_hash"`
}

type RegisterResponse struct{
	Message		string `json:"message"`
}
func Register(w http.ResponseWriter, r *http.Request) {
	db := database.Dbi
	stmt, err := db.Prepare("INSERT INTO users (user_id, user_type, email, user_name, broker,password_hash) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 
	}
	UserId:= r.URL.Query().Get("userid")
	UserType := r.URL.Query().Get("usertype")
	Email := r.URL.Query().Get("email")
	UserName := r.URL.Query().Get("username")
	Broker := r.URL.Query().Get("broker")
	Password := r.URL.Query().Get("password")
	
	_, err = stmt.Exec(UserId, UserType, Email, UserName, Broker, Password)
	if err != nil {
		return 
	}

	if err != nil {
		http.Error(w, "Error inserting user data", http.StatusInternalServerError)
		return
	}


	response := RegisterResponse{
		Message: "Registration successful",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write(jsonResponse)
}

func GetHoldingsHandler(w http.ResponseWriter, r *http.Request) {
	// Load mock data from holdings_response.json
	holdingsData := `{
		"status": "success",
		"data": [
			{
				"tradingsymbol": "GOLDBEES",
				"exchange": "BSE",
				"isin": "INF204KB17I5",
				"quantity": 2,
				"authorised_date": "2021-06-08 00:00:00",
				"average_price": 40.67,
				"last_price": 42.47,
				"close_price": 42.28,
				"pnl": 3.5999999999999943,
				"day_change": 0.18999999999999773,
				"day_change_percentage": 0.44938505203405327
			},
			{
				"tradingsymbol": "IDEA",
				"exchange": "NSE",
				"isin": "INE669E01016",
				"quantity": 5,
				"authorised_date": "2021-06-08 00:00:00",
				"average_price": 8.466,
				"last_price": 10,
				"close_price": 10.1,
				"pnl": 7.6700000000000035,
				"day_change": -0.09999999999999964,
				"day_change_percentage": -0.9900990099009866
			}
		]
	}`

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte(holdingsData))
}

func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Load mock data from profile_response.json
	profileData := `{
		"status": "success",
		"data": {
			"user_id": "AB1234",
			"user_type": "individual",
			"email": "xxxyyy@gmail.com",
			"user_name": "AxAx Bxx",
			"broker": "ZERODHA"
		}
	}`

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte(profileData))
}

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Load mock data from place_order_response.json
	placeOrderData := `{
		"status": "success",
		"data": {
			"message": "Order Placed Successfully",
			"order_id": "151220000000000"
		}
	}`

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte(placeOrderData))
}
