package models

import "time"

type Cart struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Amount     int       `json:"amount"`
	ProductID  uint      `json:"product_id"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime;default:null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime;default:null"`

	OrderID    uint      `json:"order_id"` // Foreign key ke Order
}
