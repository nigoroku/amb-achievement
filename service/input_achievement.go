package service

import (
	"fmt"
	"strconv"
	"time"

	// "github.com/kzpolicy/user/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/generated"

	"context"
)

type InputService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewInputService() *InputService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &InputService{ctx, db}
}

func (in *InputService) FindByUser(userId int) (*generated.InputAchievement, error) {
	// データ取得
	return generated.InputAchievements(qm.Where("user_id=? and created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)", userId)).One(in.ctx, in.db)
}

func (in *InputService) FindCategoriesBy(inputAchievementId int) (generated.MCategorySlice, error) {
	// データ取得
	var categories generated.MCategorySlice
	err := generated.NewQuery(
		qm.Select("c.*"),
		qm.From("input_achievement_tags as iat"),
		qm.InnerJoin("m_categories c ON iat.category_id = c.category_id"),
		qm.Where("iat.input_achievement_id=?", inputAchievementId),
	).Bind(in.ctx, in.db, &categories)

	return categories, err
}

func (in *InputService) AddInput(input *generated.InputAchievement, categoryIds []string) error {
	now := time.Now()

	var ia generated.InputAchievement
	ia.ReferenceURL = input.ReferenceURL
	ia.Summary = input.Summary
	ia.InputTime = input.InputTime
	ia.UserID = input.UserID
	ia.CreatedBy = input.UserID
	ia.CreatedAt = now
	err := ia.Insert(in.ctx, in.db, boil.Infer())

	for _, categoryID := range categoryIds {

		var category generated.InputAchievementTag
		c, _ := strconv.Atoi(categoryID)
		category.CategoryID = c
		category.InputAchievementID = ia.InputAchievementID
		category.CreatedBy = input.UserID
		category.CreatedAt = now
		err2 := category.Insert(in.ctx, in.db, boil.Infer())

		if err2 != nil {
			fmt.Println(err2)
			return err2
		}
	}
	return err
}

func (in *InputService) Update(input generated.InputAchievement, categoryIds []string, id int) error {
	now := time.Now()

	updCols := map[string]interface{}{
		generated.InputAchievementColumns.ReferenceURL: input.ReferenceURL,
		generated.InputAchievementColumns.Summary:      input.Summary,
		generated.InputAchievementColumns.InputTime:    input.InputTime,
		generated.InputAchievementColumns.UserID:       input.UserID,
		generated.InputAchievementColumns.ModifiedBy:   input.UserID,
		generated.InputAchievementColumns.ModifiedAt:   now,
	}

	query := qm.WhereIn(generated.InputAchievementColumns.InputAchievementID+" = ?", id)

	_, err := generated.InputAchievements(query).UpdateAll(in.ctx, in.db, updCols)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(categoryIds) == 0 {
		return nil
	}

	// categorys は、delete & insert
	_, err2 := generated.InputAchievementTags(qm.Where("input_achievement_id=?", id)).DeleteAll(in.ctx, in.db)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	for _, categoryID := range categoryIds {

		var category generated.InputAchievementTag
		c, _ := strconv.Atoi(categoryID)
		category.CategoryID = c
		category.InputAchievementID = id
		category.CreatedBy = input.UserID
		category.CreatedAt = now
		err3 := category.Insert(in.ctx, in.db, boil.Infer())

		if err3 != nil {
			fmt.Println(err3)
			return err3
		}
	}

	return nil
}
