package handler

import (
	"load_data_api/internal/database"
	"load_data_api/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Get_cases(ctx *gin.Context) {

}

func Add_case(ctx *gin.Context) {
	var body model.CreateCaseRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := database.Cases.InsertOne(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to add exercise"})
		return
	}

	bullet := model.Case{
		Id:        res.InsertedID.(primitive.ObjectID),
		Cartridge: body.Cartridge,
		Brand:     body.Brand,
		Length:    body.Length,
	}

	ctx.JSON(http.StatusCreated, bullet)
}

func Delete_case(ctx *gin.Context) {

}
