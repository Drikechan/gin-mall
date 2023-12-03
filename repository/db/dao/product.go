package dao

import (
	"context"
	"gorm.io/gorm"
	"test-gin-mall/repository/db/model"
	"test-gin-mall/types"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page types.BasePageTypes) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((page.CurrentPage - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Error
	return
}

func (dao *ProductDao) CountProductsByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) CreateProduct(product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}
