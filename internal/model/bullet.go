package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bullet struct {
	Id     primitive.ObjectID `json:"_id" bson:"_id"`
	Cal    string             `json:"cal" bson:"cal"`
	Diam   float32            `json:"diam" bson:"diam"`
	Weight float32            `json:"weight" bson:"weight"`
	Brand  string             `json:"brand" bson:"brand"`
	Name   string             `json:"name" bson:"name"`
}

type CreateBulletRequest struct {
	Cal    string  `json:"cal" bson:"cal"`
	Diam   float32 `json:"diam" bson:"diam"`
	Weight float32 `json:"weight" bson:"weight"`
	Brand  string  `json:"brand" bson:"brand"`
	Name   string  `json:"name" bson:"name"`
}
