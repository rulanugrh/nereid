package handler

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"github.com/rulanugrh/isonoe/internal/middleware"
	"github.com/rulanugrh/isonoe/internal/service"
	"golang.org/x/crypto/bcrypt"
)
type UserInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetMe(c *fiber.Ctx) error
}

type user struct {
	service service.UserInterface
}

func NewUserHandler(service service.UserInterface) UserInterface {
	return &user{ service: service}
}

func(u *user) Register(c *fiber.Ctx) error {
	var request domain.UserRegister
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(500).JSON(web.InternalServerError("[ERROR] - Parsing body request"))
	}

	data, err := u.service.Register(request)
	if err != nil {
		return c.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return c.Status(201).JSON(web.Created("success create account", data))
}

func(u *user) Login(c *fiber.Ctx) error {
	var request domain.UserLogin
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(500).JSON(web.InternalServerError("ERROR - Parsing body request"))
	}

	data, err := u.service.Login(request)
	if err != nil {
		return c.Status(400).JSON(web.BadRequest(err.Error()))
	}

	verify := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password))
	if verify != nil {
		return c.Status(401).JSON(web.Unauthorized("Your password is not matched"))
	}

	response := web.GetAccount {
		ID: data.ID,
		Name: data.Name,
		Email: data.Email,
	}

	token, err := middleware.CreateToken(response)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	c.Response().Header.Add("Authorization", *token)
	return c.Status(200).JSON(web.Success("Success Login", token))

}

func(u *user) GetMe(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	email, err := middleware.GetUserEmail(token)
	if err != nil {
		return c.Status(401).JSON(web.Unauthorized("You must login first"))
	}

	data, err := u.service.GetMe(*email)
	if err != nil {
		return c.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return c.Status(200).JSON(web.Success("Success Get Account", data))
}