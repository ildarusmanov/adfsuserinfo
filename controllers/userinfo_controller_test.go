package controllers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/adfsuserinfo/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserinfoController(t *testing.T) {
	userinfo := mocks.CreateUserinfoServiceMock()

	controller := CreateUserinfoController(userinfo)

	assert.NotNil(t, controller)
}

func TestUserinfoMethod(t *testing.T) {
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

	controller := CreateUserinfoController(userinfoService)
	w1 := httptest.NewRecorder()
	ctx1, _ := gin.CreateTestContext(w1)
	ctx1.Request, _ = http.NewRequest("GET", "/?access_token="+validTokenStr, nil)
	controller.Get(ctx1)
	resp1 := w1.Result()

	w2 := httptest.NewRecorder()
	ctx2, _ := gin.CreateTestContext(w1)

	form := url.Values{}
	form.Add("access_token", validTokenStr)

	ctx2.Request, _ = http.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	ctx2.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	controller.Get(ctx1)
	resp2 := w2.Result()

	assert := assert.New(t)
	assert.Equal(resp1.StatusCode, http.StatusOK)
	assert.Equal(resp2.StatusCode, http.StatusOK)
}
