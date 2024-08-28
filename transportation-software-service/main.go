package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sebsvt/software-prototype/transportation-software-services/repository"
	"github.com/sebsvt/software-prototype/transportation-software-services/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *mongo.Database {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	db := client.Database("aiselena")
	return db
}

func main() {
	db := initDB()
	app := fiber.New()

	logistic_transactions_repo := repository.NewLogisticTransactionRepositoryMongoDB(db.Collection("logistic_transactions"))
	logistic_transactions_srv := service.NewLogisticTransactionService(logistic_transactions_repo)

	app.Get("/:transaction_id", func(c *fiber.Ctx) error {
		transaction_id := c.Params("transaction_id")
		logistic_transaction, err := logistic_transactions_srv.GetLogsiticFromTransactionID(context.Background(), transaction_id)
		if err != nil {
			return err
		}
		return c.JSON(logistic_transaction)
	})

	app.Listen(":8000")
}
