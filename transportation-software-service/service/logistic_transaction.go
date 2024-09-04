package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sebsvt/software-prototype/transportation-software-services/aggregate"
	"github.com/sebsvt/software-prototype/transportation-software-services/entity"
)

type LogisticTransactionCreated struct{}

type LogisticTransactionService interface {
	CreateNewTransactionWithItem(ctx context.Context, consignor entity.Person, consignee entity.Person, package_item entity.Item, partner_id uuid.UUID) (uuid.UUID, error)
	GetLogsiticFromTransactionID(ctx context.Context, transaction_id uuid.UUID) (*aggregate.LogisticTransaction, error)
	// InsertMany(ctx context.Context, transactions []aggregate.LogisticTransaction) error
}
