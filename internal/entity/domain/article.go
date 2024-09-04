package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID  `json:"id,omitempty" bson:"id,omitempty"`
	Title     string              `json:"title" bson:"title"`
	Content   string              `json:"content" bson:"content"`
	Banner    string              `json:"banner" bson:"banner"`
	Author    string              `json:"author" bson:"author"`
	Tags      []string            `json:"tags" bson:"tags"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}
