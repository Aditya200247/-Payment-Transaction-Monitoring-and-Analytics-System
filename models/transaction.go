package models

import "time"

type Transaction struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	TransactionID string    `json:"transaction_id" gorm:"uniqueIndex;not null" binding:"required"`
	MerchantID    string    `json:"merchant_id" gorm:"index;not null" binding:"required"`
	Amount        float64   `json:"amount" gorm:"not null" binding:"required"`
	Status        string    `json:"status" gorm:"type:varchar(10);not null" binding:"required"`
	PaymentMethod string    `json:"payment_method" gorm:"type:varchar(10)"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
