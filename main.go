// @title           Availability Scheduler API
// @version         1.0
// @description     API for managing scheduling sessions.
// @host            localhost:8080
// @BasePath        /api/v1
package main

import (
	"context"
	"log"

	"github.com/galinavitsk/availability-scheduler-api/config"
	"github.com/galinavitsk/availability-scheduler-api/internal/database"
	"github.com/galinavitsk/availability-scheduler-api/internal/handlers"
	"github.com/galinavitsk/availability-scheduler-api/internal/repository/postgres"
	"github.com/galinavitsk/availability-scheduler-api/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if err := database.RunMigrations(); err != nil {
		log.Fatalf("run migrations: %v", err)
	}

	ctx := context.Background()
	pool, err := database.NewPool(ctx)
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer pool.Close()

	sessionSvc := services.NewSessionService(postgres.NewSessionRepository(pool))
	availabilitySvc := services.NewAvailabilityService(postgres.NewAvailabilityRepository(pool))

	r := gin.Default()
	handlers.RegisterRoutes(r, sessionSvc, availabilitySvc)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
