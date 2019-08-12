package v1

import (
	"context"

	helloworldv1 "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"
)

// Greeter service contract
type Greeter interface {
	SayHello(ctx context.Context, req *helloworldv1.HelloRequest) (res *helloworldv1.HelloReply, err error)
}
