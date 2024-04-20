package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Primer struct {
	Id    primitive.ObjectID `json:"_id" bson:"_id"`
	Brand string             `json:"brand" bson:"brand"`
	Name  string             `json:"name" bson:"name"`
	Type  string             `json:"type" bson:"type"`
}
