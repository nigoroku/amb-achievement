package service

import (

	// "github.com/kzpolicy/user/models"
	"github.com/volatiletech/sqlboiler/boil"
	"local.packages/generated"

	"context"
)

type CategoryService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewCategoryService() *CategoryService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &CategoryService{ctx, db}
}

func (c *CategoryService) FindAll() (generated.MCategorySlice, error) {
	// データ取得
	return generated.MCategories().All(c.ctx, c.db)
}
