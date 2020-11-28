package models

// AchievementPercentag OutputとInputの総学習時間の割合を格納する
type AchievementPercentag struct {
	OutputPercentag int `json:"output_percentag"`
	InputPercentag  int `json:"input_percentag"`
	OutputTotalTime int `json:"output_total_time"`
	InputTotalTime  int `json:"input_total_time"`
}

// NewAchievementPercentag コンストラクタ
func NewAchievementPercentag(outputPercentag int, inputPercentag int, outputTotalTime int, inputTotalTime int) *AchievementPercentag {
	a := new(AchievementPercentag)
	a.OutputPercentag = outputPercentag
	a.InputPercentag = inputPercentag
	a.OutputTotalTime = outputTotalTime
	a.InputTotalTime = inputTotalTime
	return a
}
