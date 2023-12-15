package services

import (
	"context"
	"mime/multipart"
	"strconv"
	"sync"
	conf "test-gin-mall/config"
	"test-gin-mall/consts"
	"test-gin-mall/pkg/utils/ctl"
	"test-gin-mall/pkg/utils/log"
	"test-gin-mall/pkg/utils/upload"
	"test-gin-mall/repository/db/dao"
	"test-gin-mall/repository/db/model"
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

func (s *ProductSrv) ProductCreate(context context.Context, files []*multipart.FileHeader, req *types.CreateProductResp) (resp interface{}, err error) {
	u, err := ctl.GetUserInfo(context)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}
	uId := u.Id
	boss, _ := dao.NewUserDao(context).GetUserById(uId)
	tmp, _ := files[0].Open()
	var path string
	if conf.Config.System.UploadModel == consts.UploadModalLocal {
		path, err = upload.ProductUploadToLocalStatic(tmp, uId, req.Name)
	} else {
		return
	}

	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
	product := &model.Product{
		Name:          req.Name,
		CategoryId:    req.CategoryID,
		Title:         req.Title,
		Info:          req.Info,
		ImgPath:       path,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Num:           req.Num,
		OnSale:        true,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(context)
	err = productDao.CreateProduct(product)
	if err != nil {
		log.LogrusObj.Error(err)
		return nil, err
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))

	for index, file := range files {
		num := strconv.Itoa(index)
		tmp, _ := file.Open()
		if conf.Config.System.UploadModel == consts.UploadModalLocal {
			path, err = upload.ProductUploadToLocalStatic(tmp, uId, req.Name+num)
		}

		if err != nil {
			log.LogrusObj.Error(err)
			return
		}

		productImg := &model.ProductImg{
			ProductId: product.ID,
			ImgPath:   path,
		}

		err = dao.NewProductImgDaoByDB(productDao.DB).CreateProductImg(productImg)
		if err != nil {
			log.LogrusObj.Error(err)
			return nil, err
		}

		wg.Done()

	}

	wg.Wait()

	return
}

func (s *ProductSrv) ProductUpdate(context context.Context, req *types.UpdateProductResp) (resp interface{}, err error) {
	product := &model.Product{
		Name:          req.Name,
		CategoryId:    req.CategoryId,
		Title:         req.Title,
		Info:          req.Info,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		OnSale:        req.OnSale,
	}
	err = dao.NewProductDao(context).UpdateProduct(req.ID, product)
	if err != nil {
		log.LogrusObj.Error(err)
		return
	}
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
