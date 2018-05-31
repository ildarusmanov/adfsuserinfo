package interfaces

type UserinfoService interface {
	GetUserinfoByToken(token string) (Userinfo, error)
}
