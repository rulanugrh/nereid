package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/rulanugrh/isonoe/config"
	handler "github.com/rulanugrh/isonoe/internal/http"
	"github.com/rulanugrh/isonoe/internal/middleware"
	"github.com/rulanugrh/isonoe/internal/repository"
	"github.com/rulanugrh/isonoe/internal/service"
)

func main() {
	conf := config.GetConfig()
	// Create config Connection Instance
	conn := config.InitialDB(conf)
	conn.ToConnectMongoDB()

	// Inject repository layer
	userRepo := repository.NewUserRepository(conn, conf)
	articleRepo := repository.NewArticleRepository(conn, conf)
	commentRepo := repository.NewCommentRepository(conn, conf)

	// Inject repository layer into service layer
	userService := service.NewUserService(userRepo)
	articleService := service.NewArticleService(articleRepo)
	commentService := service.NewCommentService(commentRepo)

	// Inject to handler layer
	userHandler := handler.NewUserHandler(userService)
	articleHandler := handler.NewArticleHandler(articleService)
	commentHandler := handler.NewCommentHandler(commentService)

	router(conf, articleHandler, userHandler, commentHandler)
}

func router(conf *config.App, article handler.ArticleInterface, user handler.UserInterface, comment handler.CommentInterface) {
	engine := html.New("./views", ".html")

	f := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base.layout",
	})

	f.Static("/static", "./views/static")
	f.Static("/partials", "./views/partials")

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
	f.Get("/article/:id", article.GetById)
	f.Post("/login", user.Login)
	f.Post("/register", user.Register)
	f.Get("/comment/:id", comment.GetAllComment)

	f.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(conf.Server.Secret),
		},
		ErrorHandler: middleware.ErrorHandlerJWT,
		TokenLookup:  "header:Authorization",
	}))

	f.Post("/article", article.Create)
	f.Delete("/article/:id", article.Delete)
	f.Get("/me", user.GetMe)
	f.Post("/comment/", comment.CreateComment)
	f.Delete("/comment/:id", comment.DeleteComment)

	server := fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)
	log.Fatal(f.Listen(server))
}
