package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received")
		fmt.Fprint(w, "Hello, from load balancer",)
	})
	http.ListenAndServe(":8089", nil)
}