package natsrpc_test

import (
	"context"
	"testing"

	"github.com/LeKovr/natsrpc"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

// Тестируем подключение клиента к NATS
func TestClientConnection(t *testing.T) {
	// Подключаемся к локальному NATS-серверу
	conn, err := nats.Connect(nats.DefaultURL)
	assert.NoError(t, err, "Должны подключиться без ошибок")
	defer conn.Close()

	// Создаём сервер natsrpc
	server, err := natsrpc.NewServer(conn)
	defer server.Close(context.Background())
	assert.NoError(t, err)

	// Создаём клиента natsrpc
	client := natsrpc.NewClient(conn)
	assert.NotNil(t, client, "Клиент не должен быть nil")
}
