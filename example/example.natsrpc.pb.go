// Code generated by protoc-gen-natsrpc. DO NOT EDIT.
// versions:
// - protoc-gen-natsrpc v0.7.0
// source: example.proto

package example

import (
	context "context"
	natsrpc "github.com/LeKovr/natsrpc"
	reflect "reflect"
)

var _ context.Context
var _ reflect.Value
var _ = natsrpc.SupportVersion_0_7_0

const (
	Greeting_NRServiceName = "natsrpc.example.Greeting"
)

type GreetingNRClient interface {
	Hello(ctx context.Context, req *HelloRequest, opt ...natsrpc.CallOption) (*HelloReply, error)
}

type _GreetingNRClientImpl struct {
	c natsrpc.ClientInterface
}

// NewGreetingNRClient
func NewGreetingNRClient(c natsrpc.ClientInterface) GreetingNRClient {
	ret := &_GreetingNRClientImpl{
		c: c,
	}
	return ret
}
func (c *_GreetingNRClientImpl) Hello(ctx context.Context, req *HelloRequest, opt ...natsrpc.CallOption) (*HelloReply, error) {
	rep := &HelloReply{}
	err := c.c.Request(ctx, Greeting_NRServiceName, "Hello", req, rep, opt...)
	if err != nil {
		return nil, err
	}
	return rep, err
}

var Greeting_NRServiceDesc = natsrpc.ServiceDesc{
	ServiceName: Greeting_NRServiceName,
	Methods: []natsrpc.MethodDesc{
		{
			MethodName:  "Hello",
			Handler:     _Greeting_Hello_NRHandler,
			RequestType: reflect.TypeOf(HelloRequest{}),
			IsPublish:   false,
		},
	},
	Metadata: "example.proto",
}

type GreetingNRServer interface {
	Hello(ctx context.Context, req *HelloRequest) (*HelloReply, error)
}

func _Greeting_Hello_NRHandler(svc interface{}, ctx context.Context, req any) (any, error) {
	return svc.(GreetingNRServer).Hello(ctx, req.(*HelloRequest))
}

func RegisterGreetingNRServer(register natsrpc.ServiceRegistrar, s GreetingNRServer, opts ...natsrpc.ServiceOption) (natsrpc.ServiceInterface, error) {
	return register.Register(Greeting_NRServiceDesc, s, opts...)
}

const (
	GreetingToAll_NRServiceName = "natsrpc.example.GreetingToAll"
)

type GreetingToAllNRClient interface {
	HelloToAll(ctx context.Context, notify *HelloRequest, opt ...natsrpc.CallOption) error
}

type _GreetingToAllNRClientImpl struct {
	c natsrpc.ClientInterface
}

// NewGreetingToAllNRClient
func NewGreetingToAllNRClient(c natsrpc.ClientInterface) GreetingToAllNRClient {
	ret := &_GreetingToAllNRClientImpl{
		c: c,
	}
	return ret
}
func (c *_GreetingToAllNRClientImpl) HelloToAll(ctx context.Context, notify *HelloRequest, opt ...natsrpc.CallOption) error {
	return c.c.Publish(ctx, GreetingToAll_NRServiceName, "HelloToAll", notify, opt...)
}

var GreetingToAll_NRServiceDesc = natsrpc.ServiceDesc{
	ServiceName: GreetingToAll_NRServiceName,
	Methods: []natsrpc.MethodDesc{
		{
			MethodName:  "HelloToAll",
			Handler:     _GreetingToAll_HelloToAll_NRHandler,
			RequestType: reflect.TypeOf(HelloRequest{}),
			IsPublish:   true,
		},
	},
	Metadata: "example.proto",
}

type GreetingToAllNRServer interface {
	HelloToAll(ctx context.Context, req *HelloRequest) (*Empty, error)
}

func _GreetingToAll_HelloToAll_NRHandler(svc interface{}, ctx context.Context, req any) (any, error) {
	return svc.(GreetingToAllNRServer).HelloToAll(ctx, req.(*HelloRequest))
}

func RegisterGreetingToAllNRServer(register natsrpc.ServiceRegistrar, s GreetingToAllNRServer, opts ...natsrpc.ServiceOption) (natsrpc.ServiceInterface, error) {
	return register.Register(GreetingToAll_NRServiceDesc, s, opts...)
}
