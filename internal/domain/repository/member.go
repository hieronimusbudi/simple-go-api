package repository

import (
	"context"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
)

type IMember interface {
	Create(ctx context.Context, member domain.Member) (id int64, err error)
	Get(ctx context.Context, args domain.MemberArgs) (members []domain.Member, err error)
	Update(ctx context.Context, member domain.Member) (err error)
	Delete(ctx context.Context, args domain.MemberArgs) (err error)
}
