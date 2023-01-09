package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Koders \n")
}

func main() {
	http.HandleFunc("/", helloworld)
	fmt.Println("Server started and listening port on localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
