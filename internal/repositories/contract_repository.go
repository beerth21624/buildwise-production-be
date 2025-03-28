package repositories

import (
	"boonkosang/internal/domain/models"
	"context"

	"github.com/google/uuid"
)

type ContractRepository interface {
	Create(ctx context.Context, contract *models.Contract) error
	Update(ctx context.Context, contract *models.Contract) error

	Delete(ctx context.Context, projectID uuid.UUID) error
	GetByID(ctx context.Context, projectID uuid.UUID) (*models.Contract, error)
	GetByProjectID(ctx context.Context, projectID uuid.UUID) (*models.Contract, error)
	ValidateProjectStatus(ctx context.Context, projectID uuid.UUID) error

	ChangeStatus(ctx context.Context, projectID uuid.UUID, status string) error
}
