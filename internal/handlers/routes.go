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

func RegisterRoutes(r *gin.Engine, sessionSvc *services.SessionService, availabilitySvc *services.AvailabilityService) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3005", "http://localhost:3000", "https://quest-scheduler.vercel.app"},
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
		ah := &availabilityHandler{service: availabilitySvc}
		sessions := v1.Group("/sessions")
		{
			sessions.POST("", sh.Create)
			sessions.GET("/:slug", sh.GetBySlug)
			sessions.PUT("/:id", sh.Update)
			sessions.DELETE("/:id", sh.Delete)
		}
		availability := v1.Group("/availability")
		{
			availability.POST("", ah.CreateAvailability)
			availability.GET("/:slug", ah.GetAllAvailabilitiesForSlug)
			availability.POST("/update", ah.UpdateAvailability)
		}
	}
}
