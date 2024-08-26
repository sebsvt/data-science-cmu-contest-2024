package aggregate

import (
	"context"
	"time"

	"github.com/sebsvt/software-prototype/transportation-software-services/entity"
)

type LogisticTransaction struct {
	TransactionID string        `bson:"transaction_id"` // Unique identifier for the transaction
	Consignor     entity.Person `bson:"consignor"`      // Sender's details
	Consignee     entity.Person `bson:"consignee"`      // Recipient's details
	Package       entity.Item   `bson:"package_item"`   // Details of the package
	Status        string        `bson:"status"`         // Status of the transaction (e.g., "In Transit")
	PartnerID     string        `bson:"partner_id"`     // ID for the partner (SaaS identifier)
	Timestamp     time.Time     `bson:"time_stamp"`     // Date and time of the transaction
	DeletedAt     *time.Time    `bson:"deleted_at"`     // Timestamp for soft deletion (optional)
}

type LogisticTransactionRepository interface {
	FromTransactionID(ctx context.Context, transaction_id string) (*LogisticTransaction, error)
	Save(ctx context.Context, entity *LogisticTransaction) error
	InsertMany(ctx context.Context, entities []LogisticTransaction) error
}
