BINARY_NAME=driver.exe


run:	
	go run cmd/driver/main.go


test:
	go test .



build: 
	go build -o ${BINARY_NAME} cmd/driver/main.go

protogen: 
	protoc -I protocol --go-grpc_out=pkg/driver protocol/stream.proto
