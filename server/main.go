package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	var result map[string]int
	json.NewDecoder(req.Body).Decode(&result)
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")

	mapD := make(map[string]int)
	mapD["apple"] = 5
	mapD["lettuce"] = 7
	mapD["banana"] = 8
	mapD["mango"] = 9
	//mapD = map[string]int{"qas": 1, "err": 5}
	bytesRepresentation, err := json.Marshal(mapD)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytesRepresentation))
	json.NewEncoder(w).Encode(bytesRepresentation)
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
