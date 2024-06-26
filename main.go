package main

import (
	"net/http"
	"time"
	"io"
	"fmt"
	"flag"
	"net/url"
)

var (
	counter int
	servers []string
	client *http.Client
	heathCheckInterval int
)


func getNextServerAddr() string {
	if len(servers) == 0 {
		return ""
	}
	counter++
	return servers[counter % len(servers)]
}


func redirectToServer(w http.ResponseWriter, r *http.Request) {
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

func healthCheckServers(initialServers []string) {
	client = &http.Client{
        Timeout: 1 * time.Second, 
    }
	for {
		healthyServers := []string{}
		for _, server := range initialServers {
			resp, err := client.Get(server)
			if err != nil {
				fmt.Printf("failed to check health of %s with %s\n", server, err)
				continue
			} 
			if resp.StatusCode != http.StatusOK {
				fmt.Printf("healthcheck failed for %s with %d\n", server, resp.StatusCode)
				continue
			}
			healthyServers = append(healthyServers, server)
		}
		servers = healthyServers
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
	servers = validateServerURLs(flag.Args())
	go healthCheckServers(servers)

	client = &http.Client{
        Timeout: 5 * time.Second, 
    }

	http.HandleFunc("/", redirectToServer)
	http.ListenAndServe(":8089", nil)
}