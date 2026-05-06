package postgres

import (
	"context"
	"fmt"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type availabilityRepo struct {
	db *pgxpool.Pool
}

func NewAvailabilityRepository(db *pgxpool.Pool) *availabilityRepo {
	return &availabilityRepo{db: db}
}

func (r *availabilityRepo) CreateAvailability(ctx context.Context, req models.CreateAvailabilityRequest) (*models.Availability, error) {
	var a models.Availability
	err := r.db.QueryRow(ctx,
		`INSERT INTO availability (slug, name, local_timezone, slots_by_date, hero_class)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, slug, name, local_timezone, slots_by_date, hero_class`,
		req.Slug, req.Name, req.LocalTimezone, req.SlotsByDate, req.HeroClass,
	).Scan(&a.ID, &a.Slug, &a.Name, &a.LocalTimezone, &a.SlotsByDate, &a.HeroClass)
	if err != nil {
		return nil, fmt.Errorf("create availability: %w", err)
	}
	return &a, nil
}

func (r *availabilityRepo) UpdateAvailability(ctx context.Context, req models.UpdateAvailabilityRequest) (*models.Availability, error) {
	var a models.Availability
	err := r.db.QueryRow(ctx,
		`UPDATE availability
		 SET slots_by_date = $1
		 WHERE id = $2
		 RETURNING id, slug, name, local_timezone, slots_by_date, hero_class`,
		req.SlotsByDate, req.Id,
	).Scan(&a.ID, &a.Slug, &a.Name, &a.LocalTimezone, &a.SlotsByDate, &a.HeroClass)
	if err != nil {
		return nil, fmt.Errorf("update availability: %w", err)
	}
	return &a, nil
}

func (r *availabilityRepo) GetAllAvailabilitiesForSlug(ctx context.Context, slug string) ([]models.Availability, error) {
	rows, err := r.db.Query(ctx,
		`SELECT id, slug, name, local_timezone, slots_by_date, hero_class
		 FROM availability WHERE slug = $1`,
		slug,
	)
	if err != nil {
		return nil, fmt.Errorf("get all availabilities: %w", err)
	}
	defer rows.Close()

	var availabilities []models.Availability
	for rows.Next() {
		var a models.Availability
		if err := rows.Scan(&a.ID, &a.Slug, &a.Name, &a.LocalTimezone, &a.SlotsByDate, &a.HeroClass); err != nil {
			return nil, fmt.Errorf("scan availability: %w", err)
		}
		availabilities = append(availabilities, a)
	}
	return availabilities, rows.Err()
}
