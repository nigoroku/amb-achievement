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

type OutputService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewOutputService() *OutputService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &OutputService{ctx, db}
}

func (o *OutputService) FindByUser(userId int) (*models.OutputAchievement, error) {
	// データ取得
	return models.OutputAchievements(qm.Where("user_id=? and created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)", userId)).One(o.ctx, o.db)
}

func (ou *OutputService) FindCategoriesBy(outputAchievementId int) (models.MCategorySlice, error) {
	// データ取得
	var categories models.MCategorySlice
	err := models.NewQuery(
		qm.Select("c.*"),
		qm.From("output_achievement_tags as oat"),
		qm.InnerJoin("m_categories c ON oat.category_id = c.category_id"),
		qm.Where("oat.output_achievement_id=?", outputAchievementId),
	).Bind(ou.ctx, ou.db, &categories)

	return categories, err
}

func (o *OutputService) AddOutput(output *models.OutputAchievement, categoryIds []string) error {
	now := time.Now()

	var oa models.OutputAchievement
	oa.ReferenceURL = output.ReferenceURL
	oa.Summary = output.Summary
	oa.OutputTime = output.OutputTime
	oa.UserID = output.UserID
	oa.CreatedBy = output.UserID
	oa.CreatedAt = now
	err := oa.Insert(o.ctx, o.db, boil.Infer())

	for _, categoryID := range categoryIds {

		var category models.OutputAchievementTag
		c, _ := strconv.Atoi(categoryID)
		category.CategoryID = c
		category.OutputAchievementID = oa.OutputAchievementID
		category.CreatedBy = output.UserID
		category.CreatedAt = now
		err2 := category.Insert(o.ctx, o.db, boil.Infer())

		if err2 != nil {
			fmt.Println(err2)
			return err2
		}
	}

	return err
}

func (o *OutputService) Update(output models.OutputAchievement, categoryIds []string, id int) error {
	now := time.Now()

	updCols := map[string]interface{}{
		models.OutputAchievementColumns.ReferenceURL: output.ReferenceURL,
		models.OutputAchievementColumns.Summary:      output.Summary,
		models.OutputAchievementColumns.OutputTime:   output.OutputTime,
		models.OutputAchievementColumns.UserID:       output.UserID,
		models.OutputAchievementColumns.ModifiedBy:   output.UserID,
		models.OutputAchievementColumns.ModifiedAt:   now,
	}

	query := qm.WhereIn(models.OutputAchievementColumns.OutputAchievementID+" = ?", id)

	_, err := models.OutputAchievements(query).UpdateAll(o.ctx, o.db, updCols)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// categorys は、delete & insert
	_, err2 := models.OutputAchievementTags(qm.Where("output_achievement_id=?", id)).DeleteAll(o.ctx, o.db)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	for _, categoryID := range categoryIds {

		var category models.OutputAchievementTag
		c, _ := strconv.Atoi(categoryID)
		category.CategoryID = c
		category.OutputAchievementID = id
		category.CreatedBy = output.UserID
		category.CreatedAt = now
		err3 := category.Insert(o.ctx, o.db, boil.Infer())

		if err3 != nil {
			fmt.Println(err3)
			return err3
		}
	}

	return nil
}
