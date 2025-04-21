package models

import "time"

type Review struct {
	Id        uint      `json:"id"`
	IdUser    uint      `json:"id_user"`
	IdProduct uint      `json:"id_product"`
	User      User      `gorm:"foreignKey:IdUser" json:"user"`
	Product   Product   `gorm:"foreignKey:IdProduct" json:"product"`
	CreateAt  time.Time `json:"create_at" gorm:"autoCreateTime;default:null"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:null" json:"updated_at"`
}
