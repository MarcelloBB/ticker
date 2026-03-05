package dto

import "time"

type CreateUptimeTargetRequest struct {
	Name           string `json:"name" binding:"required"`
	URL            string `json:"url" binding:"required"`
	ExpectedStatus int    `json:"expected_status"`
}

type UptimeTargetResponse struct {
	ID             uint       `json:"id"`
	Name           string     `json:"name"`
	URL            string     `json:"url"`
	ExpectedStatus int        `json:"expected_status"`
	LastStatusCode *int       `json:"last_status_code,omitempty"`
	LastCheckedAt  *time.Time `json:"last_checked_at,omitempty"`
	LastResponseMS *int64     `json:"last_response_ms,omitempty"`
	IsUp           bool       `json:"is_up"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type UptimeCheckResponse struct {
	Target         UptimeTargetResponse `json:"target"`
	ObservedStatus int                  `json:"observed_status"`
	ResponseMS     int64                `json:"response_ms"`
	IsUp           bool                 `json:"is_up"`
}
