package service

import (

	// "github.com/kzpolicy/user/generated"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/generated"
	"local.packages/models"
	"local.packages/utils"

	"context"
)

const (
	YearFormat  = string("%Y")
	MonthFormat = string("%Y/%m")
	DaysFormat  = string("%Y/%m/%d")
	YearUnit    = string("year")
	MonthUnit   = string("month")
	DaysUnit    = string("days")
)

// LearningTimeAggregateService 学習時間算出サービス
type LearningTimeAggregateService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

// NewLearningTimeAggregateService 学習時間算出サービスコンストラクタ
func NewLearningTimeAggregateService() *LearningTimeAggregateService {
	ctx := context.Background()
	db := boil.GetContextDB()

	return &LearningTimeAggregateService{ctx, db}
}

// AggregateLearningTime 学習時間を集計する
func (lt *LearningTimeAggregateService) AggregateLearningTime(userID int) (*models.AggregateResultForm, error) {

	inputTimeCategories, err1 := lt.CalcTotalForInputCategory(userID)
	outputTimeCategories, err2 := lt.CalcTotalForOutputCategory(userID)

	if err1 != nil {
		return new(models.AggregateResultForm), err1
	}

	if err2 != nil {
		return new(models.AggregateResultForm), err2
	}

	// 総学習時間
	outputTime := 0
	for _, input := range inputTimeCategories {
		outputTime += input.TotalTime
	}
	inputTime := 0
	for _, output := range outputTimeCategories {
		inputTime += output.TotalTime
	}
	totalTime := outputTime + inputTime

	// Output.Input割合
	outputPercentag := float64(0)
	inputPercentag := 0
	if outputTime != 0 && inputTime != 0 {
		outputPercentag = float64(outputTime) / float64(totalTime) * 100
		inputPercentag = 100 - int(outputPercentag)
	}
	achievementPercentag := models.NewAchievementPercentag(int(outputPercentag), inputPercentag, outputTime, inputTime)

	// カテゴリごとの学習時間
	var categoryDistribution []*models.CategoryDistribution
	// inputに含まれていてoutputにも含まれているカテゴリの学種時間を設定
	for _, ic := range inputTimeCategories {
		for _, oc := range outputTimeCategories {
			if oc.CategoryName == ic.CategoryName {
				ic.TotalTime += oc.TotalTime
			}
		}
		categoryDistribution = append(categoryDistribution, models.NewCategoryDistribution(ic.CategoryName, ic.TotalTime))
	}
	// outputにのみ含まれているカテゴリの学種時間を設定
	for _, oc := range outputTimeCategories {
		isContain := false
		for _, ic := range inputTimeCategories {
			if ic.CategoryName == oc.CategoryName {
				isContain = true
			}
		}
		if isContain == false {
			// outputにしかなければ設定
			categoryDistribution = append(categoryDistribution, models.NewCategoryDistribution(oc.CategoryName, oc.TotalTime))
		}
	}

	return models.NewAggregateResultForm(totalTime, *achievementPercentag, categoryDistribution), nil
}

// AggregateLearningTransition 集計単位ごとの学習時間推移を取得する
func (lt *LearningTimeAggregateService) AggregateLearningTransition(userID int) (*models.LearningTransitionForm, error) {

	yearLearningTransition, err1 := lt.calcLearningTransitionByUnit(userID, YearFormat)
	monthLearningTransition, err2 := lt.calcLearningTransitionByUnit(userID, MonthFormat)
	daysLearningTransition, err3 := lt.calcLearningTransitionByUnit(userID, DaysFormat)

	if err1 != nil {
		return new(models.LearningTransitionForm), err1
	}

	if err2 != nil {
		return new(models.LearningTransitionForm), err2
	}

	if err3 != nil {
		return new(models.LearningTransitionForm), err3
	}

	// 年ごとの集計
	years := models.LearningTransitionSlice{}
	for _, l := range createLabelsByUnit(YearUnit) {
		ltm := models.NewLearningTransition(l, 0)
		for _, y := range yearLearningTransition {
			if l == y.Label {
				ltm.Time = y.Time
			}
		}
		years = append(years, ltm)
	}
	sort.Sort(years)

	// 月ごとの集計
	months := models.LearningTransitionSlice{}
	for _, l := range createLabelsByUnit(MonthUnit) {
		ltm := models.NewLearningTransition(l, 0)
		for _, m := range monthLearningTransition {
			if l == m.Label {
				ltm.Time = m.Time
			}
		}
		months = append(months, ltm)
	}
	sort.Sort(months)

	// 週ごとの集計
	days := models.LearningTransitionSlice{}
	for _, l := range createLabelsByUnit(DaysUnit) {
		ltm := models.NewLearningTransition(l, 0)
		for _, d := range daysLearningTransition {
			if l == d.Label {
				ltm.Time = d.Time
			}
		}
		days = append(days, ltm)
	}
	sort.Sort(days)

	return models.NewLearningTransitionForm(years, months, days), nil
}

