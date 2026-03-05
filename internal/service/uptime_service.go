package service

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/MarcelloBB/ticker/internal/dto"
	"github.com/MarcelloBB/ticker/internal/model"
	"github.com/MarcelloBB/ticker/internal/repository"
	"gorm.io/gorm"
)

var (
	ErrInvalidURL      = errors.New("invalid url")
	ErrTargetNotFound  = errors.New("uptime target not found")
	ErrUnexpectedProbe = errors.New("unexpected error during probe")
)

type UptimeService struct {
	repo       repository.UptimeRepository
	httpClient *http.Client
}

func NewUptimeService(repo repository.UptimeRepository) *UptimeService {
	return &UptimeService{
		repo: repo,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (s *UptimeService) CreateTarget(ctx context.Context, req dto.CreateUptimeTargetRequest) (*dto.UptimeTargetResponse, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.URL = strings.TrimSpace(req.URL)

	if req.ExpectedStatus == 0 {
		req.ExpectedStatus = http.StatusOK
	}

	if err := validateURL(req.URL); err != nil {
		return nil, err
	}

	target := &model.UptimeTarget{
		Name:           req.Name,
		URL:            req.URL,
		ExpectedStatus: req.ExpectedStatus,
	}

	if err := s.repo.Create(ctx, target); err != nil {
		return nil, err
	}

	response := mapTargetResponse(*target)
	return &response, nil
}

func (s *UptimeService) ListTargets(ctx context.Context) ([]dto.UptimeTargetResponse, error) {
	targets, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dto.UptimeTargetResponse, 0, len(targets))
	for _, target := range targets {
		result = append(result, mapTargetResponse(target))
	}

	return result, nil
}

func (s *UptimeService) CheckTarget(ctx context.Context, id uint) (*dto.UptimeCheckResponse, error) {
	target, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTargetNotFound
		}
		return nil, err
	}

	start := time.Now()
	resp, err := s.httpClient.Get(target.URL)
	if err != nil {
		return nil, ErrUnexpectedProbe
	}
	defer resp.Body.Close()

	responseMS := time.Since(start).Milliseconds()
	isUp := resp.StatusCode == target.ExpectedStatus
	checkedAt := time.Now().UTC()

	if err := s.repo.UpdateProbe(ctx, target.ID, resp.StatusCode, responseMS, isUp, checkedAt); err != nil {
		return nil, err
	}

	target.LastStatusCode = &resp.StatusCode
	target.LastResponseMS = &responseMS
	target.LastCheckedAt = &checkedAt
	target.IsUp = isUp

	result := dto.UptimeCheckResponse{
		Target:         mapTargetResponse(*target),
		ObservedStatus: resp.StatusCode,
		ResponseMS:     responseMS,
		IsUp:           isUp,
	}

	return &result, nil
}

func validateURL(rawURL string) error {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return ErrInvalidURL
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return ErrInvalidURL
	}
	if parsedURL.Host == "" {
		return ErrInvalidURL
	}
	return nil
}

func mapTargetResponse(target model.UptimeTarget) dto.UptimeTargetResponse {
	return dto.UptimeTargetResponse{
		ID:             target.ID,
		Name:           target.Name,
		URL:            target.URL,
		ExpectedStatus: target.ExpectedStatus,
		LastStatusCode: target.LastStatusCode,
		LastCheckedAt:  target.LastCheckedAt,
		LastResponseMS: target.LastResponseMS,
		IsUp:           target.IsUp,
		CreatedAt:      target.CreatedAt,
		UpdatedAt:      target.UpdatedAt,
	}
}
