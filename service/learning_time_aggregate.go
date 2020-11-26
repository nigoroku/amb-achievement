package service

import (

	// "github.com/kzpolicy/user/generated"
	"github.com/volatiletech/sqlboiler/boil"

	"context"
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

// CalcTotalLearningTime 総学習時間を算出する
func (in *LearningTimeAggregateService) CalcTotalLearningTime(userID int) (int, error) {

	// var categories generated.MCategorySlice
	// err := generated.NewQuery(
	// 	qm.Select("c.*"),
	// 	qm.From("input_achievement_tags as iat"),
	// 	qm.InnerJoin("m_categories c ON iat.category_id = c.category_id"),
	// 	qm.Where("iat.input_achievement_id=?", inputAchievementId),
	// ).Bind(in.ctx, in.db, &categories)

	// return categories, err
	return 1, nil
}
