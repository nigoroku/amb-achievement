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

type OutputForm struct {
	CategoryIds []string                 `json:"category_ids"`
	Achievement models.OutputAchievement `json:"achievement"`
}

func FindOutputByUser(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	outputService := service.NewOutputService()
	ou, err := outputService.FindByUser(userID)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message":    "OK",
			"output":     ou,
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

	categories, err2 := outputService.FindCategoriesBy(ou.OutputAchievementID)

	switch {
	case err2 == sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message":    "OK",
			"output":     ou,
			"categories": "",
		})
		return
	case err2 != nil:
		fmt.Println(err2)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err2,
		})
	default:
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "OK",
		"output":     ou,
		"categories": categories,
	})
}

func AddOrEditOutput(c *gin.Context) {

	var outputForm OutputForm
	c.ShouldBindJSON(&outputForm)

	outputService := service.NewOutputService()
	out, err := outputService.FindByUser(outputForm.Achievement.UserID)

	var err2 error
	switch {
	case err == sql.ErrNoRows:
		// まだ本日のoutputが登録されていない場合
		err2 = outputService.Create(outputForm.Achievement, outputForm.CategoryIds)
	case err != nil:
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ng",
			"err":     err,
		})
		return
	default:
		// 既に登録されている場合
		err2 = outputService.Update(outputForm.Achievement, outputForm.CategoryIds, out.OutputAchievementID)
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
		"output":  out,
	})
}
