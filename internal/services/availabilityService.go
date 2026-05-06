package services

import (
	"context"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
	"github.com/galinavitsk/availability-scheduler-api/internal/repository"
)

type AvailabilityService struct {
	repo repository.AvailabilityRepository
}

func NewAvailabilityService(repo repository.AvailabilityRepository) *AvailabilityService {
	return &AvailabilityService{repo: repo}
}

func (service *AvailabilityService) CreateAvailability(ctx context.Context, req models.CreateAvailabilityRequest) (*models.Availability, error) {
	return service.repo.CreateAvailability(ctx, req)
}

func (service *AvailabilityService) GetAllAvailabilitiesForSlug(ctx context.Context, slug string) ([]models.Availability, error) {
	return service.repo.GetAllAvailabilitiesForSlug(ctx, slug)
}

func (service *AvailabilityService) UpdateAvailability(ctx context.Context, req models.UpdateAvailabilityRequest) (*models.Availability, error) {
	return service.repo.UpdateAvailability(ctx, req)
}
