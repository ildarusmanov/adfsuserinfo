package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	jwtExample    = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IjA5MlZ3UmtaaDVGb2d5ZUItTWxhRktTaTlvcyJ9.eyJhdWQiOiJ1cm46bWljcm9zb2Z0OnVzZXJpbmZvIiwiaXNzIjoiaHR0cDovL3Nzby51bml2ZXJzaXR5Lmlubm9wb2xpcy5ydS9hZGZzL3NlcnZpY2VzL3RydXN0IiwiaWF0IjoxNTIzOTYxOTkwLCJleHAiOjE1MjM5NjU1OTAsImFwcHR5cGUiOiJDb25maWRlbnRpYWwiLCJhcHBpZCI6Ik1vb2RsZS1jZjA1NzQyNC0yNTNkLTQ1NjktYmNlMy02MWZjZjNmYWE3NjMiLCJhdXRobWV0aG9kIjoidXJuOm9hc2lzOm5hbWVzOnRjOlNBTUw6Mi4wOmFjOmNsYXNzZXM6UGFzc3dvcmRQcm90ZWN0ZWRUcmFuc3BvcnQiLCJhdXRoX3RpbWUiOiIyMDE4LTA0LTE3VDEwOjQ2OjE3LjA2OFoiLCJ2ZXIiOiIxLjAiLCJzY3AiOiJvcGVuaWQiLCJzdWIiOiJVbXdsU1NmSDFSS2RQbkdYUHJ1UlQrT3RHcHZKQVlUZ2FhQ0dMR0kxbkVrPSJ9.hsRhRbQMtfP-bix7romHG7duBKzLJPVCyZ2F-h8PymE4vO5PfT0SEa4gt-6CHVPYqrsfknQb3z7uuHropf6yCjUGTNaM3pbmA2FLfAXmIyy_Ya3aLoHrqdFWOFB72Ophkz5QnZDhK3sjtPfXo5dSlVojlmgOMjEWm00YUPpsmdCzUCyu3NrdWzMlw-qb-npQ8c2Kh1gi1FvnyprIvFDLE5ETBps0I8DwNGyPa4WQ8H_7ZApY95xqqHXmp24UujCu5uQE_sApxX7L4z_QZQVp0ypFuQzES4Mat_i3vp23X64IC-Gw896Bs4NSSphW5520AW6aRYUA1AXYYlfPfOdcrQ"
	userinfoHost  = ""
	userinfoToken = ""
)

func TestCreateUserinfoServie(t *testing.T) {
	service := CreateUserinfoService()
	assert := assert.New(t)
	assert.NotNil(service)
}

func TestGetUserinfoByToken(t *testing.T) {
	validToken := jwtExample
	invalidToken := "123"

	service := CreateUserinfoService()

	userinfoValidToken, errValidToken := service.GetUserinfoByToken(validToken)
	userinfoInvalidToken, errInvalidToken := service.GetUserinfoByToken(invalidToken)

	assert := assert.New(t)
	assert.NotNil(userinfoValidToken)
	assert.Nil(errValidToken)
	assert.Nil(userinfoInvalidToken)
	assert.NotNil(errInvalidToken)
}
