package response

type OauthDeviceCodeResp struct {
	DeviceCode      string `json:"device_code"`      //设备码，可用于生成单次凭证 Access Token。
	UserCode        string `json:"user_code"`        //用户码。 如果选择让用户输入 user code 方式，来引导用户授权，设备需要展示 user code 给用户。
	VerificationUrl string `json:"verification_url"` //用户输入 user code 进行授权的 url。
	QrcodeUrl       string `json:"qrcode_url"`       //二维码url，用户用手机等智能终端扫描该二维码完成授权。
	ExpiresIn       int    `json:"expires_in"`       //deviceCode 的过期时间，单位：秒。 到期后 deviceCode 不能换 Access Token。
	Interval        int    `json:"interval"`         //deviceCode 换 Access Token 轮询间隔时间，单位：秒。 轮询次数限制小于 expire_in/interval。
}
