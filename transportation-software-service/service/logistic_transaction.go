package service

import (
	"context"

	"github.com/sebsvt/software-prototype/transportation-software-services/domain"
	"github.com/sebsvt/software-prototype/transportation-software-services/entity"
)

type LogisticTransactionCreated struct {
	Consignor entity.Person `json:"consignor"`
	Consignee entity.Person `json:"consignee"`
	Package   entity.Item   `json:"package"`
}

type LogisticTransactionService interface {
	CreateNewTransactionWithItem(ctx context.Context, new_transaction LogisticTransactionCreated, partner_id string) (string, error)
	GetLogsiticFromTransactionID(ctx context.Context, transaction_id string) (*domain.LogisticTransactionAggregate, error)
	GetLogsiticFromID(ctx context.Context, id string) (*domain.LogisticTransactionAggregate, error)
	// InsertMany(ctx context.Context, transactions []aggregate.LogisticTransaction) error
}
