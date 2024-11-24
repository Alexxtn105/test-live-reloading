package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/test", testHandler)

	addr := "localhost:8080"

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello world!!!"))
}
