package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)


var Dbi *sql.DB
func Connect() (*sql.DB, error) {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "./database/data.sqlite")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	// Ensure the connection is established
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		db.Close() // Close the connection if Ping fails
		return nil, err
	}

	fmt.Println("Database connection established successfully")

	// Create the historical_prices table if it doesn't exist
	err = createTable(db)
	if err != nil {
		fmt.Println("Error creating table:", err)
		db.Close()
		return nil, err
	}

	// Import data from the CSV file
	err = importData(db)
	if err != nil {
		fmt.Println("Error importing data:", err)
		db.Close()
		return nil, err
	}
	Dbi = db
	return db, nil
}


func createTable(db *sql.DB) error {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS historical_prices (
		id INTEGER,
		date TEXT,
		price REAL,
		symbol TEXT
	);
	`

	_, err := db.Exec(createTableQuery)
	return err
}

// importData imports data from a CSV file into the historical_prices table
func importData(db *sql.DB) error {
	// Open the CSV file
	if printTotalEntries(db)>0{
		fmt.Println("Data already present")
		return nil
	}
	fmt.Println("Importing Data")
	file, err := os.Open("database/historical_prices.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Prepare the SQL statement for insertion
	insertStatement, err := db.Prepare("INSERT INTO historical_prices (id,date, price, symbol) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer insertStatement.Close()

	// Iterate over the records and insert into the database
	for _, record := range records {
		_, err := insertStatement.Exec(record[0], record[1], record[2], record[3])
		if err != nil {
			return err
		}
	}

	fmt.Println("Data imported successfully")
	return nil
}
func printTotalEntries(db *sql.DB)  int {
	query := "SELECT COUNT(*) FROM historical_prices"
	var totalEntries int
	err := db.QueryRow(query).Scan(&totalEntries)
	if err != nil {
		fmt.Println("Error querying total entries:", err)
		return -1
	}

	return totalEntries
}
