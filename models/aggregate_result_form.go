package models

// AggregateResultForm 学習時間の集計結果を格納する
type AggregateResultForm struct {
	// 総学習時間
	TotalLearningTime int `json:"total_learning_time"`
	// Output.Input割合
	AchievementPercentag AchievementPercentag `json:achievement_percentag`
	// 学習カテゴリ分布
	CategoryDistribution []*CategoryDistribution `json:category_distribution`
}

// NewAggregateResultForm コンストラクタ
func NewAggregateResultForm(totalLearningTime int,
	achievementPercentag AchievementPercentag,
	categoryDistribution []*CategoryDistribution) *AggregateResultForm {
	c := new(AggregateResultForm)
	c.TotalLearningTime = totalLearningTime
	c.AchievementPercentag = achievementPercentag
	c.CategoryDistribution = categoryDistribution
	return c
}
