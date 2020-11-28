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

func (o *OutputService) FindByUser(userId int) (*generated.OutputAchievement, error) {
	// データ取得
	return generated.OutputAchievements(qm.Where("user_id=? and created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)", userId)).One(o.ctx, o.db)
}

func (ou *OutputService) FindCategoriesBy(outputAchievementId int) (generated.MCategorySlice, error) {
	// データ取得
	var categories generated.MCategorySlice
	err := generated.NewQuery(
		qm.Select("c.*"),
		qm.From("output_achievement_tags as oat"),
		qm.InnerJoin("m_categories c ON oat.category_id = c.category_id"),
		qm.Where("oat.output_achievement_id=?", outputAchievementId),
	).Bind(ou.ctx, ou.db, &categories)

	return categories, err
}

func (o *OutputService) AddOutput(output *generated.OutputAchievement, categoryIds []string) error {
	now := time.Now()

	var oa generated.OutputAchievement
	oa.ReferenceURL = output.ReferenceURL
	oa.Summary = output.Summary
	oa.OutputTime = output.OutputTime
	oa.UserID = output.UserID
	oa.CreatedBy = output.UserID
	oa.CreatedAt = now
	err := oa.Insert(o.ctx, o.db, boil.Infer())

	for _, categoryID := range categoryIds {

		var category generated.OutputAchievementTag
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

func (o *OutputService) Update(output generated.OutputAchievement, categoryIds []string, id int) error {
	now := time.Now()

	updCols := map[string]interface{}{
		generated.OutputAchievementColumns.ReferenceURL: output.ReferenceURL,
		generated.OutputAchievementColumns.Summary:      output.Summary,
		generated.OutputAchievementColumns.OutputTime:   output.OutputTime,
		generated.OutputAchievementColumns.UserID:       output.UserID,
		generated.OutputAchievementColumns.ModifiedBy:   output.UserID,
		generated.OutputAchievementColumns.ModifiedAt:   now,
	}

	query := qm.WhereIn(generated.OutputAchievementColumns.OutputAchievementID+" = ?", id)

	_, err := generated.OutputAchievements(query).UpdateAll(o.ctx, o.db, updCols)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(categoryIds) == 0 {
		return nil
	}

	// categorys は、delete & insert
	_, err2 := generated.OutputAchievementTags(qm.Where("output_achievement_id=?", id)).DeleteAll(o.ctx, o.db)

	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	for _, categoryID := range categoryIds {

		var category generated.OutputAchievementTag
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
