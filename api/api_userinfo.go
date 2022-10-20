package api

import "github.com/766800551/mopler/model/response"

// UserInfoer 用户信息
type UserInfoer interface {
	UserInfo() (*response.UserInfoResp, error)
}
