package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	// internal
	"github.com/jaredonline/roommates-graphql/schema"
)

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()["query"][0]
		result := schema.ExecuteQuery(query)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8888")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8888/graphql?query={users{first_name}}'")
	http.ListenAndServe(":8888", nil)
}
