// Package request 所有网盘接口入参的定义
package request

import "github.com/766800551/mopler/model/response"

//授权相关

// ImplicitGrantReq 简化模式获取授权码
type ImplicitGrantReq struct {
	ClientId string //  M    您应用的AppKey。
	Display  string //  O    授权页面展示方式。参见授权展示方式。
	/*
		page	全屏形式的授权页面(默认)，适用于 web 应用。
		popup	弹框形式的授权页面，适用于桌面软件应用和 web 应用
		dialog	浮层形式的授权页面，只能用于站内 web 应用
		mobile	Iphone/Android 等智能移动终端上用的授权页面，适用于 Iphone/Android 等智能移动终端上的应用
		tv	电视等超大显示屏使用的授权页面
		pad	适配 IPad/Android 等智能平板电脑使用的授权页面
	*/
}

type OauthDeviceCodeReq struct {
	ClientId     string //您应用的AppKey。
	ClientSecret string //您应用的SecretKey
	//根据授权模式来引导用户授权
	AuthMethod func(resp *response.OauthDeviceCodeResp)
}
