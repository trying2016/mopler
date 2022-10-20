// Package authorization 实现了百度授权
package authorization

// QuickToken 快速授权，直接赋予token
type QuickToken struct {
	AToken string
}

func NewQuickTokenImpl(token string) *QuickToken {
	return &QuickToken{
		AToken: token,
	}
}

func (t *QuickToken) AccessToken(param any) error {
	panic("implement me")
}

func (t *QuickToken) RefreshToken(param any) error {
	panic("implement me")
}

func (t *QuickToken) GetToken() string {
	return t.AToken
}
