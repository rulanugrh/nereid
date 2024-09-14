package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Content     string             `json:"content" bson:"content"`
	Banner      string             `json:"banner" bson:"banner"`
	Author      string             `json:"author" bson:"author"`
	Tags        []string           `json:"tags" bson:"tags"`
	CreatedAt   string             `json:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	Conclusion  string             `json:"conclusion" form:"conclusion"`
}
