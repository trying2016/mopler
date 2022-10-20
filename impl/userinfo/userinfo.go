// Package userinfo 实现获取网盘用户信息
package userinfo

import (
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/response"
	"github.com/pkg/errors"
)

type UserInfo struct {
	api.Sdker
}

func NewUserInfoImpl(sdker api.Sdker) *UserInfo {
	return &UserInfo{
		Sdker: sdker,
	}
}

// UserInfo 获取用户信息
func (u *UserInfo) UserInfo() (*response.UserInfoResp, error) {
	var reval response.UserInfoResp
	err := helper.HttpGet(u.HttpClient(),
		fmt.Sprintf("https://pan.baidu.com/rest/2.0/xpan/nas?method=uinfo&access_token=%v", u.GetToken()), &reval)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user info")
	}
	return &reval, nil
}
