// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scraly/hello-world/pkg/protocol/helloworld/v1 (interfaces: GreeterClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/scraly/hello-world/pkg/protocol/helloworld/v1"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGreeterClient is a mock of GreeterClient interface
type MockGreeterClient struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterClientMockRecorder
}

// MockGreeterClientMockRecorder is the mock recorder for MockGreeterClient
type MockGreeterClientMockRecorder struct {
	mock *MockGreeterClient
}

// NewMockGreeterClient creates a new mock instance
func NewMockGreeterClient(ctrl *gomock.Controller) *MockGreeterClient {
	mock := &MockGreeterClient{ctrl: ctrl}
	mock.recorder = &MockGreeterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGreeterClient) EXPECT() *MockGreeterClientMockRecorder {
	return m.recorder
}

// SayHello mocks base method
func (m *MockGreeterClient) SayHello(arg0 context.Context, arg1 *v1.HelloRequest, arg2 ...grpc.CallOption) (*v1.HelloReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SayHello", varargs...)
	ret0, _ := ret[0].(*v1.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello
func (mr *MockGreeterClientMockRecorder) SayHello(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockGreeterClient)(nil).SayHello), varargs...)
}