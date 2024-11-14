package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name       string `json:"name"`
	Descrption string `json:"descrption"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (Category) TableName() string {
	return "category"
}

func NewCategoryQuery(db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		db: db,
	}
}

type CategoryQuery struct {
	db *gorm.DB
}

func (cq *CategoryQuery) GetProductsByCategoryName(ctx context.Context, name string) (categories []*Category, err error) {
	// err = cq.db.WithContext(ctx).Where("name = ?", name).Preload("Products").Find(&categories).Error
	err = cq.db.WithContext(ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	return
}
