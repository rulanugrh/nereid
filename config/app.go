package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Database struct {
		Host string
		Pass string
		User string
		Name string
	}

	Server struct {
		Host string
		Port string
		Secret string
		Credential struct {
			Email string
			Pass string
		}
	}
}

var app *App

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	conf := App{}
	if err := godotenv.Load(); err != nil {
		conf.Database.Host = "localhost"
		conf.Database.User = "root"
		conf.Database.Pass = ""

		conf.Server.Host = "0.0.0.0"
		conf.Server.Port = "4000"
		conf.Server.Secret = "s3cr3t12345"

		conf.Server.Credential.Email = "admin@admin.co.id"
		conf.Server.Credential.Pass = "123456789"

		return &conf
	}

	conf.Database.Host = os.Getenv("DATABASE_HOST")
	conf.Database.User = os.Getenv("DATABASE_USER")
	conf.Database.Pass = os.Getenv("DATABASE_PASS")
	conf.Database.Name = os.Getenv("DATABASE_NAME")

	conf.Server.Host = os.Getenv("SERVER_HOST")
	conf.Server.Port = os.Getenv("SERVER_PORT")
	conf.Server.Secret = os.Getenv("SERVER_SECRET")

	conf.Server.Credential.Email = os.Getenv("SERVER_EMAIL")
	conf.Server.Credential.Pass = os.Getenv("SERVER_PASSWORD")

	return &conf
}