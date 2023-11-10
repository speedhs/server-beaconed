package main

import (
    "net/http"
    "server-beaconed/api"
	"fmt"
    
)

func main() {
   
	fmt.Println("Setting up routes...")
    api.SetupRoutes()
	fmt.Println("Starting server at 8080")
    http.ListenAndServe(":8080", nil)
}
