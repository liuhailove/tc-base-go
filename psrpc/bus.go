package psrpc

import (
	"github.com/liuhailove/tc-base-go/psrpc/internal/bus"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
)

type MessageBus bus.MessageBus

func NewLocalMessageBus() MessageBus {
	return bus.NewLocalMessageBus()
}

func NewNatsMessageBus(nc *nats.Conn) MessageBus {
	return bus.NewNatsMessageBus(nc)
}

func NewRedisMessageBus(rc redis.UniversalClient) MessageBus {
	return bus.NewRedisMessageBus(rc)
}
