package usecase

import (
	"context"

	"github.com/monandkey/panforyou/internal/pkg/domain"
	"github.com/monandkey/panforyou/internal/pkg/usecase/model"
)

type FindBreadUsecase interface {
	List(context.Context) ([]*model.Bread, error)
	GetByID(context.Context, FindBreadInput) (*model.Bread, error)
}

type FindBreadInput struct {
	EntryID string
}

type FindBreadPresenter interface {
	Output(domain.Bread) *model.Bread
}

type findBreadInteractor struct {
	repo      domain.BreadRepository
	presenter FindBreadPresenter
}

func NewFindBreadUsecase(
	repo domain.BreadRepository,
	presenter FindBreadPresenter,
) FindBreadUsecase {
	return findBreadInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (f findBreadInteractor) List(ctx context.Context) ([]*model.Bread, error) {
	breads, err := f.repo.FindAll(ctx)
	if err != nil {
		return []*model.Bread{}, err
	}

	var outputBreads []*model.Bread
	for _, v := range breads {
		o := f.presenter.Output(v)
		outputBreads = append(outputBreads, o)
	}
	return outputBreads, nil
}

func (f findBreadInteractor) GetByID(ctx context.Context, input FindBreadInput) (*model.Bread, error) {
	bread, err := f.repo.FindByID(
		ctx,
		domain.BreadID(input.EntryID),
	)
	if err != nil {
		return &model.Bread{}, err
	}
	return f.presenter.Output(bread), nil
}