// CalcTotalForInputCategory インプットカテゴリごとの総学習時間を算出する
func (lt *LearningTimeAggregateService) CalcTotalForInputCategory(userID int) ([]models.CategoryDistribution, error) {

	var categoryDistributions []models.CategoryDistribution
	err := generated.NewQuery(
		qm.Select("mc.name as category_name", "SUM(ia.input_time) as total_time"),
		qm.From("users u"),
		qm.InnerJoin("input_achievements ia ON u.user_id = ia.user_id"),
		qm.InnerJoin("input_achievement_tags iat ON ia.input_achievement_id = iat.input_achievement_id"),
		qm.InnerJoin("m_categories mc ON iat.category_id = mc.category_id"),
		qm.Where("u.user_id=?", userID),
		qm.GroupBy("mc.name"),
		qm.OrderBy("mc.name DESC"),
	).Bind(lt.ctx, lt.db, &categoryDistributions)

	return categoryDistributions, err
}

// CalcTotalForOutputCategory インプットカテゴリごとの総学習時間を算出する
func (lt *LearningTimeAggregateService) CalcTotalForOutputCategory(userID int) ([]models.CategoryDistribution, error) {

	var categoryDistributions []models.CategoryDistribution
	err := generated.NewQuery(
		qm.Select("mc.name as category_name", "SUM(oa.output_time) as total_time"),
		qm.From("users u"),
		qm.InnerJoin("output_achievements oa ON u.user_id = oa.user_id"),
		qm.InnerJoin("output_achievement_tags oat ON oa.output_achievement_id = oat.output_achievement_id"),
		qm.InnerJoin("m_categories mc ON oat.category_id = mc.category_id"),
		qm.Where("u.user_id=?", userID),
		qm.GroupBy("mc.name"),
		qm.OrderBy("mc.name DESC"),
	).Bind(lt.ctx, lt.db, &categoryDistributions)

	return categoryDistributions, err
}

// CalcLearningTransitionByUnit 学習時間の推移を集計単位ごとに算出する
func (lt *LearningTimeAggregateService) calcLearningTransitionByUnit(userID int, dateFormat string) ([]*models.LearningTransition, error) {

	var learningTransition []*models.LearningTransition
	err := generated.NewQuery(
		qm.Select("SUM(ia.input_time) as time", "DATE_FORMAT(ia.created_at,'"+dateFormat+"') as label"),
		qm.From("users u"),
		qm.InnerJoin("input_achievements ia ON u.user_id = ia.user_id"),
		qm.Where("u.user_id=?", userID),
		qm.GroupBy("DATE_FORMAT(ia.created_at, '"+dateFormat+"')"),
		qm.OrderBy("label DESC"),
	).Bind(lt.ctx, lt.db, &learningTransition)

	return learningTransition, err
}

// createLabelsByUnit 集計単位ごとのグラフのラベルを生成する
func createLabelsByUnit(unit string) []string {
	now := time.Now()
	var labels []string
	switch unit {
	case "year":
		for i, _ := range make([]int, 12) {
			y := now.Year() - i
			labels = append(labels, strconv.Itoa(y))
		}
	case "month":
		for i, _ := range make([]int, 12) {
			m := utils.AddMonth(now, i)
			labels = append(labels, m.Format("2006/01"))
		}
	case "days":
		for i, _ := range make([]int, 7) {
			d := now.AddDate(0, 0, -i)
			// wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
			// w := wdays[d.Weekday()]
			labels = append(labels, d.Format("2006/01/02"))
		}
	default:
		fmt.Println("the two strings are NOT equal")
	}
	return labels
}
