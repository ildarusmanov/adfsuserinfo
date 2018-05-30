package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/adfsuserinfo/interfaces"
)

// UserinfoController provides access to userinfo by given token
type UserinfoController struct {
	userinfoService interfaces.UserinfoService
}

// CreateUserinfoController creates controller
func CreateUserinfoController(us interfaces.UserinfoService) *UserinfoController {
	return &UserinfoController{userinfoService: us}
}

// Get userinfo by access_token param
// Request example:
// GET /api/v1/me?access_token=[[jwt access token]]
func (uc *UserinfoController) Get(ctx *gin.Context) {
	accessToken, ok := ctx.GetQuery("access_token")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
	}

	userinfo, err := uc.userinfoService.GetUserinfoByToken(accessToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not found"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        userinfo.GetId(),
		"email":     userinfo.GetEmail(),
		"full_name": userinfo.GetFullName(),
		"photo_url": userinfo.GetPhotoUrl(),
		"username":  userinfo.GetUsername(),
	})
}
