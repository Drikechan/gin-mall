package dao

import (
	"gorm.io/gorm"
	"test-gin-mall/repository/db/model"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao {
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	err = dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error
	return
}
