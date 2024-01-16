# SingleProcess
s-proto:
	@echo "Generating Go files"
	cd singleProcess/proto && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto

# Single Process Client
sc:
	go build -o bin/client \
		tgrziminiar/grpcStreaming/singleProcess/client
	./bin/client
	
# Single Process Server
ss:
	go build -o bin/server \
		tgrziminiar/grpcStreaming/singleProcess/server
	./bin/server

b-proto:
	@echo "Generating Go files"
	cd bidirectional/proto && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto

# Bidirectional Client
bc:
	go build -o bin/client \
		tgrziminiar/grpcStreaming/bidirectional/client
	./bin/client
	
# Bidirectional Server
bs:
	go build -o bin/server \
		tgrziminiar/grpcStreaming/bidirectional/server
	./bin/server

# ClientStream Proto
c-proto:
	@echo "Generating Go files"
	cd clientStream/proto && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto


# ClientStream Server
cc:
	go build -o bin/client \
		tgrziminiar/grpcStreaming/clientStream/client
	./bin/client

# ClientStream Client
cs:
	go build -o bin/server \
		tgrziminiar/grpcStreaming/clientStream/server
	./bin/server


# ServerStream Proto
ss-proto:
	@echo "Generating Go files"
	cd serverStream/proto && protoc --go_out=. --go-grpc_out=. \
		--go-grpc_opt=paths=source_relative --go_opt=paths=source_relative *.proto


# ServerStream Server
ssc:
	go build -o bin/client \
		tgrziminiar/grpcStreaming/serverStream/client
	./bin/client

# ServerStream Client
sss:
	go build -o bin/server \
		tgrziminiar/grpcStreaming/serverStream/server
	./bin/server



.PHONY: client server protoc