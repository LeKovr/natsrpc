package natsrpc_test

import (
	"context"
	"testing"

	"github.com/LeKovr/natsrpc"
	"github.com/LeKovr/natsrpc/example"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

// Реализация обработчика сервиса
type TestService struct{}

func (s *TestService) Hello(ctx context.Context, req *example.HelloRequest) (*example.HelloReply, error) {
	return &example.HelloReply{
		Message: "Hello, " + req.Name,
	}, nil
}

// Тестируем RPC-вызов
func TestRPC(t *testing.T) {
	// Подключаемся к локальному NATS-серверу
	conn, err := nats.Connect(nats.DefaultURL)
	defer conn.Close()

	assert.NoError(t, err, "Должны подключиться без ошибок")

	// Создаём сервер natsrpc
	server, err := natsrpc.NewServer(conn)
	defer server.Close(context.Background())

	assert.NoError(t, err)

	// Регистрируем сервис
	tstService := &TestService{}
	svc, err := example.RegisterGreetingNRServer(server, tstService)
	defer svc.Close()

	assert.NoError(t, err)

	// Создаём клиента
	client := natsrpc.NewClient(conn)

	// Отправляем запрос через NATS
	name := "Tester"
	req := &example.HelloRequest{Name: name}
	cli := example.NewGreetingNRClient(client)
	rsp, err := cli.Hello(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Hello, "+name, rsp.Message)
}
