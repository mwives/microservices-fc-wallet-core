package event

import "time"

type TransactionCreatedEvent struct {
	Name    string
	Payload interface{}
}

func NewTransactionCreatedEvent() *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		Name: "transaction_created",
	}
}

func (e *TransactionCreatedEvent) GetName() string {
	return e.Name
}

func (e *TransactionCreatedEvent) GetDateTime() time.Time {
	return time.Now()
}

func (e *TransactionCreatedEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TransactionCreatedEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}
