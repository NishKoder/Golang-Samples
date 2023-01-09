package backend

import (
	"fmt"
	"log"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
func Run(addr string) {
	http.HandleFunc("/", helloworld)
	fmt.Println("Sever started and listening port:", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
