## Example Of Grpc Streaming In This Repo
- Unary RPC (normal request response)
- Bidirectional streaming RPC
- Client streaming RPC
- Server streaming RPC 


## Install Grpc
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

## Export Path Grpc
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
<br/><br/>
<hr></hr>

# Run this following command to make the code run

## Unary RPC
```
<!-- to generate proto file -->
make s-proto

<!-- start server -->
make ss

<!-- start client -->
make sc
```


## Bidirectional streaming RPC
```
<!-- to generate proto file -->
make b-proto

<!-- start server -->
make bs

<!-- start client -->
make bc
```

## Client streaming RPC
```
<!-- to generate proto file -->
make c-proto

<!-- start server -->
make cs

<!-- start client -->
make cc
```

## Server streaming RPC 
```
<!-- to generate proto file -->
make ss-proto

<!-- start server -->
make sss

<!-- start client -->
make ssc
```