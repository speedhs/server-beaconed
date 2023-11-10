package api

import (
    "net/http"
    "encoding/json"
)
//creating sample response structure
type HelloResponse struct {
    Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    response := HelloResponse{
        Message: "Hello, World!",
    }
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}
