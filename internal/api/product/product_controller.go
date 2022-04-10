package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/category"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	products, err := c.ProductService.GetProductList()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	outOut := make([]product.Product, len(products))
	context.JSON(http.StatusOK, gin.H{
		"products": outOut,
	})
}

func (c Controller) GetProductDetail(context *gin.Context) {

}

func (c Controller) SearchProduct(context *gin.Context) {
	search, isOk := context.GetQuery("search")
	if !isOk {
		context.JSON(http.StatusOK, gin.H{
			"message": "search param is required",
		})
		context.Abort()
		return
	}
	products, err := c.ProductService.SearchProduct(search)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	outOut := make([]product.Product, len(products))
	context.JSON(http.StatusOK, gin.H{
		"products": outOut,
	})
}

func (c Controller) UpdateProduct(context *gin.Context) {

}

func (c Controller) CreateProduct(context *gin.Context) {
	var body product.Product

	err := helpers.DecodeBody(&body, context)
	if err != nil {
		context.Abort()
		return
	}

	err = c.ProductService.CreateProduct(body.ProductName, body.SKU)
	if err != nil {
		context.JSON(http.StatusAlreadyReported, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	categoryService := category.NewService()
	_, err = categoryService.CreateCategory(body.CategoryName)
	if err == nil {
		category.NewModel(body.CategoryName)
	}

	newProduct := product.NewModel(body.ProductName, body.CategoryName, body.Price, body.Stock, body.SKU)
	product.Create(newProduct)

	context.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
	})

}

func (c Controller) DeleteProduct(context *gin.Context) {
	productId, isOk := context.GetQuery("productId")
	cartService := cart.NewService()
	if !isOk {
		context.JSON(http.StatusOK, gin.H{
			"message": "productId param is required",
		})
		context.Abort()
		return
	}

	id, err := strconv.Atoi(productId)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": helpers.InvalidIdError,
		})
		context.Abort()
		return
	}

	chosenProduct, err := c.ProductService.GetProduct(uint(id))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	cartsHasChosenProduct, err := cartService.GetCartsHasProduct(chosenProduct.ID)

	if err == nil {
		for _, cartDetail := range *cartsHasChosenProduct {
			newTotalPrice := -cartDetail.TotalPrice
			newAmount := -cartDetail.Amount
			cart.UpdateUserCart(cartDetail.CartId, newAmount, newTotalPrice)
			cart.DeleteProductInCart(cartDetail.CartId, chosenProduct.ID)
		}
	}

	product.DeleteProduct(*chosenProduct)

	context.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
