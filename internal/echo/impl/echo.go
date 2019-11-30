package impl

import (
	"github.com/jsreed/echo/api/echo"
	"golang.org/x/net/context"
	"log"
)

type Echo struct {}

func (e Echo) Echo(ctx context.Context, request *api.EchoRequest) (response *api.EchoResponse, err error) {
	log.Print("echoing!!")
	response = &api.EchoResponse{
		RequestPayload: request.Payload,
	}
	return
}
