package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/LeKovr/natsrpc"
	"github.com/LeKovr/natsrpc/example"
	"github.com/nats-io/nats.go"
)

var (
	nats_url = flag.String("nats_url", "nats://127.0.0.1:4222", "nats-server地址")
)

func main() {
	conn, err := nats.Connect(*nats_url)
	example.IfNotNilPanic(err)
	defer conn.Close()

	server, err := natsrpc.NewServer(conn,
		natsrpc.WithServerRecovery(func(err interface{}) {
			fmt.Println("recover from", err)
		}),
	)
	example.IfNotNilPanic(err)
	client := natsrpc.NewClient(conn)

	defer server.Close(context.Background())

	svc, err := example.RegisterGreetingNRServer(server, &HelloSvc{})
	example.IfNotNilPanic(err)
	defer svc.Close()

	cli := example.NewGreetingNRClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = cli.Hello(ctx, &example.HelloRequest{
		Name: "bruce",
	})
	fmt.Println(err.Error())
}

type HelloSvc struct {
}

func (s *HelloSvc) Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloReply, error) {
	panic("some error")
}
