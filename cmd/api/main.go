package main

import (
	"fmt"
	"load_data_api/internal/database"
	"load_data_api/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Init("exercise"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Database connected...")

	defer func() {
		if err := database.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	r := gin.Default()

	r.GET("/bullets/calibers", handler.Get_calibers)
	r.GET("/bullets/:cal", handler.Get_bullets_by_cal)
	r.POST("/bullets", handler.Add_bullet)
	r.DELETE("/bullets/:id", handler.Delete_bullet)

	r.Run(":8080")
}