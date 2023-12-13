package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrBreadNotFound = errors.New("bread not found")
)

type BreadID string

func (b BreadID) String() string {
	return string(b)
}

type BreadRepository interface {
	Create(context.Context, Bread) error
	UpdateByID(context.Context, Bread) (Bread, error)
	FindAll(context.Context) ([]Bread, error)
	FindByID(context.Context, BreadID) (Bread, error)
	DeleteByID(context.Context, BreadID) error
}

type Bread struct {
	id        BreadID
	name      string
	createdAt time.Time
}

func NewBread(id BreadID, name string, createdAt time.Time) Bread {
	return Bread{
		id:        id,
		name:      name,
		createdAt: createdAt,
	}
}

func (b Bread) GetID() BreadID {
	return b.id
}

func (b Bread) GetName() string {
	return b.name
}

func (b Bread) GetCreatedAt() time.Time {
	return b.createdAt
}
