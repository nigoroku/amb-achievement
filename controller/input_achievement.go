package controller

import (
	"database/sql"
	"fmt"
	"strconv"

	"local.packages/models"
	"local.packages/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type InputForm struct {
	CategoryIds []string                `json:"category_ids"`
	Achievement models.InputAchievement `json:"achievement"`
}

func FindInputByUser(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	inputService := service.NewInputService()
	in, err := inputService.FindByUser(userID)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message":    "OK",
			"input":      in,
			"categories": "",
		})
		return
	case err != nil:
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	default:
	}

	categories, err2 := inputService.FindCategoriesBy(in.InputAchievementID)

	switch {
	case err2 == sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message":    "OK",
			"input":      in,
			"categories": "",
		})
		return
	case err2 != nil:
		fmt.Println(err2)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err2,
		})
		return
	default:
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "OK",
		"input":      in,
		"categories": categories,
	})
}

func AddOrEditInput(c *gin.Context) {

	var inputForm InputForm
	c.ShouldBindJSON(&inputForm)

	inputService := service.NewInputService()
	in, err := inputService.FindByUser(inputForm.Achievement.UserID)

	var err2 error
	switch {
	case err == sql.ErrNoRows:
		// まだ本日のoutputが登録されていない場合
		err2 = inputService.Create(inputForm.Achievement, inputForm.CategoryIds)
	case err != nil:
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	default:
		// 既に登録されている場合
		err2 = inputService.Update(inputForm.Achievement, inputForm.CategoryIds, in.InputAchievementID)
	}

	if err2 != nil {
		fmt.Println(err2)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err2,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"input":   in,
	})
}
