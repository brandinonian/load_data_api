package main

import (
	"fmt"
	"load_data_api/internal/database"
	"load_data_api/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Init("load_data"); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Database connected...")

	defer func() {
		if err := database.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	r := gin.Default()

	r.GET("/calibers", handler.Get_all_calibers)
	r.GET("/bullets/:cal", handler.Get_bullets_by_cal)
	r.GET("/bullets/id/:id", handler.Get_bullet_by_id)
	r.POST("/bullets", handler.Add_bullet)
	r.DELETE("/bullets/:id", handler.Delete_bullet)

	r.GET("/cases", handler.Get_all_cases)
	r.POST("/cases", handler.Add_case)
	r.DELETE("/cases/:id", handler.Delete_case)

	r.GET("/powders", handler.Get_all_powders)
	r.POST("/powders", handler.Add_powder)
	r.DELETE("/powders/:id", handler.Delete_powder)

	r.GET("/primers", handler.Get_all_primers)
	r.POST("/primers", handler.Add_primer)
	r.DELETE("/primers/:id", handler.Delete_primer)

	r.Run(":8080")
}
