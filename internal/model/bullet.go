package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bullet struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id"`
	Cal         string             `json:"cal" bson:"cal"`
	Diam        float32            `json:"diam" bson:"diam"`
	Weight      float32            `json:"weight" bson:"weight"`
	Brand       string             `json:"brand" bson:"brand"`
	Name        string             `json:"name" bson:"name"`
	MaxVelocity int                `json:"maxV" bson:"maxV"`
	G1          float32            `json:"g1" bson:"g1"`
	G7          float32            `json:"g7" bson:"g7"`
}

type CreateBulletRequest struct {
	Cal         string  `json:"cal" bson:"cal"`
	Diam        float32 `json:"diam" bson:"diam"`
	Weight      float32 `json:"weight" bson:"weight"`
	Brand       string  `json:"brand" bson:"brand"`
	Name        string  `json:"name" bson:"name"`
	MaxVelocity int     `json:"maxV" bson:"maxV"`
	G1          float32 `json:"g1" bson:"g1"`
	G7          float32 `json:"g7" bson:"g7"`
}
