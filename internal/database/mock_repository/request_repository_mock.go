// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/database/repository/request_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	model "github.com/AntoniKania/nbp-api-wrapper-go/internal/api/model"
	gomock "go.uber.org/mock/gomock"
)

// MockRequestRepository is a mock of RequestRepository interface.
type MockRequestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRequestRepositoryMockRecorder
}

// MockRequestRepositoryMockRecorder is the mock recorder for MockRequestRepository.
type MockRequestRepositoryMockRecorder struct {
	mock *MockRequestRepository
}

// NewMockRequestRepository creates a new mock instance.
func NewMockRequestRepository(ctrl *gomock.Controller) *MockRequestRepository {
	mock := &MockRequestRepository{ctrl: ctrl}
	mock.recorder = &MockRequestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestRepository) EXPECT() *MockRequestRepositoryMockRecorder {
	return m.recorder
}

// SaveRequest mocks base method.
func (m *MockRequestRepository) SaveRequest(ctx context.Context, request *model.Request) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveRequest", ctx, request)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveRequest indicates an expected call of SaveRequest.
func (mr *MockRequestRepositoryMockRecorder) SaveRequest(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveRequest", reflect.TypeOf((*MockRequestRepository)(nil).SaveRequest), ctx, request)
}
