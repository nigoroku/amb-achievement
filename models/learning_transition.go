package models

import "github.com/volatiletech/null"

// LearningTransition 学習時間の推移を格納する
type LearningTransition struct {
	Label null.String `json:"label"`
	Time  null.Int    `json:"time"`
}

// NewLearningTransition コンストラクタ
func NewLearningTransition(label null.String, time null.Int) *LearningTransition {
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
	return p[i].Label.String < p[j].Label.String
}
