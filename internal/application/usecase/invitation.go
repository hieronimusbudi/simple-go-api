package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/repository"
)

type (
	invitationUsecase struct {
		invitationRepository repository.IInvitation
	}

	InvitationUsecaseArgs struct {
		InvitationRepository repository.IInvitation
	}

	IInvitationUsecase interface {
		Create(ctx context.Context, invitation domain.Invitation) (NewInvitation domain.Invitation, err error)
		Get(ctx context.Context, args domain.InvitationArgs) (invitations []domain.Invitation, err error)
		GetByID(ctx context.Context, id int64) (invitation domain.Invitation, err error)
		Accept(ctx context.Context, args domain.InvitationArgs) (err error)
		Reject(ctx context.Context, args domain.InvitationArgs) (err error)
		Cancel(ctx context.Context, args domain.InvitationArgs) (err error)
	}
)

func NewInvitationUsecase(args InvitationUsecaseArgs) IInvitationUsecase {
	return &invitationUsecase{
		invitationRepository: args.InvitationRepository,
	}
}

func (u *invitationUsecase) Create(ctx context.Context, invitation domain.Invitation) (NewInvitation domain.Invitation, err error) {
	id, err := u.invitationRepository.Create(ctx, invitation)
	if err != nil {
		log.Println(err)
		return
	}
	NewInvitation, err = u.GetByID(ctx, id)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *invitationUsecase) Get(ctx context.Context, args domain.InvitationArgs) (invitations []domain.Invitation, err error) {
	invitations, err = u.invitationRepository.Get(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *invitationUsecase) GetByID(ctx context.Context, id int64) (invitation domain.Invitation, err error) {
	invitations, err := u.invitationRepository.Get(ctx, domain.InvitationArgs{IDs: []int64{id}})
	if err != nil {
		log.Println(err)
		return
	}
	if len(invitations) == 0 {
		err = errors.New("cannot find invitation")
		return
	}
	invitation = invitations[0]
	return
}

func (u *invitationUsecase) Accept(ctx context.Context, args domain.InvitationArgs) (err error) {
	err = u.invitationRepository.UpdateStatus(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *invitationUsecase) Reject(ctx context.Context, args domain.InvitationArgs) (err error) {
	err = u.invitationRepository.UpdateStatus(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *invitationUsecase) Cancel(ctx context.Context, args domain.InvitationArgs) (err error) {
	err = u.invitationRepository.UpdateStatus(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}
