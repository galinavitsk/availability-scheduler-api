package repository

import (
	"context"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
)

type SessionRepository interface {
	Create(ctx context.Context, req models.CreateSessionRequest) (*models.Session, error)
	GetBySlug(ctx context.Context, slug string) (*models.Session, error)
	Update(ctx context.Context, id string, req models.UpdateSessionRequest) (*models.Session, error)
	Delete(ctx context.Context, id string) error
}
