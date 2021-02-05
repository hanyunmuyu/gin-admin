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
