package main

import (
	"log"
	
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pandudpn/blog/controller/login/handler"
	"github.com/pandudpn/blog/dbc"
	"github.com/pandudpn/blog/repository/cache"
	usersrepo "github.com/pandudpn/blog/repository/sql/users_repo"
	"github.com/pandudpn/blog/usecase/login"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error load .env file %s", err)
	}
	
	db, err := dbc.NewConnectionSql()
	if err != nil {
		log.Fatalf("error connect to database %s", err)
	}
	
	redis, err := dbc.NewConnectionRedis()
	if err != nil {
		log.Fatalf("error connect to redis %s", err)
	}
	
	userRepo := usersrepo.NewUsersRepo(db)
	cacheRepo := cache.New(redis)
	
	loginUc := login.NewLoginUseCase(userRepo, cacheRepo)
	
	// validator
	valid := validator.New()
	
	app := fiber.New()
	handler.NewHandlerLogin(valid, loginUc).Route(app)
	
	log.Fatalln(app.Listen(":8888"))
}
