package test

import (
	"context"
	"testing"
	"time"

	"github.com/LeKovr/natsrpc"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct{}

func (t *TestStruct) Test1(ctx context.Context, req *TestRequest) (*TestReply, error) {
	return &TestReply{
		Answer: "test" + req.Name,
	}, nil
}

func (t *TestStruct) Test2(ctx context.Context, req *TestRequest) (*Nil, error) {
	return &Nil{}, nil
}

func TestClient_Request(t *testing.T) {
	natsUrl := "http://89.46.131.181:4222"

    conn, err := nats.Connect(natsUrl)
	assert.NoError(t, err, "Connection error")
	defer conn.Close()

	server, err := natsrpc.NewServer(conn)
	assert.NoError(t, err, "Server creation error")
	defer server.Close(context.Background())

	srv, err := RegisterTestServerNRServer(server, &TestStruct{})
	assert.NoError(t, err, "NRserver register error")
	defer srv.Close()

	client := natsrpc.NewClient(conn)
	cli := NewTestServerNRClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := cli.Test1(ctx, &TestRequest{
		Name: "Test",
	})

	assert.NoError(t, err)
	time.Sleep(time.Second)
	assert.Equal(t, "testTest", reply.Answer)
}

func TestClient_Publish(t *testing.T) {
	natsUrl := "http://89.46.131.181:4222"

    conn, err := nats.Connect(natsUrl)
	assert.NoError(t, err, "Connection error")
	defer conn.Close()

	server, err := natsrpc.NewServer(conn)
	assert.NoError(t, err, "Server creation error")
	defer server.Close(context.Background())

	srv, err := RegisterTestServerNRServer(server, &TestStruct{})
	assert.NoError(t, err, "NRserver register error")
	defer srv.Close()

	client := natsrpc.NewClient(conn)
	cli := NewTestServerNRClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = cli.Test2(ctx, &TestRequest{
		Name: "Test",
	})

	assert.NoError(t, err)
	time.Sleep(time.Second)
}