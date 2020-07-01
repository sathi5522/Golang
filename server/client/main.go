package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	makeRequest()
}

func makeRequest() {

	/*message := map[string]interface{}{
		"hello": "world",
		"life":  42,
		"embedded": map[string]string{
			"yes": "of course!",
		},
	}*/

	mapD := map[string]int{"banana": 8, "mango": 9}
	bytesRepresentation, err := json.Marshal(mapD)
	if err != nil {
		log.Fatalln(err)
	}
	req, _ := http.NewRequest("POST", "http://localhost:8080/headers", bytes.NewBuffer(bytesRepresentation))
	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var resultmap map[string]int
	err = json.NewDecoder(resp.Body).Decode(&resultmap)
	// result, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	defer resp.Body.Close()
	//json.NewDecoder(resp.Body).Decode(&result)

	//fmt.Printf("%s", result)
	fmt.Println("result is", resultmap)
	//log.Println("data: ", result["data"])
}
