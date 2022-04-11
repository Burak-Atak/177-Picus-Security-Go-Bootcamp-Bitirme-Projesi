package product

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/category"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/pagination"
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

// SearchProduct returns a list of products by given search query
func (c Controller) SearchProduct(context *gin.Context) {
	search, isOk := context.GetQuery("search")
	if !isOk {
		context.JSON(http.StatusOK, gin.H{
			"message": helpers.SearchParamRequireError.Error(),
		})
		context.Abort()
		return
	}

	allProducts, err := c.ProductService.SearchProduct(search)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products := c.ProductService.SearchProductWithPagination(search, pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(context, len(allProducts))

	if len(products) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": helpers.PageCouldNotBeFoundError.Error(),
		})
		context.Abort()
		return
	}

	outOut := make([]Product, len(products))
	for i, p := range products {
		outOut[i] = Product{
			ProductName:  p.ProductName,
			Price:        p.Price,
			Stock:        p.Stock,
			ProductId:    p.ID,
			CategoryName: p.CategoryName,
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"Info":     paginatedResult,
		"products": outOut,
	})
}

// UpdateProduct updates a product by given params if stock and price changed it will update the cart
func (c Controller) UpdateProduct(context *gin.Context) {
	cartService := cart.NewService()
	id, isProductId := context.GetQuery("id")
	newStock, isStock := context.GetQuery("stock")
	newPrice, isPrice := context.GetQuery("price")
	newName, isName := context.GetQuery("name")
	newSku, isSku := context.GetQuery("sku")

	if !isProductId {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": helpers.IdIsRequiredError.Error(),
		})
		context.Abort()
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": helpers.InvalidIdError.Error(),
		})
		context.Abort()
		return
	}

	chosenProduct, err := c.ProductService.GetProduct(uint(productId))

	// Updates stock
	if isStock {
		stock, err := strconv.Atoi(newStock)
		if err != nil || stock <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": helpers.InvalidStockError.Error(),
			})
			context.Abort()
			return
		}
		product.UpdateStock(*chosenProduct, stock)

	}

	// Updates sku
	if isSku {
		if chosenProduct.SKU == newSku {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": helpers.SkuIsSameError.Error(),
			})
			context.Abort()
			return
		}

		sameSkuProduct := product.SearchBySKU(newSku)
		if sameSkuProduct.ID != 0 {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": helpers.SkuAlreadyExistError.Error(),
			})
			context.Abort()
			return
		}
		product.UpdateSKU(*chosenProduct, newSku)
	}

	cartsHasChosenProduct, cartIsEmpty := cartService.GetCartsHasProduct(chosenProduct.ID)

	// Updates price
	if isPrice {
		price, err := strconv.ParseFloat(newPrice, 64)
		if err != nil || price <= 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": helpers.InvalidPriceError.Error(),
			})
			context.Abort()
			return
		}

		if cartIsEmpty == nil {
			for _, cartDetail := range *cartsHasChosenProduct {
				newTotalPrice := (chosenProduct.Price - price) * float64(cartDetail.Amount)
				newAmount := cartDetail.Amount

				cartDetail.UnitPrice = price
				cartDetail.TotalPrice = newTotalPrice
				cart.UpdateUserCart(cartDetail.CartId, newAmount, newTotalPrice)
				cart.UpdateModel(&cartDetail)
			}
		}

		product.UpdatePrice(*chosenProduct, price)
	}

	// Updates name
	if isName {
		if chosenProduct.ProductName == newName {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": helpers.NameIsSameError.Error(),
			})
			context.Abort()
			return
		}

		sameNameProduct := product.SearchByProductName(newName)
		if sameNameProduct.ID != 0 {
			context.JSON(http.StatusAlreadyReported, gin.H{
				"message": helpers.NameAlreadyExistError.Error(),
			})
			context.Abort()
			return
		}

		if cartIsEmpty == nil {
			for _, cartDetail := range *cartsHasChosenProduct {
				cartDetail.ProductName = newName
				cart.UpdateModel(&cartDetail)
			}
		}
		product.UpdateName(*chosenProduct, newName)
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "product updated",
	})

}

// CreateProduct creates a new product
func (c Controller) CreateProduct(context *gin.Context) {
	var body product.Product

	err := helpers.DecodeBody(&body, context)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
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
	err = categoryService.GetCategoryByName(body.CategoryName)
	if err == nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": helpers.CategoryNotFoundError.Error(),
		})
		context.Abort()
		return
	}

	newProduct := product.NewModel(body.ProductName, body.CategoryName, body.Price, body.Stock, body.SKU)
	product.Create(newProduct)

	context.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
	})
}

// DeleteProduct deletes a product and delete it from all carts
func (c Controller) DeleteProduct(context *gin.Context) {
	productId, isOk := context.GetQuery("id")
	cartService := cart.NewService()
	if !isOk {
		context.JSON(http.StatusOK, gin.H{
			"message": helpers.InvalidIdError.Error(),
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

// GetProductList returns a list of products
func (c Controller) GetProductList(context *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(context)
	products, allProducts := c.ProductService.GetProductList(pageIndex, pageSize)
	paginatedResult := pagination.NewFromGinRequest(context, allProducts)

	if len(products) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"message": helpers.PageCouldNotBeFoundError.Error(),
		})
		context.Abort()
		return
	}

	outOut := make([]Product, len(products))
	for i, p := range products {
		outOut[i] = Product{
			ProductName:  p.ProductName,
			Price:        p.Price,
			Stock:        p.Stock,
			ProductId:    p.ID,
			CategoryName: p.CategoryName,
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"Info":     paginatedResult,
		"Products": outOut,
	})
}
