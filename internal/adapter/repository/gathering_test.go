package repository_test

import (
	"context"
	"testing"

	"github.com/hieronimusbudi/simple-go-api/internal/adapter/repository"
	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/stretchr/testify/require"
)

// This test is integration test
func Test_gatheringAdapterRepository_Create(t *testing.T) {
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
	type args struct {
		gathering domain.Gathering
	}
	tests := []struct {
		name    string
		args    args
		wantId  int64 // based on last ID from test data.sql
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				gathering: gathering,
			},
			wantId: 2, // based on last ID from test data.sql
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewGatheringRepository(repository.GatheringAdapterRepositoryArgs{
				DB: db,
			})
			gotId, err := repo.Create(context.Background(), tt.args.gathering)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantId, gotId)
			}
		})
	}
}

func Test_gatheringAdapterRepository_Get(t *testing.T) {
	gatherings := []domain.Gathering{
		{
			ID:        1,
			CreatorID: 1,
			Creator: domain.Member{
				ID: 1,
			},
			Type:        0,
			ScheduledAt: "2023-10-06T05:00:00Z",
			CreatedAt:   "2023-10-02T11:06:52Z",
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
	}{
		{
			name: "success",
			args: args{
				domain.GatheringArgs{
					IDs: []int64{1},
				},
			},
			wantGatherings: gatherings,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewGatheringRepository(repository.GatheringAdapterRepositoryArgs{
				DB: db,
			})
			gotGatherings, err := repo.Get(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantGatherings, gotGatherings)
			}
		})
	}
}

func Test_gatheringAdapterRepository_Update(t *testing.T) {
	gathering := domain.Gathering{
		ID:        1,
		CreatorID: 1,
		Creator: domain.Member{
			ID: 1,
		},
		Type:        0,
		ScheduledAt: "2023-10-06 04:53",
		CreatedAt:   "2023-10-02T11:06:52Z",
		Name:        "Update Private Meeting",
		Location:    "update pramuka street",
		Attendees: []domain.Member{
			{
				ID: 1,
			},
		},
	}
	wantGathering := gathering
	wantGathering.ScheduledAt = "2023-10-06T04:53:00Z"
	type args struct {
		gathering domain.Gathering
	}
	tests := []struct {
		name          string
		args          args
		wantGathering domain.Gathering
		wantErr       bool
	}{
		{
			name: "success",
			args: args{
				gathering: gathering,
			},
			wantGathering: wantGathering,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewGatheringRepository(repository.GatheringAdapterRepositoryArgs{
				DB: db,
			})
			err := repo.Update(context.Background(), tt.args.gathering)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// check data
				gatherings, err := repo.Get(context.Background(), domain.GatheringArgs{
					IDs: []int64{tt.args.gathering.ID},
				})
				require.NoError(t, err)
				gotGathering := gatherings[0]
				require.Equal(t, tt.wantGathering, gotGathering)
			}
		})
	}
}

func Test_gatheringAdapterRepository_Delete(t *testing.T) {
	type args struct {
		args domain.GatheringArgs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				args: domain.GatheringArgs{
					ID: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewGatheringRepository(repository.GatheringAdapterRepositoryArgs{
				DB: db,
			})
			err := repo.Delete(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// check data
				gatherings, err := repo.Get(context.Background(), domain.GatheringArgs{
					IDs: []int64{tt.args.args.ID},
				})
				require.NoError(t, err)
				require.Equal(t, 0, len(gatherings))
			}
		})
	}
}
