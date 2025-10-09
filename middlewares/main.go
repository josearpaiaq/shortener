package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// create a log message
	message := fmt.Sprintf(
		"Time: %s | METHOD: %s | PATH: %s | IP: %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
	)

	fileName := fmt.Sprintf("logs/log_%s.txt", time.Now().Format("2006-01-02"))

	fmt.Print(message)

	// check if file exists, if not create it and write the message -> logs/log_{today}.txt
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		
		_, err = file.WriteString(message)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		_, err = file.WriteString(message)
		if err != nil {
			fmt.Println(err)
		}
	}

  next.ServeHTTP(w, r)
 })
}