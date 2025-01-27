package handler

import (
	"fmt"
	"sync"

	"github.com/mwives/event-driven-architecture-fc-3.0/walletcore/pkg/events"
	"github.com/mwives/event-driven-architecture-fc-3.0/walletcore/pkg/kafka"
)

type BalanceUpdatedKafkaHandler struct {
	Kakfa *kafka.Producer
}

func NewBalanceUpdatedKafkaHandler(kafka *kafka.Producer) *BalanceUpdatedKafkaHandler {
	return &BalanceUpdatedKafkaHandler{Kakfa: kafka}
}

func (h *BalanceUpdatedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kakfa.Publish(message, nil, "balances")
	fmt.Printf("BalanceUpdatedKafkaHandler: %v\n", message.GetPayload())
}
