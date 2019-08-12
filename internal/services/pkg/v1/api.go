package v1

import (
	"context"

	helloworldv1 "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"
)

//go:generate mockgen -destination test/mock/greeter.gen.go -package mock github.com/scraly/hello-world/pkg/protocol/helloworld/v1 GreeterClient

// Greeter service contract
type Greeter interface {
	SayHello(ctx context.Context, req *helloworldv1.HelloRequest) (res *helloworldv1.HelloReply, err error)
}
