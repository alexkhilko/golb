package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/alexkhilko/golb/servers"
)

var (
	pool               *servers.Pool
	client             *http.Client
	heathCheckInterval int
)

func redirectToServer(w http.ResponseWriter, r *http.Request) {
	addr := pool.GetNextServerAddr()
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

func healthCheckServers() {
	client = &http.Client{
		Timeout: 1 * time.Second,
	}
	healthy, unhealthy := pool.GetHealthyServerURLs(), pool.GetUnhealthyServerURLs()
	for {
		for _, server := range healthy {
			resp, err := client.Get(server)
			if err != nil || resp.StatusCode != http.StatusOK {
				fmt.Printf("healthcheck failed for %s with %s\n", server, err)
				pool.Suspend(server)
				continue
			}
		}
		for _, server := range unhealthy {
			resp, err := client.Get(server)
			if err == nil && resp.StatusCode == http.StatusOK {
				fmt.Printf("server %s recovered\n", server)
				pool.Activate(server)
				continue
			}
		}
		time.Sleep(time.Duration(heathCheckInterval) * time.Second)
	}
}

func validateServerURLs(urls []string) []string {
	if urls == nil {
		return []string{}
	}
	validUrls := []string{}
	for _, server := range urls {
		_, err := url.ParseRequestURI(server)
		if err != nil {
			fmt.Println("Incorrect url", server)
			continue
		}
		validUrls = append(validUrls, server)
	}
	return validUrls
}

func main() {
	flag.IntVar(&heathCheckInterval, "h", 5, "Interval between health checks in seconds")
	flag.Parse()
	serverURLs := validateServerURLs(flag.Args())
	pool = servers.NewPool(serverURLs)
	go healthCheckServers()

	client = &http.Client{
		Timeout: 5 * time.Second,
	}
	http.HandleFunc("/", redirectToServer)
	http.ListenAndServe(":8089", nil)
}
