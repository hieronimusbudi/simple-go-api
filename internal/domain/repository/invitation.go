package repository

import (
	"context"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
)

type IInvitation interface {
	Create(ctx context.Context, invitation domain.Invitation) (ID int64, err error)
	Get(ctx context.Context, args domain.InvitationArgs) (invitations []domain.Invitation, err error)
	UpdateStatus(ctx context.Context, args domain.InvitationArgs) (err error)
}
