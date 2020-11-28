package models

// CategoryDistribution カテゴリごとの学習時間を格納する
type CategoryDistribution struct {
	CategoryName string `boil:"category_name" json:"category_name"`
	TotalTime    int    `boil:"total_time" json:"total_time"`
}

// NewCategoryDistribution コンストラクタ
func NewCategoryDistribution(categoryName string, totalTime int) *CategoryDistribution {
	c := new(CategoryDistribution)
	c.CategoryName = categoryName
	c.TotalTime = totalTime
	return c
}
