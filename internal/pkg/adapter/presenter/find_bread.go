package presenter

import (
	"github.com/monandkey/panforyou/internal/pkg/domain"
	"github.com/monandkey/panforyou/internal/pkg/usecase"
	"github.com/monandkey/panforyou/internal/pkg/usecase/model"
)

type findBreadPresenter struct{}

func NewFindBreadPresenter() usecase.FindBreadPresenter {
	return findBreadPresenter{}
}

func (b findBreadPresenter) Output(bread domain.Bread) *model.Bread {
	var (
		id        = bread.GetID().String()
		name      = bread.GetName()
		createdAt = bread.GetCreatedAt()
	)
	return &model.Bread{
		ID:        &id,
		Name:      &name,
		CreatedAt: &createdAt,
	}
}
