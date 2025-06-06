package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var data Person
		json.Unmarshal(body, &data)

		fmt.Fprintf(w, "Hello %s %s", data.FirstName, data.LastName)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
