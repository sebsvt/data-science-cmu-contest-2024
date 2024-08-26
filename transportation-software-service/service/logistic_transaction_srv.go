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

// InsertMany implements LogisticTransactionService.
func (srv logisticTransactionService) InsertMany(ctx context.Context, transactions []aggregate.LogisticTransaction) error {
	err := srv.repo.InsertMany(ctx, transactions)
	return err
}
