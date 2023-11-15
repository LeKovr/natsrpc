package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/byebyebruce/natsrpc"
	"github.com/byebyebruce/natsrpc/example"
	"github.com/nats-io/nats.go"
)

var (
	nats_url = flag.String("nats_url", "nats://127.0.0.1:4222", "nats-server地址")
)

func main() {
	conn, err := nats.Connect(*nats_url)
	example.IfNotNilPanic(err)
	defer conn.Close()

	server, err := natsrpc.NewServer(conn)
	example.IfNotNilPanic(err)

	defer server.Close(context.Background())

	const namespace = "example"

	svc, err := example.RegisterGreetingNATSRPCServer(server, &HelloSvc{namespace: namespace},
		natsrpc.WithServiceNamespace(namespace))
	example.IfNotNilPanic(err)
	defer svc.Close()

	cli := example.NewGreetingNATSRPCClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = cli.Hello(ctx, &example.HelloRequest{
		Name: "bruce",
	})
	fmt.Println(err.Error())

	cli1 := example.NewGreetingNATSRPCClient(conn,
		natsrpc.WithClientNamespace(namespace))
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rep, err := cli1.Hello(ctx, &example.HelloRequest{
		Name: "bruce",
	})
	example.IfNotNilPanic(err)
	fmt.Println(rep.Message)
}

type HelloSvc struct {
	namespace string
}

func (s *HelloSvc) Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloReply, error) {
	return &example.HelloReply{
		Message: "hello " + req.Name + " from " + s.namespace,
	}, nil
}
