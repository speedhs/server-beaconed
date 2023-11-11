package main

import (
    "net/http"
    "server-beaconed/api"
    //"server-beaconed/database"
	"fmt"
    
)

func main() {
	
fmt.Println("Setting up routes...")
    //database.Connect()
    api.SetupRoutes()
	fmt.Println("Starting server at 8080")
    http.ListenAndServe(":8080", nil)
}
