package api

import (
	cartApi "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/api/cart"
	orderApi "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/api/order"
	productApi "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/api/product"
	userApi "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/api/user"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/order"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/product"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/user"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterHandlers registers all the handlers
func RegisterHandlers(r *gin.Engine) {

	userService := user.NewService()
	userController := userApi.NewUserController(userService)

	cartService := cart.NewService()
	cartController := cartApi.NewCartController(cartService)

	orderService := order.NewService()
	orderController := orderApi.NewOrderController(orderService)

	productService := product.NewService()
	productController := productApi.NewProductController(productService)

	userGroup := r.Group("/user")
	userGroup.POST("/register", userController.CreateUser)
	userGroup.POST("/login", userController.Login)

	cartGroup := r.Group("/cart")
	cartGroup.GET("/list", middleware.AuthForGeneral(), cartController.GetCartList)
	cartGroup.DELETE("/delete", middleware.AuthForGeneral(), cartController.DeleteProductFromCart)
	cartGroup.PUT("/update", middleware.AuthForGeneral(), cartController.UpdateProductInCart)
	cartGroup.POST("/add", middleware.AuthForGeneral(), cartController.AddProductToCart)

	orderGroup := r.Group("/order")
	orderGroup.GET("/list", middleware.AuthForGeneral(), orderController.GetOrderList)
	orderGroup.POST("/create", middleware.AuthForGeneral(), orderController.CreateOrder)
	orderGroup.GET("/detail", middleware.AuthForGeneral(), orderController.GetOrderDetail)
	orderGroup.DELETE("/cancel", middleware.AuthForGeneral(), orderController.CancelOrder)

	productGroup := r.Group("/product")
	productGroup.GET("/list", productController.GetProductList)
	productGroup.GET("/detail", productController.GetProductDetail)
	productGroup.GET("/search", productController.SearchProduct)
	productGroup.POST("/create", middleware.AuthForAdmin(), productController.CreateProduct)
	productGroup.PUT("/update", middleware.AuthForAdmin(), productController.UpdateProduct)
	productGroup.DELETE("/delete", middleware.AuthForAdmin(), productController.DeleteProduct)

}
