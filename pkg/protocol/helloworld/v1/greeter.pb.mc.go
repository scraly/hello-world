// Code generated by protoc-gen-defaults. DO NOT EDIT.

package helloworldv1

import (
	"context"
	"github.com/bxcodec/faker"
)

// MockGreeterServer is the mock implementation of the GreeterServer. Use this to create mock services that
// return random data. Useful in UI Testing.
type MockGreeterServer struct{}

// SayHello is mock implementation of the method SayHello
func (MockGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	var res HelloReply
	if err := faker.FakeData(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
