package response

type UserInfoResp struct {
	BaiduName   string `json:"baidu_name"`   //百度帐号
	NetdiskName string `json:"netdisk_name"` //网盘帐号
	AvatarUrl   string `json:"avatar_url"`   //头像地址
	VipType     int    `json:"vip_type"`     //会员类型，0普通用户、1普通会员、2超级会员
	Uk          int    `json:"uk"`           //用户ID
	Errno       int    `json:"errno"`
	RequestId   string `json:"request_id"`
	Errmsg      string `json:"errmsg"`
}
