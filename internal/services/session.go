package services

import (
	"context"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
	"github.com/galinavitsk/availability-scheduler-api/internal/repository"
)

type SessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (service *SessionService) Create(ctx context.Context, req models.CreateSessionRequest) (*models.Session, error) {
	return service.repo.Create(ctx, req)
}

func (service *SessionService) GetByID(ctx context.Context, id string) (*models.Session, error) {
	return service.repo.GetByID(ctx, id)
}

func (service *SessionService) Update(ctx context.Context, id string, req models.UpdateSessionRequest) (*models.Session, error) {
	return service.repo.Update(ctx, id, req)
}

func (service *SessionService) Delete(ctx context.Context, id string) error {
	return service.repo.Delete(ctx, id)
}
