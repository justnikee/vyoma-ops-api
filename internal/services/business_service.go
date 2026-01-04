package services

import (
	"context"
	"errors"

	"vyoma-api/internal/models"
	"vyoma-api/internal/repositories"

	"github.com/google/uuid"
)

type BusinessService struct {
	repo *repositories.BusinessRepository
}

func NewBusinessService(repo *repositories.BusinessRepository) *BusinessService {
	return &BusinessService{repo: repo}
}

func (s *BusinessService) CreateBusiness(
	ctx context.Context,
	business *models.Business,
) (*models.Business, error) {

	// basic validation
	if business.Name == "" {
		return nil, errors.New("business name is required")
	}
	if business.State == "" {
		return nil, errors.New("state is required")
	}
	if business.IndustryType == "" {
		return nil, errors.New("industry type is required")
	}

	// create business in DB
	createdBusiness, err := s.repo.CreateBusiness(ctx, business)
	if err != nil {
		return nil, err
	}

	// 🔜 later: auto-assign compliance rules here

	return createdBusiness, nil
}

func (s *BusinessService) GetBusinessByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Business, error) {

	if id == uuid.Nil {
		return nil, errors.New("invalid business id")
	}

	return s.repo.GetBusinessByID(ctx, id)
}
