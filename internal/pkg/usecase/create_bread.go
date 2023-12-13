package usecase

import (
	"context"
	"time"

	"github.com/monandkey/panforyou/internal/pkg/domain"
)

type CreateBreadUsecase interface {
	Create(context.Context, CreateBreadInput) error
}

type CreateBreadInput struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type createBreadInteractor struct {
	repo       domain.BreadRepository
	ctxTimeout time.Duration
}

func NewCreateBreadUsecase(
	repo domain.BreadRepository,
	t time.Duration,
) CreateBreadUsecase {
	return createBreadInteractor{
		repo:       repo,
		ctxTimeout: t,
	}
}

func (c createBreadInteractor) Create(ctx context.Context, input CreateBreadInput) error {
	ctx, cancel := context.WithTimeout(ctx, c.ctxTimeout)
	defer cancel()

	bread := domain.NewBread(
		domain.BreadID(input.ID),
		input.Name,
		input.CreatedAt,
	)

	if err := c.repo.Create(ctx, bread); err != nil {
		return err
	}
	return nil
}
