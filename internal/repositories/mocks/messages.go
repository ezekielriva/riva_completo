// Code generated by MockGen. DO NOT EDIT.
// Source: messages.go
//
// Generated by this command:
//
//	mockgen -source=messages.go
//

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	entities "github.com/ezekielriva/riva_completo/internal/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockMessagesRepository is a mock of MessagesRepository interface.
type MockMessagesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessagesRepositoryMockRecorder
}

// MockMessagesRepositoryMockRecorder is the mock recorder for MockMessagesRepository.
type MockMessagesRepositoryMockRecorder struct {
	mock *MockMessagesRepository
}

// NewMockMessagesRepository creates a new mock instance.
func NewMockMessagesRepository(ctrl *gomock.Controller) *MockMessagesRepository {
	mock := &MockMessagesRepository{ctrl: ctrl}
	mock.recorder = &MockMessagesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessagesRepository) EXPECT() *MockMessagesRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockMessagesRepository) Create(m *entities.Message) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMessagesRepositoryMockRecorder) Create(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessagesRepository)(nil).Create), m)
}
