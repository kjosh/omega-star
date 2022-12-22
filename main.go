package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get the current time
		now := time.Now()

		// Format the time in a non-ISO-conforming way
		timeString := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%03d",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second(),
			now.Nanosecond()/1000000)

		// Write the time string to the response
		w.Write([]byte(timeString))
	})

	fmt.Println("omega-star returns non-ISO-timestamps. Listening on :8082")
	http.ListenAndServe(":8082", nil)
}
