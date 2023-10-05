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

func Test_memberUsecase_Create(t *testing.T) {
	type args struct {
		member domain.Member
	}
	member := domain.Member{
		ID:        0,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}
	wantNewMember := member
	wantNewMember.ID = 1
	tests := []struct {
		name          string
		args          args
		wantNewMember domain.Member
		wantErr       bool
		funcCreate    helpers.TestFuncCall
		funcGet       helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{
				member: member,
			},
			wantNewMember: wantNewMember,
			funcCreate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{wantNewMember.ID, nil},
			},
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{[]domain.Member{wantNewMember}, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMember := new(mocks.IMember)
			usecase := usecase.NewMemberUsecase(usecase.MemberUsecaseArgs{
				MemberRepository: mockMember,
			})
			if tt.funcCreate.Called {
				mockMember.On("Create", tt.funcCreate.Input...).Return(tt.funcCreate.Output...)
			}
			if tt.funcGet.Called {
				mockMember.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotNewMember, err := usecase.Create(context.Background(), tt.args.member)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantNewMember, gotNewMember)
			}
		})
	}
}

func Test_memberUsecase_Get(t *testing.T) {
	members := []domain.Member{{
		ID:        1,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}}
	type args struct {
		args domain.MemberArgs
	}
	tests := []struct {
		name        string
		args        args
		wantMembers []domain.Member
		wantErr     bool
		funcGet     helpers.TestFuncCall
	}{
		{
			name:        "success",
			wantMembers: members,
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{members, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMember := new(mocks.IMember)
			usecase := usecase.NewMemberUsecase(usecase.MemberUsecaseArgs{
				MemberRepository: mockMember,
			})
			if tt.funcGet.Called {
				mockMember.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotMembers, err := usecase.Get(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantMembers, gotMembers)
			}
		})
	}
}

func Test_memberUsecase_GetByID(t *testing.T) {
	members := []domain.Member{{
		ID:        1,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name       string
		args       args
		wantMember domain.Member
		wantErr    bool
		funcGet    helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			wantMember: members[0],
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{members, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMember := new(mocks.IMember)
			usecase := usecase.NewMemberUsecase(usecase.MemberUsecaseArgs{
				MemberRepository: mockMember,
			})
			if tt.funcGet.Called {
				mockMember.On("Get", tt.funcGet.Input...).Return(tt.funcGet.Output...)
			}
			gotMember, err := usecase.GetByID(tt.args.ctx, tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantMember, gotMember)
			}
		})
	}
}

func Test_memberUsecase_Update(t *testing.T) {
	type args struct {
		member domain.Member
	}
	member := domain.Member{
		ID:        1,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
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
				member: member,
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
			mockMember := new(mocks.IMember)
			usecase := usecase.NewMemberUsecase(usecase.MemberUsecaseArgs{
				MemberRepository: mockMember,
			})
			if tt.funcUpdate.Called {
				mockMember.On("Update", tt.funcUpdate.Input...).Return(tt.funcUpdate.Output...)
			}
			err := usecase.Update(context.Background(), tt.args.member)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_memberUsecase_Delete(t *testing.T) {
	type args struct {
		args domain.MemberArgs
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		funcDelete helpers.TestFuncCall
	}{
		{
			name: "success",
			args: args{domain.MemberArgs{
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
		t.Run(tt.name, func(t *testing.T) {
			mockMember := new(mocks.IMember)
			usecase := usecase.NewMemberUsecase(usecase.MemberUsecaseArgs{
				MemberRepository: mockMember,
			})
			if tt.funcDelete.Called {
				mockMember.On("Delete", tt.funcDelete.Input...).Return(tt.funcDelete.Output...)
			}
			err := usecase.Delete(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
