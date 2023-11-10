package api

import (
    "net/http"
	"github.com/gorilla/mux"
)

func SetupRoutes() {
    router := mux.NewRouter()

    // Sample route
    router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
    http.Handle("/", router)
}
