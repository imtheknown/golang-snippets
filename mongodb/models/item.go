package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Item is Exported
type Item struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	CoverImage  string             `bson:"coverimage" json:"coverimage"`
	Description string             `bson:"description" json:"description"`
}
