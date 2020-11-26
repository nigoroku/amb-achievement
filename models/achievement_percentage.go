package models

// AchievementPercentag OutputとInputの総学習時間の割合を格納する
type AchievementPercentag struct {
	OutputPercentag float64 `json:"output_percentag"`
	InputPercentag  float64 `json:"input_percentag"`
	OutputTotalTime int     `json:"output_total_time"`
	InputTotalTime  int     `json:"input_total_time"`
}
