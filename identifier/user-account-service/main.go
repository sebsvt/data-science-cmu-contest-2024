package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/standardise-software/user-account-service/event"
	"github.com/standardise-software/user-account-service/logs"
)

func main() {
	// godotenv.Load()
	// db := initDB()
	// app := fiber.New()
	// api := app.Group("/api")

	// user_account_repo := repository.NewUserAccountRepositoryDB(db)
	// user_account_service := service.NewUserAccountService(user_account_repo)
	// user_account_handler := handler.NewUserAccountHandler(user_account_service)

	// api.Post("/users/create", user_account_handler.CreateNewUserAccount)
	// api.Get("/users/:user_id", user_account_handler.GetAccountFromID)

	// app.Listen(":8000")
	ch := initRabbitMQ()
	defer ch.Close()
	message_event := event.NewAccountEventHandler(ch)
	emailEvent := event.UpdatedEmail{
		Email: "new-email@example.com",
	}
	if err := message_event.Sender("email_update_queue", []byte(emailEvent.Email)); err != nil {
		logs.Error(err)
	}
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DB"),
		os.Getenv("DATABASE_SSLMODE"),
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}

func initRabbitMQ() *amqp.Channel {
	conn, err := amqp.Dial("amqp://aiselena:s4cret@localhost:5672/")
	if err != nil {
		panic(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return channel
}
