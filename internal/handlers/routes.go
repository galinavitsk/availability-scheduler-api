package handlers

import (
	"time"

	_ "github.com/galinavitsk/availability-scheduler-api/docs"
	"github.com/galinavitsk/availability-scheduler-api/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine, sessionSvc *services.SessionService) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3005"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/health", HealthCheck)

	v1 := r.Group("/api/v1")
	{
		sh := &sessionHandler{service: sessionSvc}
		sessions := v1.Group("/sessions")
		{
			sessions.POST("", sh.Create)
			sessions.GET("/:id", sh.GetByID)
			sessions.PUT("/:id", sh.Update)
			sessions.DELETE("/:id", sh.Delete)
		}
	}
}
