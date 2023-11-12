package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryId    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string
	View          uint64
	CreateAt      int64
	Num           int
	OnSale        bool `gorm:"default:false"`
	BossID        uint
	BossName      string
	BossAvatar    string
}
