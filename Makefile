gqlgen:
	go run github.com/99designs/gqlgen -v -c ./api/gateway/gqlgen.yml

protobuf:
	protoc --go_out=plugins=grpc:. api/echo/echo.proto

dev:
	skaffold dev