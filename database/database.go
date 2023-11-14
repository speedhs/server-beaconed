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
	db, err := sql.Open("sqlite3", "./database/data.sqlite")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		db.Close() 
		return nil, err
	}

	fmt.Println("Database connection established successfully")

	err = createTable(db)
	if err != nil {
		fmt.Println("Error creating table:", err)
		db.Close()
		return nil, err
	}

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

func importData(db *sql.DB) error {
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

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	insertStatement, err := db.Prepare("INSERT INTO historical_prices (id,date, price, symbol) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer insertStatement.Close()

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

