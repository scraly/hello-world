package greeter

import (
	"context"

	apiv1 "github.com/scraly/hello-world/internal/services/pkg/v1"
	helloworldv1 "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"
	"go.zenithar.org/pkg/log"
)

type service struct {
}

// New services instance
func New() apiv1.Greeter {
	return &service{}
}

// -----------------------------------------------------------------------------

func (s *service) SayHello(ctx context.Context, req *helloworldv1.HelloRequest) (*helloworldv1.HelloReply, error) {
	res := &helloworldv1.HelloReply{}

	// Check request
	if req == nil {
		// res.Error = &helloworldv1.Error{Code: http.StatusBadRequest}
		log.Bg().Error("request must not be nil")
		return res, nil
	}

	res.Message = "hello " + req.Name

	return res, nil
}
