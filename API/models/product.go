package models

import "time"

type Product struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Price       float64   `json:"price" gorm:"not null"`
	Image       string    `json:"image" gorm:"not null"`
	CreateAt    time.Time `json:"create_at" gorm:"autoCreateTime;default:null"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;default:null" json:"updated_at"`
	IdCategory  uint      `json:"id_category"`
}
