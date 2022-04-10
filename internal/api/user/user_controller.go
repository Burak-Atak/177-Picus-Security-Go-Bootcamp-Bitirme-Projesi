package user

import (
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/helpers"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/cart"
	"github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/internal/domain/user"
	jwtHelper "github.com/Burak-Atak/177-Picus-Security-Go-Bootcamp-Bitirme-Projesi/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	UserService *user.Service
}

// NewUserController creates a new user controller
func NewUserController(userService *user.Service) *Controller {
	return &Controller{
		UserService: userService,
	}
}

// CreateUser creates a new user
func (c *Controller) CreateUser(context *gin.Context) {
	var body user.User
	err := helpers.DecodeBody(&body, context)
	if err != nil {
		return
	}

	newUser := user.NewModel(body.Email, body.Password)
	if err := c.UserService.CreateUser(newUser); err != nil {
		context.JSON(http.StatusSeeOther, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}

	user.Create(newUser)
	token := jwtHelper.CreateToken(newUser)
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"token":   token,
	})

	// Creates cart for new user
	cart.Create(cart.NewModel(newUser.ID))
}

// Login logs in a user
func (c *Controller) Login(context *gin.Context) {
	var body user.User

	err := helpers.DecodeBody(&body, context)
	if err != nil {
		return
	}

	loggedUser, err := c.UserService.Login(body.Email, body.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		context.Abort()
		return
	}
	token := jwtHelper.CreateToken(loggedUser)
	context.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
