package handler

import (
	"load_data_api/internal/database"
	"load_data_api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get_all_cases(ctx *gin.Context) {
	cursor, err := database.Cases.Find(ctx, bson.D{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var cases []model.Case
	if err := cursor.All(ctx, &cases); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, cases)
}

func Add_case(ctx *gin.Context) {
	var body model.CreateCaseRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.Cases.InsertOne(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add case"})
		return
	}

	casing := model.Case{
		Id:        res.InsertedID.(primitive.ObjectID),
		Cartridge: body.Cartridge,
		Brand:     body.Brand,
		Length:    body.Length,
	}

	ctx.JSON(http.StatusCreated, casing)
}

func Delete_case(ctx *gin.Context) {
	id := ctx.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := database.Cases.DeleteOne(ctx, bson.M{"_id": _id})

	if res.DeletedCount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "case not found"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "case deleted"})
}
