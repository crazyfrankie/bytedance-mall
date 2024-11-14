package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"category" gorm:"many2many:product_category"`
}

func (Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	db *gorm.DB
}

func NewProductQuery(db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		db: db,
	}
}

func (pt *ProductQuery) GetByID(ctx context.Context, productId int) (product Product, err error) {
	err = pt.db.WithContext(ctx).Where("id = ?", productId).First(&product).Error
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (pt *ProductQuery) SearchProducts(ctx context.Context, query string) (products []*Product, err error) {
	err = pt.db.WithContext(ctx).Where("name like ? or description like ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	// err = pt.db.WithContext(ctx).Model(&Product{}).Find(&products, "name like ? or description like ?", "%" + query + "%", "%" + query + "%").Error
	return
}
