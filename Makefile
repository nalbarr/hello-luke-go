PROJ=hello-luke-go
SRC=main.go
BIN=hello-luke-go

all: run

clean:
	rm $(BIN)
	
build: $(SRC)
	go build -o $(BIN)

run: build
	go run $(BIN)
