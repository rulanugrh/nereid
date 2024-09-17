package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ArticleID primitive.ObjectID `json:"article_id" bson:"article_id"`
	Content   string             `json:"content" bson:"content" form:"content"`
	Author    string             `json:"author" bson:"author"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
}

type CommentRequest struct {
	ArticleID string `json:"article_id" bson:"article_id" validate:"required"`
	Content   string `json:"content" bson:"content" form:"content" validate:"required"`
	Author    string `json:"author" bson:"author"`
}
