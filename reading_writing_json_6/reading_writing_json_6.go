package main

import (
	"log"
	"net/http"
)

func main() {
	cathandler := http.FileServer(http.Dir("/tmp/images"))
	http.Handle("/cat/", http.StripPrefix("/cat", cathandler))
	log.Fatal(http.ListenAndServe(":8080", cathandler))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	cathandler := http.FileServer(http.Dir("/tmp/images"))
	http.Handle("/cat/", http.StripPrefix("/cat", cathandler))

}
