package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/repository"
)

type (
	gatheringUsecase struct {
		gatheringRepository repository.IGathering
	}

	GatheringUsecaseArgs struct {
		GatheringRepository repository.IGathering
	}

	IGatheringUsecase interface {
		Create(ctx context.Context, gathering domain.Gathering) (NewGathering domain.Gathering, err error)
		Get(ctx context.Context, args domain.GatheringArgs) (gatherings []domain.Gathering, err error)
		GetByID(ctx context.Context, id int64) (gathering domain.Gathering, err error)
		Update(ctx context.Context, gathering domain.Gathering) (err error)
		Delete(ctx context.Context, args domain.GatheringArgs) (err error)
	}
)

func NewGatheringUsecase(args GatheringUsecaseArgs) IGatheringUsecase {
	return &gatheringUsecase{
		gatheringRepository: args.GatheringRepository,
	}
}

func (u *gatheringUsecase) Create(ctx context.Context, gathering domain.Gathering) (NewGathering domain.Gathering, err error) {
	id, err := u.gatheringRepository.Create(ctx, gathering)
	if err != nil {
		log.Println(err)
		return
	}
	NewGathering, err = u.GetByID(ctx, id)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *gatheringUsecase) Get(ctx context.Context, args domain.GatheringArgs) (gatherings []domain.Gathering, err error) {
	gatherings, err = u.gatheringRepository.Get(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *gatheringUsecase) GetByID(ctx context.Context, id int64) (gathering domain.Gathering, err error) {
	gatherings, err := u.gatheringRepository.Get(ctx, domain.GatheringArgs{IDs: []int64{id}})
	if err != nil {
		log.Println(err)
		return
	}
	if len(gatherings) == 0 {
		err = errors.New("cannot find gathering")
		return
	}
	gathering = gatherings[0]
	return
}

func (u *gatheringUsecase) Update(ctx context.Context, gathering domain.Gathering) (err error) {
	err = u.gatheringRepository.Update(ctx, gathering)
	if err != nil {
		log.Println(err)
	}
	return
}

func (u *gatheringUsecase) Delete(ctx context.Context, args domain.GatheringArgs) (err error) {
	err = u.gatheringRepository.Delete(ctx, args)
	if err != nil {
		log.Println(err)
	}
	return
}
