package repository

import (
	"context"

	"github.com/sebsvt/software-prototype/transportation-software-services/aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type logisticTransactionRepositoryMongoDB struct {
	collection *mongo.Collection
}

func NewLogisticTransactionRepositoryMongoDB(collection *mongo.Collection) aggregate.LogisticTransactionRepository {
	return logisticTransactionRepositoryMongoDB{collection: collection}
}

// FromTransactionID implements aggregate.LogisticTransactionRepository.
func (repo logisticTransactionRepositoryMongoDB) FromTransactionID(ctx context.Context, transaction_id string) (*aggregate.LogisticTransaction, error) {
	// Create a filter to find the document with the specific transaction_id
	filter := bson.M{"transaction_id": transaction_id}

	var transaction aggregate.LogisticTransaction
	err := repo.collection.FindOne(ctx, filter).Decode(&transaction)
	if err != nil {
		// if errors.Is(err, mongo.ErrNoDocuments) {
		// 	// Handle the case where no document was found
		// 	return nil, nil
		// }
		// Return the error if there was an issue with the query
		return nil, err
	}

	return &transaction, nil
}

// Save implements aggregate.LogisticTransactionRepository.
func (repo logisticTransactionRepositoryMongoDB) Save(ctx context.Context, entity *aggregate.LogisticTransaction) error {
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

func (repo logisticTransactionRepositoryMongoDB) InsertMany(ctx context.Context, entities []aggregate.LogisticTransaction) error {
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
