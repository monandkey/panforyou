package repository

import (
	"context"
	"time"

	"github.com/monandkey/panforyou/internal/pkg/domain"
)

type BreadSQL struct {
	db SQL
}

func NewBreadSQL(db SQL) domain.BreadRepository {
	return BreadSQL{
		db: db,
	}
}

func (b BreadSQL) Create(ctx context.Context, bread domain.Bread) error {
	query := `
		INSERT INTO
			breads (id, name, createdAt)
		VALUES
			($1, $2, $3);
	`

	return b.db.ExecuteContext(
		ctx,
		query,
		bread.GetID(),
		bread.GetName(),
		bread.GetCreatedAt(),
	)
}

func (b BreadSQL) FindAll(ctx context.Context) ([]domain.Bread, error) {
	query := "SELECT * FROM breads;"
	rows, err := b.db.QueryContext(ctx, query)
	if err != nil {
		return []domain.Bread{}, err
	}

	var breads = make([]domain.Bread, 0)
	for rows.Next() {
		var (
			id        string
			name      string
			createdAt time.Time
		)

		if err = rows.Scan(&id, &name, &createdAt); err != nil {
			return []domain.Bread{}, err
		}
		breads = append(breads, domain.NewBread(
			domain.BreadID(id),
			name,
			createdAt,
		))
	}
	defer rows.Close()
	if err = rows.Err(); err != nil {
		return []domain.Bread{}, err
	}
	return breads, nil
}

func (b BreadSQL) UpdateByID(ctx context.Context, bread domain.Bread) (domain.Bread, error) {
	query := `
		INSERT INTO
			breads (id, name, createdAt)
		VALUES
			($1, $2, $3)
		ON CONFLICT(id) DO UPDATE SET
			id = EXCLUDED.id,
			name = EXCLUDED.name,
			createdAt = EXCLUDED.createdAt;
	`

	if err := b.db.ExecuteContext(
		ctx,
		query,
		bread.GetID(),
		bread.GetName(),
		bread.GetCreatedAt(),
	); err != nil {
		return domain.Bread{}, err
	}
	return bread, nil
}

func (b BreadSQL) FindByID(ctx context.Context, id domain.BreadID) (domain.Bread, error) {
	query := "SELECT * FROM breads WHERE id = $1;"
	rows, err := b.db.QueryContext(ctx, query, id)
	if err != nil {
		return domain.Bread{}, err
	}
	var breads = make([]domain.Bread, 0)
	for rows.Next() {
		var (
			id        string
			name      string
			createdAt time.Time
		)

		if err = rows.Scan(&id, &name, &createdAt); err != nil {
			return domain.Bread{}, err
		}
		breads = append(breads, domain.NewBread(
			domain.BreadID(id),
			name,
			createdAt,
		))
	}
	defer rows.Close()
	if err = rows.Err(); err != nil {
		return domain.Bread{}, err
	}
	return breads[0], nil
}

func (b BreadSQL) DeleteByID(ctx context.Context, id domain.BreadID) error {
	query := "SELECT FROM breads WHERE id = $1;"
	return b.db.ExecuteContext(ctx, query, id)
}
