package postgres

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type sessionRepo struct {
	db *pgxpool.Pool
}

func NewSessionRepository(db *pgxpool.Pool) *sessionRepo {
	return &sessionRepo{db: db}
}

func generateSlug() (string, error) {
	b := make([]byte, 4) // 4 bytes = 8 hex chars
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("generate slug: %w", err)
	}
	return hex.EncodeToString(b), nil
}

func (r *sessionRepo) Create(ctx context.Context, req models.CreateSessionRequest) (*models.Session, error) {
	slug, err := generateSlug()
	if err != nil {
		return nil, err
	}

	var s models.Session
	err = r.db.QueryRow(ctx,
		`INSERT INTO sessions (name, start_time, end_time, slug, time_zone)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, name, start_time, end_time, slug, time_zone`,
		req.Name, req.StartTime, req.EndTime, slug, req.TimeZone,
	).Scan(&s.ID, &s.Name, &s.StartTime, &s.EndTime, &s.Slug, &s.TimeZone)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	return &s, nil
}

func (r *sessionRepo) GetByID(ctx context.Context, id string) (*models.Session, error) {
	var s models.Session
	err := r.db.QueryRow(ctx,
		`SELECT id, name, start_time, end_time, slug, time_zone
		 FROM sessions WHERE id = $1`,
		id,
	).Scan(&s.ID, &s.Name, &s.StartTime, &s.EndTime, &s.Slug, &s.TimeZone)
	if err != nil {
		return nil, fmt.Errorf("get session: %w", err)
	}
	return &s, nil
}

func (r *sessionRepo) Update(ctx context.Context, id string, req models.UpdateSessionRequest) (*models.Session, error) {
	var s models.Session
	err := r.db.QueryRow(ctx,
		`UPDATE sessions
		 SET start_time = COALESCE($1, start_time),
		     end_time   = COALESCE($2, end_time),
		     updated_at = now()
		 WHERE id = $3
		 RETURNING id, name, start_time, end_time, slug, time_zone`,
		req.StartTime, req.EndTime, id,
	).Scan(&s.ID, &s.Name, &s.StartTime, &s.EndTime, &s.Slug, &s.TimeZone)
	if err != nil {
		return nil, fmt.Errorf("update session: %w", err)
	}
	return &s, nil
}

func (r *sessionRepo) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM sessions WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}
	return nil
}
