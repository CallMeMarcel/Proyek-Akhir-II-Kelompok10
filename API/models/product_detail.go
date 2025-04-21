package models

type ProductDetail struct {
	IdUser    int     `json:"id_user"`
	IdProduct int     `json:"id_product"`
	User      User    `gorm:"foreignKey:IdUser" json:"user"`
	Product   Product `gorm:"foreignKey:IdProduct" json:"product"`
}
