package models

// LearningTransitionForm 学習時間の推移を格納する
type LearningTransitionForm struct {
	// 年ごとの学習時間推移
	YearLearningTransition LearningTransitionSlice `json:year_transition`
	// 月ごとの学習時間推移
	MonthLearningTransition LearningTransitionSlice `json:month_transition`
	// 週ごとの学習時間推移
	DaysLearningTransition LearningTransitionSlice `json:days_transition`
}

// NewLearningTransitionForm コンストラクタ
func NewLearningTransitionForm(yearLearningTransition LearningTransitionSlice,
	monthLearningTransition LearningTransitionSlice,
	daysLearningTransition LearningTransitionSlice) *LearningTransitionForm {
	l := new(LearningTransitionForm)
	l.YearLearningTransition = yearLearningTransition
	l.MonthLearningTransition = monthLearningTransition
	l.DaysLearningTransition = daysLearningTransition
	return l
}
