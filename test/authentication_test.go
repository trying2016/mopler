package test

import (
	"fmt"
	"github.com/766800551/mopler"
	"github.com/766800551/mopler/impl/authorization"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"github.com/pkg/browser"
	"log"
	"testing"
)

//授权相关的测试,请在TestMain中直接赋予了AccessToken，方便后续接口测试

func TestQuickToken(t *testing.T) {
	if sdk.GetToken() == "" {
		t.Fatal("token is empty")
	}
	fmt.Println(sdk.GetToken())
}

//以下的这些方法由于涉及到浏览器交互，请单独进行测试
//请不要在TestMain中直接赋予了AccessToken!!!!

// 测试简化模式授权，此调用会打开浏览器，认证后会重定向到新的链接
func TestImplicitGrant(t *testing.T) {
	c := authorization.NewImplicitGrantImpl()
	sdk = mopler.New(c)
	err := sdk.AccessToken(&request.ImplicitGrantReq{
		ClientId: mopler.ClientId,
	})
	if err != nil {
		log.Printf("%+v", err)
		t.Fatal(err)
	}
}

// 上面的测试的认证链接会返回token的数据，这里测试解析是否正常
func TestImplicitGrantToken(t *testing.T) {
	c := authorization.NewImplicitGrantImpl()
	sdk = mopler.New(c)
	err := sdk.AccessToken(&request.ImplicitGrantReq{
		ClientId: mopler.ClientId,
	})
	if err != nil {
		t.Fatal(err)
	}
	//将重定向的网页传入这个方法里
	c.SetToken("https://openapi.baidu.com/oauth/2.0/login_success#expires_in=2592000&access_token=xxxxxx&session_secret=&session_key=&scope=basic+netdisk")
	if sdk.GetToken() == "" {
		t.Fatal("token is empty")
	}
	fmt.Println(sdk.GetToken())
}

// 设备码获取授权（设备码授权会将token进行回写）
func TestOauthDeviceCode(t *testing.T) {
	c := authorization.NewOauthDeviceCodeImpl()
	sdk = mopler.New(c)
	err := sdk.AccessToken(&request.OauthDeviceCodeReq{
		ClientId:     mopler.ClientId,
		ClientSecret: mopler.ClientSecret,
		AuthMethod: func(resp *response.OauthDeviceCodeResp) {
			err := browser.OpenURL(resp.QrcodeUrl)
			if err != nil {
				t.Fatal(err)
			}
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sdk.GetToken())
}
