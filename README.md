# golb
Simple load balancer implemented in Go. 

# Features
- Request routing between servers using round robin algorithm.
- Health checks

# Management
1. `make build`
2. `make test`
3. `make clean`

# Usage
Assume servers are running locally at `http://localhost:8091` and `http://localhost:8092`
1. `go run main.go http://localhost:8091 http://localhost:8092'` - run load balancer with routing requests between servers.
3. try call load balancer with `curl http://localhost:8089` several times. Requests should be served from different servers.
