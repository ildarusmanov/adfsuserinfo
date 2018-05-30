package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ildarusmanov/adfsuserinfo/interfaces"
	"github.com/ildarusmanov/adfsuserinfo/models"
)

type UserinfoService struct{}

func CreateUserinfoService() *UserinfoService {
	return &UserinfoService{}
}

func (us *UserinfoService) GetUserinfoByToken(tokenStr string) (interfaces.Userinfo, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, make(jwt.MapClaims))

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return us.buildUserinfo(claims), nil
	}

	return nil, err
}

func (us *UserinfoService) buildUserinfo(claims jwt.MapClaims) *models.Userinfo {
	return models.CreateUserinfo(
		getMapClaimsStringValue("id", claims),
		getMapClaimsStringValue("username", claims),
		getMapClaimsStringValue("email", claims),
		getMapClaimsStringValue("FullName", claims),
		getMapClaimsStringValue("PhotoURL", claims),
	)
}

func getMapClaimsStringValue(key string, claims jwt.MapClaims) string {
	if v, ok := claims[key]; ok {
		return v.(string)
	}

	return ""
}
