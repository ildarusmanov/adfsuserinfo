package mocks

import (
	"github.com/stretchr/testify/mock"
)

type UserinfoMock struct {
	mock.Mock
}

func CreateUserinfoMock() *UserinfoMock {
	return new(UserinfoMock)
}

func (m *UserinfoMock) GetId() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *UserinfoMock) GetUsername() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *UserinfoMock) GetEmail() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *UserinfoMock) GetFullName() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *UserinfoMock) GetPhotoUrl() string {
	args := m.Called()
	return args.Get(0).(string)
}
