package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/repository"
)

type (
	memberUsecase struct {
		memberRepository repository.IMember
	}

	MemberUsecaseArgs struct {
		MemberRepository repository.IMember
	}

	IMemberUsecase interface {
		Create(ctx context.Context, member domain.Member) (newMember domain.Member, err error)
		Get(ctx context.Context, args domain.MemberArgs) (members []domain.Member, err error)
		GetByID(ctx context.Context, id int64) (member domain.Member, err error)
		Update(ctx context.Context, member domain.Member) (err error)
		Delete(ctx context.Context, args domain.MemberArgs) (err error)
	}
)

func NewMemberUsecase(args MemberUsecaseArgs) IMemberUsecase {
	return &memberUsecase{
		memberRepository: args.MemberRepository,
	}
}

func (u *memberUsecase) Create(ctx context.Context, member domain.Member) (newMember domain.Member, err error) {
	id, err := u.memberRepository.Create(ctx, member)
	if err != nil {
		log.Println(err)
		return
	}
	newMember, err = u.GetByID(ctx, id)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *memberUsecase) Get(ctx context.Context, args domain.MemberArgs) (members []domain.Member, err error) {
	members, err = u.memberRepository.Get(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *memberUsecase) GetByID(ctx context.Context, id int64) (member domain.Member, err error) {
	members, err := u.memberRepository.Get(ctx, domain.MemberArgs{IDs: []int64{id}})
	if err != nil {
		log.Println(err)
		return
	}
	if len(members) == 0 {
		err = errors.New("cannot find member")
		return
	}
	member = members[0]
	return
}

func (u *memberUsecase) Update(ctx context.Context, member domain.Member) (err error) {
	err = u.memberRepository.Update(ctx, member)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *memberUsecase) Delete(ctx context.Context, args domain.MemberArgs) (err error) {
	err = u.memberRepository.Delete(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}
