package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sebsvt/software-prototype/transportation-software-services/entity"
	"github.com/sebsvt/software-prototype/transportation-software-services/valueobject"
)

var (
	ErrCannotAcceptTransaction = errors.New("cannot accept this transaction")
	ErrMemberHaveNoPermission  = errors.New("you have no permission")
)

type LogisticTransactionAggregate struct {
	TransactionID string             `bson:"transaction_id" json:"transaction_id"` // Unique identifier for the transaction
	Consignor     entity.Person      `bson:"consignor" json:"consignor"`           // Sender's details
	Consignee     entity.Person      `bson:"consignee" json:"consignee"`           // Recipient's details
	Package       entity.Item        `bson:"package_item" json:"package"`          // Details of the package
	Status        valueobject.Status `bson:"status" json:"status"`                 // Status of the transaction (e.g., "In Transit")
	PartnerID     string             `bson:"partner_id" json:"partner_id"`         // ID for the partner (SaaS identifier)
	Timestamp     time.Time          `bson:"time_stamp" json:"timestamp"`          // Date and time of the transaction
	DeletedAt     *time.Time         `bson:"deleted_at" json:"deleted_at"`         // Timestamp for soft deletion (optional)
}

type LogisticTransactionRepository interface {
	FromTransactionID(ctx context.Context, transaction_id string) (*LogisticTransactionAggregate, error)
	FromID(ctx context.Context, id string) (*LogisticTransactionAggregate, error)
	Save(ctx context.Context, entity *LogisticTransactionAggregate) error
	InsertMany(ctx context.Context, entities []LogisticTransactionAggregate) error
}

func NewLogisticTransactionAggregate(consignor entity.Person, consignee entity.Person, package_item entity.Item, partner_id string) (LogisticTransactionAggregate, error) {
	return LogisticTransactionAggregate{
		TransactionID: uuid.New().String(),
		Consignor:     consignor,
		Consignee:     consignee,
		Package:       package_item,
		Status:        valueobject.Pending,
		Timestamp:     time.Now(),
		PartnerID:     partner_id,
	}, nil
}

func (transaction *LogisticTransactionAggregate) AccepetTransaction(member Member) error {
	if member.Role == Guess {
		return ErrMemberHaveNoPermission
	}
	if transaction.Status != valueobject.Pending {
		return ErrCannotAcceptTransaction
	}
	transaction.Status = valueobject.Accepted
	return nil
}
