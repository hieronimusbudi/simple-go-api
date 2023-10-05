package repository_test

import (
	"context"
	"testing"

	"github.com/hieronimusbudi/simple-go-api/internal/adapter/repository"
	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
	"github.com/stretchr/testify/require"
)

// This test is integration test
func Test_invitationAdapterRepository_Create(t *testing.T) {
	invitation := domain.Invitation{
		Member: domain.Member{
			ID: 2,
		},
		Gathering: domain.Gathering{
			ID: 2,
		},
	}
	type args struct {
		invitation domain.Invitation
	}
	tests := []struct {
		name    string
		args    args
		wantId  int64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				invitation: invitation,
			},
			wantId: 2, // based on last ID from test data.sql
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewInvitationRepository(repository.InvitationAdapterRepositoryArgs{
				DB: db,
			})
			gotId, err := repo.Create(context.Background(), tt.args.invitation)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantId, gotId)
			}
		})
	}
}

func Test_invitationAdapterRepository_Get(t *testing.T) {
	invitations := []domain.Invitation{
		{
			ID:       1,
			MemberID: 2,
			Member: domain.Member{
				ID: 2,
			},
			GatheringID: 1,
			Gathering: domain.Gathering{
				ID: 1,
			},
			Status:    valueobject.INVITATION_CREATED,
			CreatedAt: "2023-10-02T11:09:22Z",
		},
	}
	type args struct {
		args domain.InvitationArgs
	}
	tests := []struct {
		name            string
		args            args
		wantInvitations []domain.Invitation
		wantErr         bool
	}{
		{
			name: "success",
			args: args{
				domain.InvitationArgs{
					IDs: []int64{1},
				},
			},
			wantInvitations: invitations,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewInvitationRepository(repository.InvitationAdapterRepositoryArgs{
				DB: db,
			})
			gotInvitations, err := repo.Get(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantInvitations, gotInvitations)
			}
		})
	}
}

func Test_invitationAdapterRepository_UpdateStatus(t *testing.T) {
	invitationAccept := domain.Invitation{
		ID:       1,
		MemberID: 2,
		Member: domain.Member{
			ID: 2,
		},
		GatheringID: 1,
		Gathering: domain.Gathering{
			ID: 1,
		},
		Status:    valueobject.INVITATION_ACCEPT,
		CreatedAt: "2023-10-02T11:09:22Z",
	}
	invitationReject := domain.Invitation{
		ID:       1,
		MemberID: 2,
		Member: domain.Member{
			ID: 2,
		},
		GatheringID: 1,
		Gathering: domain.Gathering{
			ID: 1,
		},
		Status:    valueobject.INVITATION_REJECT,
		CreatedAt: "2023-10-02T11:09:22Z",
	}
	type args struct {
		args domain.InvitationArgs
	}
	tests := []struct {
		name           string
		args           args
		wantInvitation domain.Invitation
		wantErr        bool
	}{
		{
			name: "success accept",
			args: args{
				domain.InvitationArgs{
					ID:          1,
					MemberID:    invitationAccept.MemberID,
					GatheringID: invitationAccept.GatheringID,
					Status:      valueobject.INVITATION_ACCEPT,
				},
			},
			wantInvitation: invitationAccept,
		},
		{
			name: "success reject",
			args: args{
				domain.InvitationArgs{
					ID:          1,
					MemberID:    invitationReject.MemberID,
					GatheringID: invitationReject.GatheringID,
					Status:      valueobject.INVITATION_REJECT,
				},
			},
			wantInvitation: invitationReject,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewInvitationRepository(repository.InvitationAdapterRepositoryArgs{
				DB: db,
			})
			err := repo.UpdateStatus(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// check data
				invitations, err := repo.Get(context.Background(), domain.InvitationArgs{
					IDs: []int64{tt.args.args.ID},
				})
				require.NoError(t, err)
				gotInvitation := invitations[0]
				require.Equal(t, tt.wantInvitation, gotInvitation)
			}
		})
	}
}
