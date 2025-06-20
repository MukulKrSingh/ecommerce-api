package handlers

import (
	"log"
	"net/http"

	"github.com/ecommerce-api/internal/api/rest"
	"github.com/ecommerce-api/internal/api/service"
	"github.com/ecommerce-api/internal/model"
	"github.com/ecommerce-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func InitUserRoutes(rh *rest.RestHandler) {

	app := rh.App
	handler := UserHandler{
		svc: service.UserService{
			UserRepo: repository.NewUserRepository(rh.DB),
			Auth:     rh.Auth,
		},
	}

	pubRoutes := app.Group("/")

	//Public routes for user registration and login
	pubRoutes.POST("/register", handler.Register)
	pubRoutes.POST("/login", handler.Login)

	//Private routes for user profile management
	pvtRoutes := app.Group("/users")

	// Private endpoint
	pvtRoutes.GET("/verify", handler.GetVerificationCode)
	pvtRoutes.POST("/verify", handler.Verify)

	pvtRoutes.POST("/profile", handler.CreateProfile)
	pvtRoutes.GET("/profile", handler.GetProfile)
	pvtRoutes.PATCH("/profile", handler.UpdateProfile)

	pvtRoutes.POST("/cart", handler.AddToCart)
	pvtRoutes.GET("/cart", handler.GetCart)

	pvtRoutes.GET("/order", handler.GetOrders)
	pvtRoutes.GET("/order/:id", handler.GetOrder)

	pvtRoutes.POST("/become-seller", handler.BecomeSeller)

}

func (uh *UserHandler) Register(ctx *gin.Context) {
	// Handler logic for user registration

	log.Println("User registration handler called")
	var user model.User
	var token string
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, user, err := uh.svc.RegisterUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "token": token, "user": user})
}
func (uh *UserHandler) Login(ctx *gin.Context) {
	// Handler logic for user login
	log.Println("User login handler called")
	var token string
	user := model.User{}
	// Bind the JSON request body to the user model
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := uh.svc.Login(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "user": token})
}
func (uh *UserHandler) GetVerificationCode(ctx *gin.Context) {
	// Handler logic for getting verification code
	log.Println("Get verification code handler called")
}
func (uh *UserHandler) Verify(ctx *gin.Context) {
	// Handler logic for verifying user
	log.Println("Verify user handler called")
}
func (uh *UserHandler) CreateProfile(ctx *gin.Context) {
	// Handler logic for creating user profile
}
func (uh *UserHandler) GetProfile(ctx *gin.Context) {
	// Handler logic for getting user profile
}
func (uh *UserHandler) UpdateProfile(ctx *gin.Context) {
	// Handler logic for updating user profile
}
func (uh *UserHandler) AddToCart(ctx *gin.Context) {
	// Handler logic for adding item to cart
}
func (uh *UserHandler) GetCart(ctx *gin.Context) {
	// Handler logic for getting user cart
}
func (uh *UserHandler) GetOrders(ctx *gin.Context) {
	// Handler logic for getting user orders
}
func (uh *UserHandler) GetOrder(ctx *gin.Context) {
	// Handler logic for getting a specific order
}
func (uh *UserHandler) BecomeSeller(ctx *gin.Context) {
	// Handler logic for becoming a seller
}
