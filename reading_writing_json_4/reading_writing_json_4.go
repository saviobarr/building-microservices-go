package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloWorldRequest struct {
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

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var request helloWorldRequest
	err = json.Unmarshal(body, &request)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}
