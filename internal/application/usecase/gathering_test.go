package usecase_test

import (
	"context"
	"testing"

	"github.com/hieronimusbudi/simple-go-api/internal/application/usecase"
	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/helpers"
	"github.com/hieronimusbudi/simple-go-api/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_gatheringUsecase_Create(t *testing.T) {
	gathering := domain.Gathering{
		Creator: domain.Member{
			ID: 1,
		},
		Type:        0,
		ScheduledAt: "2020-10-06 11:53",
		Name:        "Private Gathering",
		Location:    "Local Street",
		Attendees: []domain.Member{
			{
				ID: 1,
			},
		},
	}
	wantNewGathering := gathering
	wantNewGathering.ID = 1
	type args struct {
		gathering domain.Gathering
	}
	tests := []struct {
		name             string
		args             args
		wantNewGathering domain.Gathering
		wantErr          bool
		funcCreate       helpers.TestFuncCall
		funcGet          helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{
				gathering: gathering,
			},
			wantNewGathering: wantNewGathering,
			funcCreate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{wantNewGathering.ID, nil},
			},
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{[]domain.Gathering{wantNewGathering}, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGathering := new(mocks.IGathering)
			usecase := usecase.NewGatheringUsecase(usecase.GatheringUsecaseArgs{
				GatheringRepository: mockGathering,
			})
			if tt.funcCreate.Called {
				mockGathering.On("Create", tt.funcCreate.Input...).Return(tt.funcCreate.Output...)
			}
			if tt.funcGet.Called {
				mockGathering.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotNewGathering, err := usecase.Create(context.Background(), tt.args.gathering)
			if (err != nil) != tt.wantErr {
				t.Errorf("gatheringUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantNewGathering, gotNewGathering)
			}
		})
	}
}

func Test_gatheringUsecase_Get(t *testing.T) {
	gatherings := []domain.Gathering{
		{
			ID:        1,
			CreatorID: 1,
			Creator: domain.Member{
				ID: 1,
			},
			Type:        0,
			ScheduledAt: "2020-10-06T04:53:02Z",
			CreatedAt:   "2023-09-28T19:09:41Z",
			Name:        "Private Meeting",
			Location:    "pramuka street",
			Attendees: []domain.Member{
				{
					ID: 1,
				},
			},
		},
	}
	type args struct {
		args domain.GatheringArgs
	}
	tests := []struct {
		name           string
		args           args
		wantGatherings []domain.Gathering
		wantErr        bool
		funcGet        helpers.TestFuncCall
	}{
		{
			name:           "success",
			wantGatherings: gatherings,
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{gatherings, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGathering := new(mocks.IGathering)
			usecase := usecase.NewGatheringUsecase(usecase.GatheringUsecaseArgs{
				GatheringRepository: mockGathering,
			})
			if tt.funcGet.Called {
				mockGathering.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotGatherings, err := usecase.Get(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantGatherings, gotGatherings)
			}
		})
	}
}

func Test_gatheringUsecase_GetByID(t *testing.T) {
	gatherings := []domain.Gathering{
		{
			ID:        1,
			CreatorID: 1,
			Creator: domain.Member{
				ID: 1,
			},
			Type:        0,
			ScheduledAt: "2020-10-06T04:53:02Z",
			CreatedAt:   "2023-09-28T19:09:41Z",
			Name:        "Private Meeting",
			Location:    "pramuka street",
			Attendees: []domain.Member{
				{
					ID: 1,
				},
			},
		},
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name          string
		args          args
		wantGathering domain.Gathering
		wantErr       bool
		funcGet       helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantGathering: gatherings[0],
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{gatherings, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGathering := new(mocks.IGathering)
			usecase := usecase.NewGatheringUsecase(usecase.GatheringUsecaseArgs{
				GatheringRepository: mockGathering,
			})
			if tt.funcGet.Called {
				mockGathering.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotGathering, err := usecase.GetByID(context.Background(), tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantGathering, gotGathering)
			}
		})
	}
}

func Test_gatheringUsecase_Update(t *testing.T) {
	gathering := domain.Gathering{
		ID:        1,
		CreatorID: 1,
		Creator: domain.Member{
			ID: 1,
		},
		Type:        0,
		ScheduledAt: "2023-10-06 04:53",
		CreatedAt:   "2023-09-28T19:09:41Z",
		Name:        "Update Private Meeting",
		Location:    "update pramuka street",
		Attendees: []domain.Member{
			{
				ID: 1,
			},
		},
	}
	type args struct {
		gathering domain.Gathering
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		funcUpdate helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{
				gathering: gathering,
			},
			funcUpdate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGathering := new(mocks.IGathering)
			usecase := usecase.NewGatheringUsecase(usecase.GatheringUsecaseArgs{
				GatheringRepository: mockGathering,
			})
			if tt.funcUpdate.Called {
				mockGathering.On("Update", tt.funcUpdate.Input...).Return(tt.funcUpdate.Output...)
			}
			err := usecase.Update(context.Background(), tt.args.gathering)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_gatheringUsecase_Delete(t *testing.T) {
	type args struct {
		args domain.GatheringArgs
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		funcDelete helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{domain.GatheringArgs{
				ID: int64(1),
			}},
			funcDelete: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{nil},
			},
		},
	}
	for _, tt := range tests {
		mockGathering := new(mocks.IGathering)
		usecase := usecase.NewGatheringUsecase(usecase.GatheringUsecaseArgs{
			GatheringRepository: mockGathering,
		})
		if tt.funcDelete.Called {
			mockGathering.On("Delete", tt.funcDelete.Input...).Return(tt.funcDelete.Output...)
		}
		err := usecase.Delete(context.Background(), tt.args.args)
		if tt.wantErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
