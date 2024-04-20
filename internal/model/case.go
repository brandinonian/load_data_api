package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Case struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	Cartridge string             `json:"cartridge" bson:"cartridge"`
	Brand     string             `json:"brand" bson:"brand"`
	Length    float32            `json:"length" bson:"length"`
}
