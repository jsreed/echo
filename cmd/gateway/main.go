package main

import (
	"fmt"
	"github.com/99designs/gqlgen/handler"
	api "github.com/jsreed/echo/api/echo"
	"github.com/jsreed/echo/internal/gateway/generated"
	"github.com/jsreed/echo/internal/gateway/resolvers"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func init() {
	viper.SetDefault("echo_url", "echo:9000")
	viper.SetDefault("port", "4200")
}

func main() {
	log.Print("Starting gateway")

	conn, err := grpc.Dial(viper.GetString("echo_url"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial Failed: %v", err)
	}

	echoClient := api.NewEchoClient(conn)

	http.Handle("/",handler.Playground("echo playground", "/query"))
	http.Handle("/query", handler.GraphQL(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolvers.Resolver{
					EchoCli: echoClient,
				},
			},
		),
	))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("port")), nil))
}
