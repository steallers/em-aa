run:
	echo "hello world"

build-grpc:
	echo "building employee connectors..."
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/employee.proto

test:
	cd ..
	$(MAKE) -C ../Server main.go
	ls
	echo "ass"