package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint32 `gorm:"type:int(11);not null;index:index_user_id"`
	ProductID uint32 `gorm:"type:int(11);not null"`
	Quantity  int32  `gorm:"type:int(11);not null"`
}

func (Cart) TableName() string {
	return "cart"
}

type CartQuery struct {
	db *gorm.DB
}

func NewCartQuery(db *gorm.DB) *CartQuery {
	return &CartQuery{
		db: db,
	}
}

func (cq *CartQuery) AddItem(ctx context.Context, cart *Cart) error {
	var row Cart
	err := cq.db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserID: cart.UserID, ProductID: cart.ProductID}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if row.ID > 0 {
		return cq.db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserID: cart.UserID, ProductID: cart.ProductID}).
			UpdateColumn("quantity", gorm.Expr("quantity+?", cart.Quantity)).Error
	}

	return cq.db.WithContext(ctx).Create(cart).Error
}

func (cq *CartQuery) EmptyCart(ctx context.Context, uid uint32) error {
	if uid == 0 {
		return errors.New("user id is required")
	}

	return cq.db.WithContext(ctx).Delete(&Cart{}, "user_id = ?", uid).Error
}

func (cq *CartQuery) GetCart(ctx context.Context, uid uint32) (carts []*Cart, err error) {
	err = cq.db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: uid}).Find(&carts).Error
	return carts, err
}
