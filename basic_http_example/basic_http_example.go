package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

	http.HandleFunc("/helloworld", helloworldhandler)
	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloworldhandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld", Date: "2021-03-07"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	/*data, err := json.Marshal(response)
	if err != nil {
		panic("Ooooops")
	}
	fmt.Fprint(w, string(data))*/
}
