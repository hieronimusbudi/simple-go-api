package repository_test

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hieronimusbudi/simple-go-api/internal/adapter/repository"
	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/test"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var (
	db *sqlx.DB
)

// This test is integration test
func TestMain(m *testing.M) {
	var err error
	var terminateContainer func() // variable to store function to terminate container
	terminateContainer, db, err = test.SetupMySQLContainer()
	defer terminateContainer() // make sure container will be terminated at the end
	if err != nil {
		log.Println("failed to setup MySQL container")
		panic(err)
	}
	os.Exit(m.Run())
}

func Test_memberAdapterRepository_Create(t *testing.T) {
	member := domain.Member{
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}
	type args struct {
		member domain.Member
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
				member: member,
			},
			wantId: 3, // based on last ID from test data.sql
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewMemberRepository(repository.MemberAdapterRepositoryArgs{
				DB: db,
			})
			gotId, err := repo.Create(context.Background(), tt.args.member)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantId, gotId)
			}
		})
	}
}

func Test_memberAdapterRepository_Get(t *testing.T) {
	members := []domain.Member{
		{
			ID:        1,
			FirstName: "linus",
			LastName:  "torvalds",
			Email:     "linus@mail.com",
			CreatedAt: "2023-10-02T11:05:01Z",
		},
		{
			ID:        2,
			FirstName: "ron",
			LastName:  "west",
			Email:     "ron@mail.com",
			CreatedAt: "2023-10-02T11:05:43Z",
		},
	}
	type args struct {
		args domain.MemberArgs
	}
	tests := []struct {
		name        string
		args        args
		wantMembers []domain.Member
		wantErr     bool
	}{
		{
			name:        "success",
			wantMembers: members,
			args: args{
				domain.MemberArgs{
					IDs: []int64{1, 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewMemberRepository(repository.MemberAdapterRepositoryArgs{
				DB: db,
			})
			gotMembers, err := repo.Get(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.wantMembers, gotMembers)
			}
		})
	}
}

func Test_memberAdapterRepository_Update(t *testing.T) {
	member := domain.Member{
		ID:        1,
		FirstName: "linus updated",
		LastName:  "torvalds updated",
		Email:     "updatedlinus@mail.com",
		CreatedAt: "2023-10-02T11:05:01Z",
	}
	type args struct {
		member domain.Member
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				member: member,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewMemberRepository(repository.MemberAdapterRepositoryArgs{
				DB: db,
			})
			err := repo.Update(context.Background(), tt.args.member)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// check data
				members, err := repo.Get(context.Background(), domain.MemberArgs{
					IDs: []int64{tt.args.member.ID},
				})
				require.NoError(t, err)
				gotMember := members[0]
				require.Equal(t, tt.args.member, gotMember)
			}
		})
	}
}

func Test_memberAdapterRepository_Delete(t *testing.T) {
	type args struct {
		args domain.MemberArgs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				args: domain.MemberArgs{
					ID: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repository.NewMemberRepository(repository.MemberAdapterRepositoryArgs{
				DB: db,
			})
			err := repo.Delete(context.Background(), tt.args.args)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				// check data
				members, err := repo.Get(context.Background(), domain.MemberArgs{
					IDs: []int64{tt.args.args.ID},
				})
				require.NoError(t, err)
				require.Equal(t, 0, len(members))
			}
		})
	}
}
