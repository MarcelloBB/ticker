package repository

import (
	"context"
	"time"

	"github.com/MarcelloBB/ticker/internal/model"
)

type UptimeRepository interface {
	Create(ctx context.Context, target *model.UptimeTarget) error
	List(ctx context.Context) ([]model.UptimeTarget, error)
	GetByID(ctx context.Context, id uint) (*model.UptimeTarget, error)
	UpdateProbe(ctx context.Context, id uint, statusCode int, responseMS int64, isUp bool, checkedAt time.Time) error
}
