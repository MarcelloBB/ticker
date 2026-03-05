package repository

import (
	"context"
	"time"

	"github.com/MarcelloBB/ticker/internal/model"
	"gorm.io/gorm"
)

type UptimeGormRepository struct {
	db *gorm.DB
}

func NewUptimeGormRepository(db *gorm.DB) UptimeRepository {
	return &UptimeGormRepository{db: db}
}

func (r *UptimeGormRepository) Create(ctx context.Context, target *model.UptimeTarget) error {
	return r.db.WithContext(ctx).Create(target).Error
}

func (r *UptimeGormRepository) List(ctx context.Context) ([]model.UptimeTarget, error) {
	var targets []model.UptimeTarget
	err := r.db.WithContext(ctx).Order("id DESC").Find(&targets).Error
	return targets, err
}

func (r *UptimeGormRepository) GetByID(ctx context.Context, id uint) (*model.UptimeTarget, error) {
	var target model.UptimeTarget
	if err := r.db.WithContext(ctx).First(&target, id).Error; err != nil {
		return nil, err
	}
	return &target, nil
}

func (r *UptimeGormRepository) UpdateProbe(ctx context.Context, id uint, statusCode int, responseMS int64, isUp bool, checkedAt time.Time) error {
	updates := map[string]interface{}{
		"last_status_code": statusCode,
		"last_response_ms": responseMS,
		"last_checked_at":  checkedAt,
		"is_up":            isUp,
	}
	return r.db.WithContext(ctx).Model(&model.UptimeTarget{}).Where("id = ?", id).Updates(updates).Error
}
