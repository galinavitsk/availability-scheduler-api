package repository

import (
	"context"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
)

type AvailabilityRepository interface {
	CreateAvailability(ctx context.Context, req models.CreateAvailabilityRequest) (*models.Availability, error)
	GetAllAvailabilitiesForSlug(ctx context.Context, slug string) ([]models.Availability, error)
	UpdateAvailability(ctx context.Context, req models.UpdateAvailabilityRequest) (*models.Availability, error)
}
