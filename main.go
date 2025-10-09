package main

import (
	"fmt"
	"net/http"

	"github.com/josearpaiaq/shortener/router"
)

func main() {
	r := router.NewRouter()

	// TODO: get port from env
	PORT := ":8080"

	fmt.Println("")
	fmt.Println(" 🚀 Server started on port", PORT)
	fmt.Println("")
	fmt.Println(" Press Ctrl+C to stop the server 👋")
	fmt.Println("")
	fmt.Println("----------------------------------")
	fmt.Println("")

	if err := http.ListenAndServe(PORT, r); err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
}