# Go parameters
SERVER_MAIN_PATH=cmd/server/main.go
CLIENT_MAIN_PATH=cmd/client/main.go
BINARY_SERVER_NAME=$(BINARY_PATH)/server
BINARY_CLIENT_NAME=$(BINARY_PATH)/client

BINARY_PATH=bin

server-run:
	go build -o $(BINARY_SERVER_NAME) -race $(SERVER_MAIN_PATH)
	./$(BINARY_SERVER_NAME)

client-run:
	go build -o $(BINARY_CLIENT_NAME) -race $(CLIENT_MAIN_PATH)
	./$(BINARY_CLIENT_NAME)

clean:
	go clean $(CLIENT_MAIN_PATH)
	go clean $(SERVER_MAIN_PATH)
	rm -f $(BINARY_PATH)/*