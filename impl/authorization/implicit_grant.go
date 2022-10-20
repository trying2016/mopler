package authorization

import (
	"fmt"
	"github.com/766800551/mopler/model/request"
	"github.com/pkg/browser"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
	"strings"
)

// ImplicitGrant 简化模式授权
type ImplicitGrant struct {
	AToken    string `json:"access_token"`  //获取到的Access QuickToken，Access Token是调用网盘开放API访问用户授权资源的凭证。
	ExpiresIn int    `json:"expires_in"`    //Access Token的有效期，单位为秒。
	RToken    string `json:"refresh_token"` //用于刷新 Access QuickToken, 有效期为10年。
	Scope     string `json:"scope"`         //Access QuickToken 最终的访问权限，即用户的实际授权列表。
}

func NewImplicitGrantImpl() *ImplicitGrant {
	return &ImplicitGrant{}
}

// AccessToken 向百度网盘请求Token
func (g *ImplicitGrant) AccessToken(param any) error {
	switch param.(type) {
	case *request.ImplicitGrantReq:
		l := fmt.Sprintf("http://openapi.baidu.com/oauth/2.0/authorize?response_type=token&client_id=%v&redirect_uri=oob&scope=basic,netdisk&display=%v",
			param.(*request.ImplicitGrantReq).ClientId,
			param.(*request.ImplicitGrantReq).Display)
		_ = browser.OpenURL(l)
	default:
		return errors.New("invalid implicit grant request")
	}
	return nil
}

// RefreshToken 刷新Token
func (g *ImplicitGrant) RefreshToken(param any) error {
	return g.AccessToken(param)
}

// GetToken 获取Token
func (g *ImplicitGrant) GetToken() string {
	return g.AToken
}

// SetToken 根据鉴权的url获取token
func (g *ImplicitGrant) SetToken(l string) {
	r, _ := url.ParseRequestURI(strings.Replace(l, "#", "?", 1))
	g.AToken = r.Query().Get("access_token")
	g.ExpiresIn, _ = strconv.Atoi(r.Query().Get("expires_in"))
	g.RToken = r.Query().Get("refresh_token")
	g.Scope = strings.Replace(r.Query().Get("scope"), " ", ",", 1)
}
