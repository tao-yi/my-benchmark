package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	port := ":3999"
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"message": "hello world",
		})
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
