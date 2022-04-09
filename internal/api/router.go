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
	userGroup.POST("/logout", userController.Logout)

	cartGroup := r.Group("/cart")
	cartGroup.GET("/cart/list", cartController.GetCartList)
	cartGroup.DELETE("/cart/delete", cartController.DeleteItemFromCart)
	cartGroup.PUT("/cart/update", cartController.UpdateItemInCart)
	cartGroup.POST("/cart/add", cartController.AddItemToCart)

	orderGroup := r.Group("/order")
	orderGroup.GET("/order/list", orderController.GetOrderList)
	orderGroup.POST("/order/create", orderController.CreateOrder)
	orderGroup.GET("/order/detail", orderController.GetOrderDetail)
	orderGroup.DELETE("/order/cancel", orderController.CancelOrder)

	productGroup := r.Group("/product")
	productGroup.GET("/list", productController.GetProductList)
	productGroup.GET("/detail", productController.GetProductDetail)
	productGroup.GET("/search", productController.SearchProduct)
	productGroup.POST("/create", productController.CreateProduct)
	productGroup.PUT("/update", productController.UpdateProduct)
	productGroup.DELETE("/delete", productController.DeleteProduct)

}
