package adapter_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hieronimusbudi/simple-go-api/internal/adapter"
	"github.com/hieronimusbudi/simple-go-api/internal/domain"
	"github.com/hieronimusbudi/simple-go-api/internal/helpers"
	"github.com/hieronimusbudi/simple-go-api/internal/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestController_CreateMember(t *testing.T) {
	type args struct {
		reqPayload io.Reader
	}

	// success
	member := domain.Member{
		ID:        0,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}
	jsonMember, err := json.Marshal(member)
	require.NoError(t, err)

	// validation fail
	member2 := domain.Member{
		ID:        0,
		FirstName: "john",
		LastName:  "doe",
		Email:     "",
	}
	jsonMember2, err := json.Marshal(member2)
	require.NoError(t, err)

	tests := []struct {
		name         string
		args         args
		funcCreate   helpers.TestFuncCall
		expectedCode int
	}{
		{
			name: "success",
			args: args{
				reqPayload: strings.NewReader(string(jsonMember)),
			},
			funcCreate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{member, nil},
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "validation fail",
			args: args{
				reqPayload: strings.NewReader(string(jsonMember2)),
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "create fail",
			args: args{
				reqPayload: strings.NewReader(string(jsonMember)),
			},
			funcCreate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{domain.Member{}, errors.New("create error")},
			},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMemberUsecase := new(mocks.IMemberUsecase)
			if tt.funcCreate.Called {
				mockMemberUsecase.On("Create", tt.funcCreate.Input...).
					Return(tt.funcCreate.Output...)
			}
			ctr := &adapter.Controller{
				MemberUsecase: mockMemberUsecase,
			}
			c, w := helpers.CreateGinContext(http.MethodPost, "/members", tt.args.reqPayload)
			ctr.CreateMember(c)
			require.Equal(t, tt.expectedCode, w.Result().StatusCode)
		})
	}
}

func TestController_GetMembers(t *testing.T) {
	// success
	members := []domain.Member{{
		ID:        0,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}}

	tests := []struct {
		name         string
		funcGet      helpers.TestFuncCall
		expectedCode int
	}{
		{
			name: "success",
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{members, nil},
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "get fail",
			funcGet: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{[]domain.Member{}, errors.New("get error")},
			},
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMemberUsecase := new(mocks.IMemberUsecase)
			if tt.funcGet.Called {
				mockMemberUsecase.On("Get", tt.funcGet.Input...).
					Return(tt.funcGet.Output...)
			}
			ctr := &adapter.Controller{
				MemberUsecase: mockMemberUsecase,
			}
			c, w := helpers.CreateGinContext(http.MethodGet, "/members", nil)
			ctr.GetMembers(c)
			require.Equal(t, tt.expectedCode, w.Result().StatusCode)
		})
	}
}

func TestController_UpdateMember(t *testing.T) {
	type args struct {
		reqPayload io.Reader
		id         string
	}

	// success
	member := domain.Member{
		ID:        1,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}
	jsonMember, err := json.Marshal(member)
	require.NoError(t, err)

	// validation fail
	member2 := domain.Member{
		ID:        0,
		FirstName: "john",
		LastName:  "doe",
		Email:     "",
	}
	jsonMember2, err := json.Marshal(member2)
	require.NoError(t, err)

	tests := []struct {
		name         string
		args         args
		funcGetByID1 helpers.TestFuncCall
		funcUpdate   helpers.TestFuncCall
		funcGetByID2 helpers.TestFuncCall
		expectedCode int
	}{
		{
			name: "success",
			args: args{
				id:         "1",
				reqPayload: strings.NewReader(string(jsonMember)),
			},
			funcGetByID1: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{member, nil},
			},
			funcUpdate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{nil},
			},
			funcGetByID2: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{member, nil},
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "validation fail",
			args: args{
				reqPayload: strings.NewReader(string(jsonMember2)),
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "update fail",
			args: args{
				id:         "1",
				reqPayload: strings.NewReader(string(jsonMember)),
			},
			funcGetByID1: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{member, nil},
			},
			funcUpdate: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{errors.New("update error")},
			},
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMemberUsecase := new(mocks.IMemberUsecase)
			if tt.funcGetByID1.Called {
				mockMemberUsecase.On("GetByID", tt.funcGetByID1.Input...).
					Return(tt.funcGetByID1.Output...)
			}
			if tt.funcUpdate.Called {
				mockMemberUsecase.On("Update", tt.funcUpdate.Input...).
					Return(tt.funcUpdate.Output...)
			}
			if tt.funcGetByID2.Called {
				mockMemberUsecase.On("GetByID", tt.funcGetByID2.Input...).
					Return(tt.funcGetByID2.Output...)
			}
			ctr := &adapter.Controller{
				MemberUsecase: mockMemberUsecase,
			}
			c, w := helpers.CreateGinContext(http.MethodPut, "/members", tt.args.reqPayload)
			c.AddParam("id", tt.args.id)
			ctr.UpdateMember(c)
			require.Equal(t, tt.expectedCode, w.Result().StatusCode)
		})
	}
}

func TestController_DeleteMember(t *testing.T) {
	type args struct {
		id string
	}
	member := domain.Member{
		ID:        1,
		FirstName: "john",
		LastName:  "doe",
		Email:     "john@mail.com",
	}
	tests := []struct {
		name         string
		args         args
		funcGetByID1 helpers.TestFuncCall
		funcDelete   helpers.TestFuncCall
		expectedCode int
	}{
		{
			name: "success",
			args: args{
				id: "1",
			},
			funcGetByID1: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{member, nil},
			},
			funcDelete: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{nil},
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "delete fail",
			args: args{
				id: "1",
			},
			funcGetByID1: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{member, nil},
			},
			funcDelete: helpers.TestFuncCall{
				Called: true,
				Input:  []interface{}{mock.Anything, mock.Anything},
				Output: []interface{}{errors.New("delete error")},
			},
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMemberUsecase := new(mocks.IMemberUsecase)
			if tt.funcGetByID1.Called {
				mockMemberUsecase.On("GetByID", tt.funcGetByID1.Input...).
					Return(tt.funcGetByID1.Output...)
			}
			if tt.funcDelete.Called {
				mockMemberUsecase.On("Delete", tt.funcDelete.Input...).
					Return(tt.funcDelete.Output...)
			}
			ctr := &adapter.Controller{
				MemberUsecase: mockMemberUsecase,
			}
			c, w := helpers.CreateGinContext(http.MethodDelete, "/members", nil)
			c.AddParam("id", tt.args.id)
			ctr.DeleteMember(c)
			require.Equal(t, tt.expectedCode, w.Result().StatusCode)
		})
	}
}
