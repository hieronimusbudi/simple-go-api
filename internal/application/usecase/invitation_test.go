package usecase_test

import (
	"context"
	"testing"

	"github.com/hieronimusbudi/simple-go-api/internal/application/usecase"
	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/domain/valueobject"
	"github.com/hieronimusbudi/simple-go-api/internal/helpers"
	"github.com/hieronimusbudi/simple-go-api/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_invitationUsecase_Create(t *testing.T) {
	invitation := domain.Invitation{
		Member: domain.Member{
			ID: 2,
		},
		Gathering: domain.Gathering{
			ID: 2,
		},
	}
	wantNewInvitation := invitation
	wantNewInvitation.ID = 1
	type args struct {
		invitation domain.Invitation
	}
	tests := []struct {
		name              string
		args              args
		wantNewInvitation domain.Invitation
		wantErr           bool
		funcCreate        helpers.TestFuncCall
		funcGet           helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{
				invitation: invitation,
			},
			wantNewInvitation: wantNewInvitation,
			funcCreate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{wantNewInvitation.ID, nil},
			},
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{[]domain.Invitation{wantNewInvitation}, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvitation := new(mocks.IInvitation)
			usecase := usecase.NewInvitationUsecase(usecase.InvitationUsecaseArgs{
				InvitationRepository: mockInvitation,
			})
			if tt.funcCreate.Called {
				mockInvitation.On("Create", tt.funcCreate.Input...).Return(tt.funcCreate.Output...)
			}
			if tt.funcGet.Called {
				mockInvitation.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotNewInvitation, err := usecase.Create(context.Background(), tt.args.invitation)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantNewInvitation, gotNewInvitation)
			}
		})
	}
}

func Test_invitationUsecase_Get(t *testing.T) {
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
			CreatedAt: "2023-09-28T16:57:27Z",
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
		funcGet         helpers.TestFuncCall
	}{
		{
			name:            "success",
			wantInvitations: invitations,
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{invitations, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvitation := new(mocks.IInvitation)
			usecase := usecase.NewInvitationUsecase(usecase.InvitationUsecaseArgs{
				InvitationRepository: mockInvitation,
			})
			if tt.funcGet.Called {
				mockInvitation.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotInvitations, err := usecase.Get(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantInvitations, gotInvitations)
			}
		})
	}
}

func Test_invitationUsecase_GetByID(t *testing.T) {
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
			CreatedAt: "2023-09-28T16:57:27Z",
		},
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name           string
		args           args
		wantInvitation domain.Invitation
		wantErr        bool
		funcGet        helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantInvitation: invitations[0],
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{invitations, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvitation := new(mocks.IInvitation)
			usecase := usecase.NewInvitationUsecase(usecase.InvitationUsecaseArgs{
				InvitationRepository: mockInvitation,
			})
			if tt.funcGet.Called {
				mockInvitation.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotInvitation, err := usecase.GetByID(context.Background(), tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantInvitation, gotInvitation)
			}
		})
	}
}

func Test_invitationUsecase_Accept(t *testing.T) {
	type args struct {
		args domain.InvitationArgs
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		funcUpdateStatus helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{domain.InvitationArgs{
				ID: int64(1),
			}},
			funcUpdateStatus: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvitation := new(mocks.IInvitation)
			usecase := usecase.NewInvitationUsecase(usecase.InvitationUsecaseArgs{
				InvitationRepository: mockInvitation,
			})
			if tt.funcUpdateStatus.Called {
				mockInvitation.On("UpdateStatus", tt.funcUpdateStatus.Input...).Return(tt.funcUpdateStatus.Output...)
			}
			err := usecase.Accept(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_invitationUsecase_Reject(t *testing.T) {
	type args struct {
		args domain.InvitationArgs
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		funcUpdateStatus helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{domain.InvitationArgs{
				ID: int64(1),
			}},
			funcUpdateStatus: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvitation := new(mocks.IInvitation)
			usecase := usecase.NewInvitationUsecase(usecase.InvitationUsecaseArgs{
				InvitationRepository: mockInvitation,
			})
			if tt.funcUpdateStatus.Called {
				mockInvitation.On("UpdateStatus", tt.funcUpdateStatus.Input...).Return(tt.funcUpdateStatus.Output...)
			}
			err := usecase.Reject(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_invitationUsecase_Cancel(t *testing.T) {
	type args struct {
		args domain.InvitationArgs
	}
	tests := []struct {
		name             string
		args             args
		wantErr          bool
		funcUpdateStatus helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{domain.InvitationArgs{
				ID: int64(1),
			}},
			funcUpdateStatus: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInvitation := new(mocks.IInvitation)
			usecase := usecase.NewInvitationUsecase(usecase.InvitationUsecaseArgs{
				InvitationRepository: mockInvitation,
			})
			if tt.funcUpdateStatus.Called {
				mockInvitation.On("UpdateStatus", tt.funcUpdateStatus.Input...).Return(tt.funcUpdateStatus.Output...)
			}
			err := usecase.Cancel(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
