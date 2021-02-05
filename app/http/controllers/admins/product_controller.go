package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/services/admins"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	productService = admins.NewProductService()
)

type ProductController struct {
	http.BaseController
}

func (p *ProductController) GetProductList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	productList := productService.GetProductList(page, 15)
	p.Success(ctx, productList)
}
func (p *ProductController) DeleteProduct(ctx *gin.Context) {
	productId := 0
	if p, err := strconv.Atoi(ctx.Param("productId")); err == nil {
		productId = p
	}
	if productId <= 0 {
		p.Error(ctx, "productId >0")
		return
	}
	product := productService.GetProductById(productId)
	if product.ID == 0 {
		p.Error(ctx, "产品不能存在")
		return
	}
	row := productService.DeleteProductByProductId(productId)
	if row == 0 {
		p.Error(ctx, "删除失败")
		return
	}
	p.Success(ctx, gin.H{})
}
