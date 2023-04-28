BINARY_NAME=driver.exe


run:	
	go run cmd/driver/main.go


test:
	go test .



build: 
	go build -o ${BINARY_NAME} cmd/driver/main.go
