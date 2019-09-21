all: run

clean:
	rm hello-luke-go
	
build:
	go build -o hello-luke-go

run: build
	go run main.go
