# server-beaconed

Getting Started
Clone the repository:

git clone https://github.com/speedhs/server-beaconed.git
cd server-backend
go run main.go

The server will be available at http://localhost:8080.

Database Setup
SQLite Database:

The application uses an SQLite database located at ./database/data.sqlite.

Import Data:

The historical data is imported from ./database/historical_prices.csv into the database. If data is already present, no import occurs.

API Endpoints
Hello API:

/api/hello (GET): Get a simple "Hello, World!" message.
Example: http://localhost:8080/api/hello
Historical Data API:

/api/historical-data (GET): Get historical data from the database.
Parameters:
symbol (string): Stock symbol.
from_date (string): Start date in the format 'YYYY-MM-DD HH:MM:SS+05:30'.
to_date (string): End date in the format 'YYYY-MM-DD HH:MM:SS+05:30'.
Example: http://localhost:8080/api/historical-data?symbol=NIFTY50&from_date=2023-01-01%2000:00:00%2B05:30&to_date=2023-12-31%2000:00:00%2B05:30
...



Built With
Go
