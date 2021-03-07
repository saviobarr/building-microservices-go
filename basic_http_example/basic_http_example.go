package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloworldhandler)
	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloworldhandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooooops")
	}
	fmt.Fprint(w, string(data))
}
