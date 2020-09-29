package main

import (
	"github.com/gin-contrib/cors"

	// "github.com/kzpolicy/user/controller"
	// "github.com/kzpolicy/user/middleware"
	"local.packages/controller"
	"local.packages/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --wipe mysql

func main() {
	r := gin.Default()

	// ミドルウェア
	r.Use(middleware.RecordUaAndTime)

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	// ルーティング
	aRoute := r.Group("/api/v1")
	{
		v1 := aRoute.Group("/achievement")
		{
			v1.GET("/category", controller.FindCategories)
			v1.GET("/input", controller.FindInputByUser)
			v1.POST("/input/register", controller.AddOrEditInput)

			v1.GET("/output", controller.FindOutputByUser)
			v1.POST("/output/register", controller.AddOrEditOutput)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// TODO:環境変数に設定する
	r.Run(":8083")
}
