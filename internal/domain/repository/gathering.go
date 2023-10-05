package repository

import (
	"context"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
)

type IGathering interface {
	Create(ctx context.Context, gathering domain.Gathering) (ID int64, err error)
	Get(ctx context.Context, args domain.GatheringArgs) (gatherings []domain.Gathering, err error)
	Update(ctx context.Context, gathering domain.Gathering) (err error)
	Delete(ctx context.Context, args domain.GatheringArgs) (err error)
}
