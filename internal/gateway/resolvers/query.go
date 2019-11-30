package resolvers

import (
	"context"
	api "github.com/jsreed/echo/api/echo"
	"log"
)

type queryResolver struct{ *Resolver }

func (q queryResolver) Echo(ctx context.Context, payload string) (returnPayload string, err error) {
	log.Print("handling echo query")

	echoResponse, err := q.EchoCli.Echo(ctx, &api.EchoRequest{
		Payload: payload,
	})
	if err != nil {
		log.Print(err)
	}
	if echoResponse == nil {
		return "nil echoResponse", nil
	}

	return echoResponse.RequestPayload, err
}


