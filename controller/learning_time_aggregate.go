package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"local.packages/service"
	// "local.packages/generated"

	"github.com/gin-gonic/gin"
)

// FindAggregateResult 学習時間の集計結果を取得する
func FindAggregateResult(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	learningService := service.NewLearningTimeAggregateService()
	// 総学習時間を算出する
	aggregateResults, err := learningService.AggregateLearningTime(userID)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message":           "OK",
			"aggregate_results": aggregateResults,
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

	c.JSON(http.StatusOK, gin.H{
		"message":           "OK",
		"aggregate_results": aggregateResults,
	})
}

// FindLearningTimeTransition 集計単位ごとの学習時間推移を取得する
func FindLearningTimeTransition(c *gin.Context) {

	userID, _ := strconv.Atoi(c.Query("user_id"))

	learningService := service.NewLearningTimeAggregateService()
	// 集計単位ごとの学習時間推移を算出する
	learningTransition, err := learningService.AggregateLearningTransition(userID)

	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusOK, gin.H{
			"message":             "OK",
			"learning_transition": learningTransition,
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

	c.JSON(http.StatusOK, gin.H{
		"message":             "OK",
		"learning_transition": learningTransition,
	})
}
