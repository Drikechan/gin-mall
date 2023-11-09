package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"imgPath"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discountPrice"`
	View          uint64 `json:"view"`
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossID        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}
