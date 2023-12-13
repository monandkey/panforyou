package controller

import (
	"context"

	"github.com/monandkey/panforyou/internal/pkg/usecase"
	"github.com/monandkey/panforyou/internal/pkg/usecase/model"
)

type GraphQLController struct {
	uc usecase.FindBreadUsecase
}

func NewGraphQLController(uc usecase.FindBreadUsecase) GraphQLController {
	return GraphQLController{
		uc: uc,
	}
}

func (c GraphQLController) List(ctx context.Context) ([]*model.Bread, error) {
	return c.uc.List(ctx)
}

func (c GraphQLController) GetByID(ctx context.Context, id string) (*model.Bread, error) {
	input := usecase.FindBreadInput{
		EntryID: id,
	}
	return c.uc.GetByID(ctx, input)
}
