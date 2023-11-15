# server-beaconed
### 1. Database Setup
- [x] Insert daily price data for NIFTY and BANK NIFTY indices into an SQLite database.

### 2. API for Historical Data
- [x] Create a `GET /historical-data` API.
  - [x] Implement query parameters for `symbol`, `from_date`, and `to_date`.
  - [x] Query the database to return data for the input symbol between the specified dates.

### 3. User Authentication Endpoints
- [x] Create authentication endpoints.
  - [x] GET `/user/login`
  - [x]  GET `/user/register`
- [x] Store user details in the SQLite database.

### 4. Mock Response Endpoints
- [x] Create endpoints for mock responses.
  - [x] GET `/portfolio/holdings` (Returns *holdings_response.json*)
  - [x] GET `/user/profile` (Returns *profile_response.json*)
  - [x] POST `/order/place_order` (Returns *place_order_response.json*)


Built With
Go
