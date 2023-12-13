package usecase

import (
	"time"

	"github.com/monandkey/panforyou/internal/pkg/domain"
)

type FindContentfulUsecase interface {
	GetByID(FindContentfulInput) (FindContentfulOutput, error)
}

type FindContentfulInput struct {
	EntryID string
}

type FindContentfulPresenter interface {
	Output([]string) FindContentfulOutput
}

type FindContentfulOutput struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type findContentfulInteractor struct {
	repo      domain.ContentfulRepository
	presenter FindContentfulPresenter
}

func NewFindContentfulUsecase(
	repo domain.ContentfulRepository,
	presenter FindContentfulPresenter,
) FindContentfulUsecase {
	return findContentfulInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (f findContentfulInteractor) GetByID(input FindContentfulInput) (FindContentfulOutput, error) {
	vecStr, err := f.repo.FindByID(
		input.EntryID,
	)
	if err != nil {
		return FindContentfulOutput{}, err
	}
	return f.presenter.Output(vecStr), nil
}
