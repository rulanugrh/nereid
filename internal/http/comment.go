package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"github.com/rulanugrh/isonoe/internal/middleware"
	"github.com/rulanugrh/isonoe/internal/service"
)

type CommentInterface interface {
	// CreateComment creates a new comment
	CreateComment(ctx *fiber.Ctx) error
	// GetComment retrieves a comment
	GetAllComment(ctx *fiber.Ctx) error
	// DeleteComment retrieves a comment
	DeleteComment(ctx *fiber.Ctx) error
}

type comment struct {
	service service.CommentInterface
}

func NewCommentHandler(service service.CommentInterface) CommentInterface {
	return &comment{
		service: service,
	}
}
// CreateComment creates a new comment
func(c *comment) CreateComment(ctx *fiber.Ctx) error {
	var request domain.CommentRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(500).JSON(web.InternalServerError("Cannot parsing request"))
	}

	data, err := c.service.CreateComment(request)
	if err != nil {
		return ctx.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return ctx.Status(201).JSON(web.Created("Success create comment", data))
}

// GetComment retrieves a comment
func(c *comment) GetAllComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	data, err := c.service.FindAllComment(id)
	if err != nil {
		return ctx.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(web.Success("success get all comment", data))
}

// DeleteComment retrieves a comment
func(c *comment) DeleteComment(ctx *fiber.Ctx) error {

	token := ctx.Get("Authorization")
	name, err := middleware.GetUserName(token)
	if err != nil {
		return ctx.Status(401).JSON(web.Unauthorized("You must login first"))
	}

	id := ctx.Params("id")
	err = c.service.DeleteComment(id, *name)
	if err != nil {
		return ctx.Status(400).JSON(web.BadRequest(err.Error()))
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "success delete comment",
	})
}