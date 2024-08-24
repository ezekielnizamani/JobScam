package v1

import (
	authHandlers "github.com/ezekielnizamani/JobScam/internal/api/v1/handlers/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/signup", authHandlers.SignUpHandler)
		authRoutes.POST("/signin", authHandlers.SignInHandler)
	}
}
