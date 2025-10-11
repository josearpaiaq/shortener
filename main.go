package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/josearpaiaq/shortener/db"
	"github.com/josearpaiaq/shortener/router"
	"github.com/josearpaiaq/shortener/utils"
	_ "github.com/lib/pq"
)

func main() {
	db, err := db.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
 
	r := router.NewRouter()

	// TODO: get port from env
	// PORT := os.Getenv("PORT")
	PORT := utils.GET_ENV("PORT")

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