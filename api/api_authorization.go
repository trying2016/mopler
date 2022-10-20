package api

// Authorizationer 获取百度网盘的授权需要实现此接口
type Authorizationer interface {
	AccessToken(any) error
	RefreshToken(any) error
	GetToken() string
}
