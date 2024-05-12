package handler

import (
	"load_data_api/internal/database"
	"load_data_api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// step 1
// use this endpoint to narrow down the bullet search
func Get_all_calibers(ctx *gin.Context) {
	calibers, err := database.Bullets.Distinct(ctx, "cal", bson.D{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, calibers)
}

// step 2
// use a caliber from above here
func Get_bullets_by_cal(ctx *gin.Context) {
	caliber := ctx.Param("cal")

	cursor, err := database.Bullets.Find(ctx, bson.M{"cal": caliber})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var bullets []model.Bullet
	if err := cursor.All(ctx, &bullets); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, bullets)
}

func Get_bullet_by_id(ctx *gin.Context) {
	id := ctx.Param("id")

	bullet := database.Bullets.FindOne(ctx, bson.M{"_id": id})

	if bullet == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bullet not found"})
		return
	}

	ctx.JSON(http.StatusOK, bullet)
}

func Add_bullet(ctx *gin.Context) {
	var body model.CreateBulletRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.Bullets.InsertOne(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add bullet"})
		return
	}

	bullet := model.Bullet{
		Id:     res.InsertedID.(primitive.ObjectID),
		Cal:    body.Cal,
		Diam:   body.Diam,
		Weight: body.Weight,
		Name:   body.Name,
		Brand:  body.Brand,
	}

	ctx.JSON(http.StatusCreated, bullet)
}

func Delete_bullet(ctx *gin.Context) {
	id := ctx.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := database.Bullets.DeleteOne(ctx, bson.M{"_id": _id})

	if res.DeletedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bullet not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "bullet deleted"})
}
