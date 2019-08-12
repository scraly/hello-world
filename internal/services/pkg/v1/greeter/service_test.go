package greeter_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"

	helloworldv1 "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"

	"github.com/scraly/hello-world/internal/services/pkg/v1/greeter"
)

func TestSayHello(t *testing.T) {
	testCases := []struct {
		name        string
		req         *helloworldv1.HelloRequest
		message     string
		expectedErr bool
	}{
		{
			name:        "req ok",
			req:         &helloworldv1.HelloRequest{Name: "me"},
			message:     "hello me",
			expectedErr: false,
		},
		{
			name:        "req with empty name",
			req:         &helloworldv1.HelloRequest{},
			expectedErr: true,
		},
		{
			name:        "nil request",
			req:         nil,
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			ctx := context.Background()

			// call
			greeterSvc := greeter.New()
			response, err := greeterSvc.SayHello(ctx, testCase.req)

			t.Log("Got : ", response)

			// assert results expectations
			if testCase.expectedErr {
				g.Expect(response).ToNot(BeNil(), "Result should be nil")
				g.Expect(err).ToNot(BeNil(), "Result should be nil")
			} else {
				g.Expect(response.Message).To(Equal(testCase.message))
			}
		})
	}
}
