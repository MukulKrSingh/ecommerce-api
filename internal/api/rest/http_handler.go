package rest

import (
	"github.com/ecommerce-api/internal/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  *gin.Engine
	DB   *gorm.DB
	Auth helpers.Auth
}
