package api

import (
    "net/http"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
    router := mux.NewRouter()

    
    router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
    router.HandleFunc("/api/historical-data", GetHistoricalData).Methods("GET")
    router.HandleFunc("/api/firstten",FirstTenEntriesHandler).Methods("GET")
    router.HandleFunc("/api/login",Login).Methods("GET")
    router.HandleFunc("/api/register",Register).Methods("GET")
    http.Handle("/", router)
}
