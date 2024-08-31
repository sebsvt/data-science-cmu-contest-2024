package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sebsvt/prototype/handlers"
	"github.com/sebsvt/prototype/middlewares"
	"github.com/sebsvt/prototype/repository"
	"github.com/sebsvt/prototype/services"
)

func main() {
	db := initDB()
	user_repo := repository.NewUserRepositoryDB(db)
	user_srv := services.NewUserService(user_repo)
	auth_srv := services.NewAuth(user_repo, []byte("my-secret"), time.Minute*2, time.Minute*4)
	user_handler := handlers.NewAuthHandler(user_srv, auth_srv)
	user_profile_repo := repository.NewUserProfileRepository(db)
	user_profile_srv := services.NewUserProfileService(user_profile_repo)
	user_profile_handler := handlers.NewUserProfileHandler(user_profile_srv)

	app := fiber.New()

	app.Use("/api/users/me", middlewares.AuthRequired(auth_srv))
	app.Use("/api/user-profiles/:user_id/create", middlewares.AuthRequired(auth_srv))

	api := app.Group("/api")

	api.Post("/auth/signup", user_handler.SignUp)
	api.Post("/auth/signin", user_handler.SignIn)
	api.Post("/auth/refresh", user_handler.RefreshToken)

	api.Get("/users/me", user_handler.GetUser)

	api.Post("/user-profiles/:user_id/create", user_profile_handler.CreateNewUserProfile)
	api.Get("/user-profiles/:user_id", user_profile_handler.GetUserProfile)
	// api.Put("/user-profiles/:user_id", user_profile_handler.UpdateUserProfile)
	app.Listen(":8000")
}

func initDB() *sqlx.DB {
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DB"),
		os.Getenv("DATABASE_SSLMODE"),
	)
	fmt.Println(dsn)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db
}
