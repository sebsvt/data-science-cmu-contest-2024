package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sebsvt/software-prototype/transportation-software-services/domain"
)

type logisticTransactionService struct {
	repo domain.LogisticTransactionRepository
}

func NewLogisticTransactionService(repo domain.LogisticTransactionRepository) LogisticTransactionService {
	return logisticTransactionService{repo: repo}
}

// GetLogsiticFromTransactionID implements LogisticTransactionService.
func (srv logisticTransactionService) GetLogsiticFromTransactionID(ctx context.Context, transaction_id string) (*domain.LogisticTransactionAggregate, error) {
	logistic_transactions, err := srv.repo.FromTransactionID(ctx, transaction_id)
	if err != nil {
		return nil, err
	}
	return logistic_transactions, nil
}

func (srv logisticTransactionService) CreateNewTransactionWithItem(ctx context.Context, transaction LogisticTransactionCreated, partner_id string) (string, error) {
	transaction_id := uuid.New().String()
	new_transaction := domain.LogisticTransactionAggregate{
		TransactionID: transaction_id,
		Consignor:     transaction.Consignor,
		Consignee:     transaction.Consignee,
		Status:        "pending",
		PartnerID:     partner_id,
		Package:       transaction.Package,
		Timestamp:     time.Now(),
		DeletedAt:     nil,
	}
	if err := srv.repo.Save(ctx, &new_transaction); err != nil {
		fmt.Println(err)
		return "", err
	}
	return transaction_id, nil
}

func (srv logisticTransactionService) GetLogsiticFromID(ctx context.Context, id string) (*domain.LogisticTransactionAggregate, error) {
	transaction, err := srv.repo.FromID(ctx, id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
