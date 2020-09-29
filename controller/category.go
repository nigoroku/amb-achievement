package controller

import (
	"fmt"

	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCategories(c *gin.Context) {

	categoryService := service.NewCategoryService()
	categories, err := categoryService.FindAll()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "NG",
			"err":     err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "OK",
		"categories": categories,
	})
}
