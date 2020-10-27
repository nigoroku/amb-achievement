package service

import (
	"fmt"
	"strconv"
	"time"

	// "github.com/kzpolicy/user/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/models"

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

func (in *InputService) FindByUser(userId int) (*models.InputAchievement, error) {
	// データ取得
	return models.InputAchievements(qm.Where("user_id=? and created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)", userId)).One(in.ctx, in.db)
}

func (in *InputService) FindCategoriesBy(inputAchievementId int) (models.MCategorySlice, error) {
	// データ取得
	var categories models.MCategorySlice
	err := models.NewQuery(
		qm.Select("c.*"),
		qm.From("input_achievement_tags as iat"),
		qm.InnerJoin("m_categories c ON iat.category_id = c.category_id"),
		qm.Where("iat.input_achievement_id=?", inputAchievementId),
	).Bind(in.ctx, in.db, &categories)

	return categories, err
}

func (in *InputService) AddInput(input *models.InputAchievement, categoryIds []string) error {
	now := time.Now()

	var ia models.InputAchievement
	ia.ReferenceURL = input.ReferenceURL
	ia.Summary = input.Summary
	ia.InputTime = input.InputTime
	ia.UserID = input.UserID
	ia.CreatedBy = input.UserID
	ia.CreatedAt = now
	err := ia.Insert(in.ctx, in.db, boil.Infer())

	for _, categoryID := range categoryIds {

		var category models.InputAchievementTag
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

func (in *InputService) Update(input models.InputAchievement, categoryIds []string, id int) error {
	now := time.Now()

	updCols := map[string]interface{}{
		models.InputAchievementColumns.ReferenceURL: input.ReferenceURL,
		models.InputAchievementColumns.Summary:      input.Summary,
		models.InputAchievementColumns.InputTime:    input.InputTime,
		models.InputAchievementColumns.UserID:       input.UserID,
		models.InputAchievementColumns.ModifiedBy:   input.UserID,
		models.InputAchievementColumns.ModifiedAt:   now,
	}

	query := qm.WhereIn(models.InputAchievementColumns.InputAchievementID+" = ?", id)

	_, err := models.InputAchievements(query).UpdateAll(in.ctx, in.db, updCols)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// categorys は、delete & insert
	_, err2 := models.InputAchievementTags(qm.Where("input_achievement_id=?", id)).DeleteAll(in.ctx, in.db)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	for _, categoryID := range categoryIds {

		var category models.InputAchievementTag
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
