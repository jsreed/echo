package main

import (
	"fmt"
	api "github.com/jsreed/echo/api/echo"
	"github.com/jsreed/echo/internal/echo/impl"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	viper.SetDefault("port", "9000")
}

func main() {
	log.Print("Starting echo")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.Get("port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	echoImpl := &impl.Echo{}

	grpcServer := grpc.NewServer()

	api.RegisterEchoServer(grpcServer, echoImpl)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
