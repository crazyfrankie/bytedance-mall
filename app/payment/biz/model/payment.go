package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId        uint32    `json:"user_id"`
	OrderId       string    `json:"order_id"`
	TransactionId string    `json:"transaction_id"`
	Amount        float32   `json:"amount"`
	PayAt         time.Time `json:"payat"`
}

func (PaymentLog) TableName() string {
	return "payment_log"
}

type PaymentQuery struct {
	db *gorm.DB
}

func NewPaymentQuery(db *gorm.DB) *PaymentQuery {
	return &PaymentQuery{
		db: db,
	}
}

func (pq *PaymentQuery) CreatePaymentLog(ctx context.Context, payment *PaymentLog) error {
	return pq.db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}
