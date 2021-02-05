package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}
func (p ProductService) GetProductList(page int, limit int) *utils.Paginate {
	var productList []models.Product
	var total int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&productList).Offset(-1).Count(&total)
	return utils.NewPaginate(total, page, limit, productList)
}
