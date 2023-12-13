package controller

import (
	"context"

	"github.com/monandkey/panforyou/internal/pkg/usecase"
)

type CommandController struct {
	cuc     usecase.FindContentfulUsecase
	buc     usecase.CreateBreadUsecase
	entryID string
}

func NewCommandController(
	cuc usecase.FindContentfulUsecase,
	buc usecase.CreateBreadUsecase,
	entryID string,
) CommandController {
	return CommandController{
		cuc:     cuc,
		buc:     buc,
		entryID: entryID,
	}
}

func (c CommandController) Create() error {
	ctx := context.Background()
	cinput := usecase.FindContentfulInput{
		EntryID: c.entryID,
	}

	bread, err := c.cuc.GetByID(cinput)
	if err != nil {
		return err
	}

	binput := usecase.CreateBreadInput{
		ID:        bread.ID,
		Name:      bread.Name,
		CreatedAt: bread.CreatedAt,
	}
	if err := c.buc.Create(ctx, binput); err != nil {
		return err
	}
	return nil
}
