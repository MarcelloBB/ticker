package model

import "time"

type UptimeTarget struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"size:120;not null" json:"name"`
	URL            string     `gorm:"size:2048;not null;uniqueIndex" json:"url"`
	ExpectedStatus int        `gorm:"not null;default:200" json:"expected_status"`
	LastStatusCode *int       `json:"last_status_code,omitempty"`
	LastCheckedAt  *time.Time `json:"last_checked_at,omitempty"`
	LastResponseMS *int64     `json:"last_response_ms,omitempty"`
	IsUp           bool       `gorm:"not null;default:false" json:"is_up"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
