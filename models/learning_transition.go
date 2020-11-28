package models

// LearningTransition 学習時間の推移を格納する
type LearningTransition struct {
	Label string `json:"label"`
	Time  int    `json:"time"`
}

// NewLearningTransition コンストラクタ
func NewLearningTransition(label string, time int) *LearningTransition {
	l := new(LearningTransition)
	l.Label = label
	l.Time = time
	return l
}

// LearningTransitionSlice ソート用スライス
type LearningTransitionSlice []*LearningTransition

func (p LearningTransitionSlice) Len() int {
	return len(p)
}

func (p LearningTransitionSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p LearningTransitionSlice) Less(i, j int) bool {
	return p[i].Label < p[j].Label
}
