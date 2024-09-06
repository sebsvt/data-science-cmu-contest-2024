package repository

import (
	"context"

	"github.com/sebsvt/software-prototype/transportation-software-services/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type logisticTransactionRepositoryMongoDB struct {
	collection *mongo.Collection
}

func NewLogisticTransactionRepositoryMongoDB(collection *mongo.Collection) domain.LogisticTransactionRepository {
	return logisticTransactionRepositoryMongoDB{collection: collection}
}

// FromTransactionID implements domain.LogisticTransactionRepository.
func (repo logisticTransactionRepositoryMongoDB) FromTransactionID(ctx context.Context, transaction_id string) (*domain.LogisticTransactionAggregate, error) {
	// Create a filter to find the document with the specific transaction_id
	filter := bson.M{"transaction_id": transaction_id}

	var transaction domain.LogisticTransactionAggregate
	err := repo.collection.FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// Save implements domain.LogisticTransactionRepository.
func (repo logisticTransactionRepositoryMongoDB) Save(ctx context.Context, entity *domain.LogisticTransactionAggregate) error {
	// Create a filter to find the document with the specific transaction_id
	filter := bson.M{"transaction_id": entity.TransactionID}

	// Define update options: upsert true to insert a new document if one doesn't exist
	opts := options.Update().SetUpsert(true)

	// Convert the LogisticTransaction entity to a BSON map for MongoDB
	update := bson.M{"$set": entity}

	_, err := repo.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		// Handle the error if the update fails
		return err
	}

	return nil
}

func (repo logisticTransactionRepositoryMongoDB) InsertMany(ctx context.Context, entities []domain.LogisticTransactionAggregate) error {
	// Convert the slice of LogisticTransaction entities into a slice of interface{}
	var docs []interface{}
	for _, entity := range entities {
		docs = append(docs, entity)
	}

	// Perform the insert many operation
	_, err := repo.collection.InsertMany(ctx, docs)
	if err != nil {
		// Handle the error if the insert fails
		return err
	}

	// Return the IDs of the inserted documents
	return nil
}

func (repo logisticTransactionRepositoryMongoDB) FromID(ctx context.Context, id string) (*domain.LogisticTransactionAggregate, error) {
	filter := bson.M{"_id": id}

	var transaction domain.LogisticTransactionAggregate
	err := repo.collection.FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
