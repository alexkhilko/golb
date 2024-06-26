package main

import (
	"net/http"
	"time"
	"io"
	"fmt"
)


func getNextServerAddr() string {
	return "http://localhost:8091"
}


func redirectToServer(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{
        Timeout: 5 * time.Second, 
    }
	addr := getNextServerAddr()
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


	// Copy headers and status code from the downstream response to the client response
	for key, value := range resp.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}
	w.WriteHeader(resp.StatusCode)

	// Copy the response body from the downstream server to the client
	io.Copy(w, resp.Body)
}


func main() {
	http.HandleFunc("/", redirectToServer)
	http.ListenAndServe(":8089", nil)
}