package main

import (
	"net/http"
	"time"
	"io"
	"fmt"
)

var counter int
var servers []string


func getNextServerAddr() string {
	counter++
	return servers[counter % len(servers)]
}


func redirectToServer(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{
        Timeout: 5 * time.Second, 
    }
	addr := getNextServerAddr()
	if addr == "" {
		http.Error(w, "No servers available", http.StatusInternalServerError)
		return
	}
	proxyReq, err := http.NewRequest(r.Method, addr, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	resp, err := client.Do(proxyReq)
	if err != nil {
		fmt.Println("Downsteam server error", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	for key, value := range resp.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}


func main() {
	servers = []string{"http://localhost:8091", "http://localhost:8092"}
	http.HandleFunc("/", redirectToServer)
	http.ListenAndServe(":8089", nil)
}