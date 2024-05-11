package handler

import (
	"load_data_api/internal/database"
	"load_data_api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get_all_primers(ctx *gin.Context) {
	cursor, err := database.Primers.Find(ctx, bson.D{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var primers []model.Primer
	if err := cursor.All(ctx, &primers); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, primers)
}

func Add_primer(ctx *gin.Context) {
	var body model.CreatePrimerRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.Primers.InsertOne(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add primer"})
		return
	}

	primer := model.Primer{
		Id:    res.InsertedID.(primitive.ObjectID),
		Brand: body.Brand,
		Name:  body.Name,
		Type:  body.Type,
	}

	ctx.JSON(http.StatusCreated, primer)
}

func Delete_primer(ctx *gin.Context) {
	id := ctx.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := database.Primers.DeleteOne(ctx, bson.M{"_id": _id})

	if res.DeletedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "primer not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "primer deleted"})
}
