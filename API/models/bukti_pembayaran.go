package models

import "time"

type BuktiPembayaran struct {
	ID         uint      `json:"id,omitempty"`
	Photo      string    `json:"photo,omitempty" gorm:"not null"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	User       User      `gorm:"foreignKey:CustomerID" json:"customer"`
	CreatedAt  time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}
