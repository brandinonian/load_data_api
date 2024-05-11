package handler

import (
	"load_data_api/internal/database"
	"load_data_api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get_all_powders(ctx *gin.Context) {
	cursor, err := database.Powders.Find(ctx, bson.D{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var powders []model.Powder
	if err := cursor.All(ctx, &powders); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, powders)
}

func Add_powder(ctx *gin.Context) {
	var body model.CreatePowderRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.Powders.InsertOne(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add powder"})
		return
	}

	powder := model.Powder{
		Id:    res.InsertedID.(primitive.ObjectID),
		Brand: body.Brand,
		Name:  body.Name,
	}

	ctx.JSON(http.StatusCreated, powder)
}

func Delete_powder(ctx *gin.Context) {
	id := ctx.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := database.Powders.DeleteOne(ctx, bson.M{"_id": _id})

	if res.DeletedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "powder not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "powder deleted"})
}
