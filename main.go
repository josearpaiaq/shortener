package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josearpaiaq/shortener/db"
	"github.com/josearpaiaq/shortener/routes"
	"github.com/josearpaiaq/shortener/utils"
)

func main() {
	// Connect to the database
	if _, err := db.ConnectToDB(); err != nil {
		log.Fatal(err)
	}
 
	// Initialize the router
	r := routes.NewRouter()

	// Get port from env
	PORT := utils.GET_ENV("PORT", "8080")

	fmt.Println(PORT)

	fmt.Println("")
	fmt.Println(" 🚀 Server started on port", PORT)
	fmt.Println("")
	fmt.Println(" Press Ctrl+C to stop the server 👋")
	fmt.Println("")
	fmt.Println("----------------------------------")
	fmt.Println("")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), r); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}