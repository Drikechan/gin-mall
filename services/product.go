package services

import (
	"context"
	"mime/multipart"
	"sync"
	"test-gin-mall/pkg/utils/log"
	"test-gin-mall/repository/db/dao"
	"test-gin-mall/types"
)

var ProductSrvIns *ProductSrv

type ProductSrv struct {
}

var ProductSrvOnce sync.Once

func GetProductSrv() *ProductSrv {
	ProductSrvOnce.Do(func() {
		ProductSrvIns = &ProductSrv{}
	})
	return ProductSrvIns
}

func (s *ProductSrv) ProductCreate(context context.Context, file *multipart.FileHeader, req *types.CreateProductResp) (resp interface{}, err error) {

	return
}

func (s *ProductSrv) ListProduct(context context.Context, req *types.ProductListReq) (resp interface{}, err error) {
	var total int64
	var condition = make(map[string]interface{})
	if req.CategoryId != 0 {
		condition["category_id"] = req.CategoryId
	}
	productDao := dao.NewProductDao(context)
	products, _ := productDao.ListProductByCondition(condition, req.BasePageTypes)
	total, err = productDao.CountProductsByCondition(condition)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	respList := make([]*types.ProductResp, 0)
	for _, p := range products {
		pResp := &types.ProductResp{
			ID:            p.ID,
			Name:          p.Name,
			CategoryId:    p.CategoryId,
			Title:         p.Title,
			Info:          p.Info,
			ImgPath:       p.ImgPath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			View:          p.View,
			CreateAt:      p.CreatedAt.Unix(),
			Num:           p.Num,
			OnSale:        p.OnSale,
			BossID:        p.BossID,
			BossName:      p.BossName,
			BossAvatar:    p.BossAvatar,
		}
		respList = append(respList, pResp)
	}
	resp = &types.DataListResp{
		Item:  respList,
		Total: total,
	}
	return
}
