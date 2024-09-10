package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rulanugrh/isonoe/config"
	handler "github.com/rulanugrh/isonoe/internal/http"
	"github.com/rulanugrh/isonoe/internal/repository"
	"github.com/rulanugrh/isonoe/internal/service"
)

func main() {
	conf := config.GetConfig()
	// Create config Connection Instance
	conn := config.InitialDB(conf)

	// Inject repository layer
	userRepo := repository.NewUserRepository(conn, conf)
	articleRepo := repository.NewArticleRepository(conn, conf)

	// Inject repository layer into service layer
	userService := service.NewUserService(userRepo)
	articleService := service.NewArticleService(articleRepo)

	// Inject to handler layer
	userHandler := handler.NewUserHandler(userService)
	articleHandler := handler.NewArticleHandler(articleService)

	router(conf, articleHandler, userHandler)
}

func router(conf *config.App, article handler.ArticleInterface, user handler.UserInterface) {
	f := fiber.New()
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodDelete,
			fiber.MethodPut,
			fiber.MethodOptions,
		}, ","),
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	file, err := os.OpenFile("./sys.log", os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	f.Use(logger.New(logger.Config{
		Format:   "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeZone: "Asia/Jakarta",
		Output:   file,
	}))

	f.Get("/", article.GetAll)
	f.Get("/find/:id", article.GetById)
	f.Post("/article", article.Create)
	f.Delete("/article/:id", article.Delete)

	f.Post("/register", user.Register)
	f.Post("/login", user.Login)
	f.Get("/me", user.GetMe)

	server := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	log.Fatal(f.Listen(server))
}