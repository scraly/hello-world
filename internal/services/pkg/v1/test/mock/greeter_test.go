package mock_test

import (
	context "context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/scraly/hello-world/internal/services/pkg/v1/test/mock"
	helloworldv1 "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"
)

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockGreeterClient := mock.NewMockGreeterClient(ctrl)
	req := &helloworldv1.HelloRequest{Name: "me"}

	// set-up mock with a helloreply returned
	mockGreeterClient.EXPECT().SayHello(
		gomock.Any(),
		req,
	).Return(&helloworldv1.HelloReply{}, nil)
	testSayHelloOK(t, mockGreeterClient)
}

func testSayHelloOK(t *testing.T, client helloworldv1.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &helloworldv1.HelloRequest{})
	t.Log("Reply : ", r)
	if err != nil || r.Message != "hello me" {
		t.Errorf("mocking failed")
	}
	t.Log("Reply : ", r.Message)
}
