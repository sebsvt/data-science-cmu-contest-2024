package service

import (
	"context"

	"github.com/sebsvt/software-prototype/transportation-software-services/aggregate"
)

type LogisticTransactionService interface {
	// CreateNewTransaction(trasaction aggregate.LogisticTransaction)
	InsertMany(ctx context.Context, transactions []aggregate.LogisticTransaction) error
}
