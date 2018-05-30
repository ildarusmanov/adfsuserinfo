package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ildarusmanov/adfsuserinfo/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestRouterEndpoints(t *testing.T) {
	var (
		validTokenStr = "123123"
	)

	validUserinfo := mocks.CreateUserinfoMock()
	validUserinfo.On("GetId").Return("1")
	validUserinfo.On("GetEmail").Return("test@email.com")
	validUserinfo.On("GetUsername").Return("test")
	validUserinfo.On("GetFullName").Return("Test Name")
	validUserinfo.On("GetPhotoUrl").Return("http://test.com/test.jpg")

	userinfoService := mocks.CreateUserinfoServiceMock()
	userinfoService.On("GetUserinfoByToken", validTokenStr).Return(validUserinfo, nil)

	serviceLocator := mocks.CreateServiceLocatorMock()
	serviceLocator.On("Get", "Userinfo").Return(userinfoService, nil)

	router := CreateRouter(serviceLocator)

	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/api/v1/me?access_token="+validTokenStr, nil)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/v1/status/check", nil)

	router.ServeHTTP(w1, req1)
	router.ServeHTTP(w2, req2)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, w1.Code)
	assert.Equal(http.StatusOK, w2.Code)
}
