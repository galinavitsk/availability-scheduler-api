package models

type Session struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	StartTime     string   `json:"startTime"`
	EndTime       string   `json:"endTime"`
	Slug          string   `json:"slug"`
	TimeZone      string   `json:"timeZone"`
	SelectedDates []string `json:"selectedDates"`
}

type CreateSessionRequest struct {
	Name          string   `json:"name" binding:"required"`
	StartTime     string   `json:"startTime" binding:"required"`
	EndTime       string   `json:"endTime" binding:"required"`
	TimeZone      string   `json:"timeZone" binding:"required"`
	SelectedDates []string `json:"selectedDates"`
}

type UpdateSessionRequest struct {
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
}
