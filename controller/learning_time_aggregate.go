package controller

import (
	"net/http"
	"strconv"

	"local.packages/service"
	// "local.packages/generated"
	"local.packages/models"

	"github.com/gin-gonic/gin"
)

type AggregateResultForm struct {
	// 総学習時間
	TotalLearningTime int `json:"total_learning_time"`
	// Output.Input割合
	AchievementPercentag models.AchievementPercentag `json:achievement_percentag`
	// 学習カテゴリ分布
	CategoryDistribution []models.CategoryDistribution `json:category_distribution`
}

func FindAggregateResult(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	learningService := service.NewLearningTimeAggregateService()
	// 総学習時間を算出する
	time, _ := learningService.CalcTotalLearningTime(userID)

	// switch {
	// case err == sql.ErrNoRows:
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":    "OK",
	// 		"input":      in,
	// 		"categories": "",
	// 	})
	// 	return
	// case err != nil:
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "ng",
	// 		"err":     err,
	// 	})
	// 	return
	// default:
	// }

	// categories, err2 := inputService.FindCategoriesBy(in.InputAchievementID)

	// switch {
	// case err2 == sql.ErrNoRows:
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":    "OK",
	// 		"input":      in,
	// 		"categories": "",
	// 	})
	// 	return
	// case err2 != nil:
	// 	fmt.Println(err2)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "ng",
	// 		"err":     err2,
	// 	})
	// 	return
	// default:
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"time":    time,
	})
}
