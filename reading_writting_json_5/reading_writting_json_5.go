package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`

	//do not output this field
	Author string `json:"-"`

	//do not output this field if the value is empty
	Date string `json:",omitempty"`

	//convert output to a string and rename id
	Id int `json:"id, string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var request HelloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name + " this is using encode"}

	encoder := json.NewEncoder(w)

	encoder.Encode(response)
}
