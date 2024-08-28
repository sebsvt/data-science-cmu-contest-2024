package service

import (
	"context"

	"github.com/sebsvt/software-prototype/transportation-software-services/aggregate"
)

type logisticTransactionService struct {
	repo aggregate.LogisticTransactionRepository
}

func NewLogisticTransactionService(repo aggregate.LogisticTransactionRepository) LogisticTransactionService {
	return logisticTransactionService{repo: repo}
}

// GetLogsiticFromTransactionID implements LogisticTransactionService.
func (srv logisticTransactionService) GetLogsiticFromTransactionID(ctx context.Context, transaction_id string) (*aggregate.LogisticTransaction, error) {
	logistic_transactions, err := srv.repo.FromTransactionID(ctx, transaction_id)
	if err != nil {
		return nil, err
	}
	return logistic_transactions, nil
}
