package greeter

import (
	"context"

	apiv1 "github.com/scraly/hello-world/internal/services/pkg/v1"
	helloworldv1 "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"
	"go.zenithar.org/pkg/log"
	"golang.org/x/xerrors"
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
		log.Bg().Error("request must not be nil")
		return res, xerrors.Errorf("request must not be nil")
	}

	if req.Name == "" {
		log.Bg().Error("name but not be empty in the request")
		return res, xerrors.Errorf("name but not be empty in the request")
	}

	res.Message = "hello " + req.Name

	return res, nil
}
