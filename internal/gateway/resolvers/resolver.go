package resolvers

import (
	api "github.com/jsreed/echo/api/echo"
	"github.com/jsreed/echo/internal/gateway/generated"
)

type Resolver struct{
	EchoCli api.EchoClient
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}