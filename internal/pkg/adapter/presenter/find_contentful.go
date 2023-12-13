package presenter

import (
	"time"

	"github.com/monandkey/panforyou/internal/pkg/usecase"
)

type findContentfulPresenter struct{}

func NewFindContentfulPresenter() usecase.FindContentfulPresenter {
	return findContentfulPresenter{}
}

func (f findContentfulPresenter) Output(vecString []string) usecase.FindContentfulOutput {
	if len(vecString) != 3 {
		return usecase.FindContentfulOutput{}
	}

	id := vecString[0]
	name := vecString[1]
	createdAt, _ := time.Parse("2006-01-02T15:04:05Z07:00", vecString[2])
	return usecase.FindContentfulOutput{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
	}
}
