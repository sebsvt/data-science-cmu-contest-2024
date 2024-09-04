package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sebsvt/software-prototype/transportation-software-services/aggregate"
	"github.com/sebsvt/software-prototype/transportation-software-services/entity"
)

type logisticTransactionService struct {
	repo aggregate.LogisticTransactionRepository
}

func NewLogisticTransactionService(repo aggregate.LogisticTransactionRepository) LogisticTransactionService {
	return logisticTransactionService{repo: repo}
}

// GetLogsiticFromTransactionID implements LogisticTransactionService.
func (srv logisticTransactionService) GetLogsiticFromTransactionID(ctx context.Context, transaction_id uuid.UUID) (*aggregate.LogisticTransaction, error) {
	logistic_transactions, err := srv.repo.FromTransactionID(ctx, transaction_id)
	if err != nil {
		return nil, err
	}
	return logistic_transactions, nil
}

func (srv logisticTransactionService) CreateNewTransactionWithItem(ctx context.Context, consignor entity.Person, consignee entity.Person, package_item entity.Item, partner_id uuid.UUID) (uuid.UUID, error) {
	transaction_id := uuid.New()
	new_transaction := aggregate.LogisticTransaction{
		TransactionID: transaction_id,
		Consignor:     consignor,
		Consignee:     consignee,
		Status:        "pending",
		PartnerID:     partner_id,
		Package:       package_item,
		Timestamp:     time.Now(),
		DeletedAt:     nil,
	}
	if err := srv.repo.Save(ctx, &new_transaction); err != nil {
		fmt.Println(err)
		return uuid.Nil, err
	}
	return transaction_id, nil
}
