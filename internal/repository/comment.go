package repository

import (
	"context"
	"fmt"
	"time"


	"github.com/rulanugrh/isonoe/config"
	"github.com/rulanugrh/isonoe/internal/entity/domain"
	"github.com/rulanugrh/isonoe/internal/entity/web"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentInterface interface {
	// Create comment for article
	CreateComment(req domain.CommentRequest) (*domain.Comment, error)
	// Get comment by article
	GetCommentByArticle(articleID string) (*[]domain.Comment, error)
	// Delete comment by id 
	DeleteComment(commentID string, author string) error
}

type comment struct {
	coll *mongo.Collection
}

func NewCommentRepository(conn *config.Connection, conf *config.App) CommentInterface {
	return &comment{
		coll: conn.DB.Database(conf.Database.Name).Collection("comments"),
	}
}

// Create comment for article
func(c *comment) CreateComment(req domain.CommentRequest) (*domain.Comment, error) {
	// Create new variable for decode response
	var response domain.Comment
	year, month, day := time.Now().Date()
	tm := fmt.Sprintf("%d %s %d", day, month, year)

	// create context with timeout with 10 second
	ctx, timeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer timeout()

	// parsing article id from hex
	articleID, err := primitive.ObjectIDFromHex(req.ArticleID)
	if err != nil {
		return nil, web.ErrorLog("Parsing ID from Hex")
	}

	request := domain.Comment {
		Content: req.Content,
		ArticleID: articleID,
		CreatedAt: tm,
		UpdatedAt: tm,
		Author: req.Author,
	}

	inserted, err := c.coll.InsertOne(ctx, &request)
	if err != nil {
		return nil, web.ErrorLog("Cannot insert into DB")
	}

	err = c.coll.FindOne(ctx, bson.M{"_id": inserted.InsertedID}).Decode(&response)
	if err != nil {
		return nil, web.ErrorLog("Cannot find data with this id")
	}

	return &response, nil
}

// Get comment by article
func(c *comment) GetCommentByArticle(articleID string) (*[]domain.Comment, error) {
	// create default response 
	var response []domain.Comment

	// create context with timeout with 10 second
	ctx, timeout := context.WithTimeout(context.Background(), 10 * time.Second)
	defer timeout()

	// parsing id from hex
	id, err := primitive.ObjectIDFromHex(articleID)
	if err != nil {
		return nil, web.ErrorLog("Parsing ID from Hex")
	}

	rows, err := c.coll.Find(ctx, bson.M{"article_id": id })
	if err != nil {
		return nil, web.ErrorLog("Cannot find comment with this article ID")
	}
	defer rows.Close(ctx)

	for rows.Next(ctx) {
		var comment domain.Comment
		if errs := rows.Decode(&comment); errs != nil {
			return nil, web.ErrorLog("Cannot decode to domain logic")
		}

		response = append(response, comment)
	}

	return &response, nil
}

// Delete comment by id 
func(c *comment) DeleteComment(commentID string, author string) error {
	// create context with timeout with 10 second
	ctx, timeout := context.WithTimeout(context.Background(), 10 * time.Second)
	defer timeout()

	// parsing id from string
	id, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return web.ErrorLog("Parsing ID from Hex")
	}

	_, err = c.coll.DeleteOne(ctx, bson.M{"_id": id, "author": author})
	if err != nil {
		return web.ErrorLog("Cannot delete this comment")
	}

	return nil
}