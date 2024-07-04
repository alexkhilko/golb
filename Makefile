build:
	go build -o bin/golb main.go

clean:
	rm -rf bin/

test:
	go test -v ./...