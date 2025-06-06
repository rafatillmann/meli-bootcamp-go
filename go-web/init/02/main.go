package main

import (
	"encoding/json"
	"fmt"
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

		var data Person
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&data); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		response := fmt.Sprintf("Hello %s %s!", data.FirstName, data.LastName)
		w.Write([]byte(response))

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
