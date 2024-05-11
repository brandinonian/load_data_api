package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Powder struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id"`
	Brand string             `json:"brand" bson:"brand"`
	Name  string             `json:"name" bson:"name"`
}

type CreatePowderRequest struct {
	Brand string `json:"brand" bson:"brand"`
	Name  string `json:"name" bson:"name"`
}
