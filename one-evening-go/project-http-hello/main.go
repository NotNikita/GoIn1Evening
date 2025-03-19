package main

import (
	"fmt"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, ", r.URL.Query().Get("name"))
}

func main() {
	http.HandleFunc("/hello", foo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
