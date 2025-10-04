package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8080"

	fmt.Println("Listening on port " + port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}