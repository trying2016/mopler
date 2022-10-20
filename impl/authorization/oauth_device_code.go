package authorization

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
	"time"
)

// OauthDeviceCode 设备码授权
type OauthDeviceCode struct {
	AToken    string `json:"access_token"`  //获取到的Access QuickToken，Access Token是调用网盘开放API访问用户授权资源的凭证。
	ExpiresIn int    `json:"expires_in"`    //Access Token的有效期，单位为秒。
	RToken    string `json:"refresh_token"` //用于刷新 Access QuickToken, 有效期为10年。
	Scope     string `json:"scope"`         //Access QuickToken 最终的访问权限，即用户的实际授权列表。
}

func NewOauthDeviceCodeImpl() *OauthDeviceCode {
	return &OauthDeviceCode{}
}

func (o *OauthDeviceCode) AccessToken(param any) error {
	switch param.(type) {
	case *request.OauthDeviceCodeReq:
		c, err := o.code(param.(*request.OauthDeviceCodeReq))
		if err != nil {
			return errors.Wrap(err, "failed to get OAuth device code")
		}

		param.(*request.OauthDeviceCodeReq).AuthMethod(c)

		err = o.pollingToken(c.DeviceCode, param.(*request.OauthDeviceCodeReq))
		if err != nil {
			return errors.Wrap(err, "failed to pollingToken")
		}
	default:
		return errors.New("invalid device code request")
	}
	return nil
}

func (o *OauthDeviceCode) RefreshToken(param any) error {
	return nil
}

func (o *OauthDeviceCode) GetToken() string {
	return o.AToken
}

func (o *OauthDeviceCode) code(param *request.OauthDeviceCodeReq) (*response.OauthDeviceCodeResp, error) {
	resp, err := http.Get(fmt.Sprintf(
		"https://openapi.baidu.com/oauth/2.0/device/code?"+
			"response_type=device_code&client_id=%v&scope=basic,netdisk",
		param.ClientId))
	if err != nil {
		return nil, errors.Wrap(err, "获取设备码接口访问失败")
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "获取设备码接口响应读取失败")
	}
	var rc response.OauthDeviceCodeResp
	err = json.Unmarshal(b, &rc)
	if err != nil || rc == (response.OauthDeviceCodeResp{}) {
		return nil, errors.Wrap(err, "反序列化失败")
	}
	return &rc, nil
}

func (o *OauthDeviceCode) pollingToken(code string, param *request.OauthDeviceCodeReq) error {
	for {
		resp, err := http.Get(fmt.Sprintf("https://openapi.baidu.com/oauth/2.0/token?"+
			"grant_type=device_token&code=%v&client_id=%v&client_secret=%v",
			code,
			param.ClientId,
			param.ClientSecret,
		))
		if err != nil {
			return errors.Wrap(err, "获取token接口访问失败")
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "获取token接口响应读取失败")
		}
		//如果获取到了access_token，那么就退出
		if strings.Contains(string(b), "access_token") {
			err = json.Unmarshal(b, &o)
			if err != nil {
				return errors.Wrap(err, "反序列化失败")
			}
			resp.Body.Close()
			return nil
		}
		resp.Body.Close()
		//轮询接口必须5秒以上，这里6秒一次
		time.Sleep(time.Second * 6)
	}
}
