package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sebsvt/software-prototype/transportation-software-services/handler"
	"github.com/sebsvt/software-prototype/transportation-software-services/middlewares"
	"github.com/sebsvt/software-prototype/transportation-software-services/repository"
	"github.com/sebsvt/software-prototype/transportation-software-services/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() *mongo.Database {
	dsn := fmt.Sprintf("mongodb://%v:%v@%v:%v/", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"))
	clientOptions := options.Client().ApplyURI(dsn)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalln(err)
	}
	db := client.Database(os.Getenv("DATABASE_DB"))
	return db
}

func main() {
	godotenv.Load()
	middlewares.LoadEnv()
	db := initDB()
	app := fiber.New()

	logistic_transactions_repo := repository.NewLogisticTransactionRepositoryMongoDB(db.Collection("logistic_transactions"))
	logistic_transactions_srv := service.NewLogisticTransactionService(logistic_transactions_repo)
	logistic_transactions_handler := handler.NewLogisticTransactionHandler(logistic_transactions_srv)

	api := app.Group("/api")

	api.Get("/logistic_transaction/:transaction_id", logistic_transactions_handler.GetTransaction)
	api.Get("/logistic_transaction/:partner_id", logistic_transactions_handler.GetTransaction)
	api.Get("/logistic_transaction/id/:id", logistic_transactions_handler.GetTransactionFromID)
	api.Post("/logistic_transaction/create", middlewares.JWTMiddleware, logistic_transactions_handler.CreateNewTransactionWithItem)

	app.Listen(":8080")
}
