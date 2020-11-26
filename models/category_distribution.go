package models

// AchievementPercentag OutputとInputの総学習時間の割合を格納する
type CategoryDistribution struct {
	CategoryName string `json:"category_name"`
	TotalTime    int    `json:"total_time"`
}
