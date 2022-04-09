package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	ProductService *product.Service
}

func NewProductController(productService *product.Service) *Controller {
	return &Controller{
		ProductService: productService,
	}
}

func (c Controller) GetProductList(context *gin.Context) {

}

func (c Controller) GetProductDetail(context *gin.Context) {

}

func (c Controller) SearchProduct(context *gin.Context) {

}

func (c Controller) UpdateProduct(context *gin.Context) {

}

func (c Controller) CreateProduct(context *gin.Context) {

}

func (c Controller) DeleteProduct(context *gin.Context) {

}
