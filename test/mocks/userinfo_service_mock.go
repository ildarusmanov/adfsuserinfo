package mocks

import (
	"github.com/ildarusmanov/adfsuserinfo/interfaces"
	"github.com/stretchr/testify/mock"
)

type UserinfoServiceMock struct {
	mock.Mock
}

func CreateUserinfoServiceMock() *UserinfoServiceMock {
	return new(UserinfoServiceMock)
}

func (m *UserinfoServiceMock) GetUserinfoByToken(token string) (interfaces.Userinfo, error) {
	args := m.Called(token)
	userinfo := args.Get(0)
	err := args.Error(1)

	if err != nil {
		return nil, err
	}

	return userinfo.(interfaces.Userinfo), err
}
