package models

type Availability struct {
	ID            string                    `json:"id"`
	Slug          string                    `json:"slug"`
	Name          string                    `json:"name"`
	LocalTimezone string                    `json:"localTimezone"`
	SlotsByDate   map[string]map[int]string `json:"slotsByDate"`
	HeroClass     string                    `json:"heroClass"`
}

type CreateAvailabilityRequest struct {
	Slug          string                    `json:"slug"`
	Name          string                    `json:"name"`
	LocalTimezone string                    `json:"localTimezone"`
	SlotsByDate   map[string]map[int]string `json:"slotsByDate"`
	HeroClass     string                    `json:"heroClass"`
}

type UpdateAvailabilityRequest struct {
	Name          string                    `json:"name"`
	LocalTimezone string                    `json:"localTimezone"`
	SlotsByDate   map[string]map[int]string `json:"slotsByDate"`
	HeroClass     string                    `json:"heroClass"`
	Id            string                    `json:"id"`
}
