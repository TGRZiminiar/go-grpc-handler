# SingleProcess
s-proto:
	@echo "Generating Go files"
	cd SingleProcess/proto && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto

# Single Process Client
sc:
	go build -o bin/client \
		tgrziminiar/grpcStreaming/SingleProcess/client
	./bin/client
	
# Single Process Server
ss:
	go build -o bin/server \
		tgrziminiar/grpcStreaming/SingleProcess/server
	./bin/server



# Bidirectional
all: client server

c:
	./client
	
s:
	./server

protoc:
	@echo "Generating Go files"
	cd src/proto && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto

server: 
	go build -o server \
		tgrziminiar/grpcStreaming/src/server
# protoc
# @echo "Building server"

client: 
	go build -o client \
		tgrziminiar/grpcStreaming/src/client
# @echo "Building client"
# protoc

clean:
	go clean tgrziminiar/grpcStreaming/...
	rm -f server client

.PHONY: client server protoc