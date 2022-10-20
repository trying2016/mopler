# 百度网盘SDK
[![Go Reference](https://pkg.go.dev/badge/github.com/766800551/mopler.svg)](https://pkg.go.dev/github.com/766800551/mopler)<br/>
实现百度网盘的一些接口，整合后方便快速调用，暂时只支持AES加密。

> 使用sdk前必须获取授权，内置了简化授权、设备码授权实现。
>
> 您也可以自行实现，参考：https://pan.baidu.com/union/doc/ol0rsap9s
>
> 官方接口文档地址：https://pan.baidu.com/union/doc/



## 安装

推荐使用go module进行依赖管理

```go
go get -u github.com/766800551/mopler
```



## 快速开始

每个接口的调用示例可以看test文件夹下的测试方法。

```go
package main

import (
	"fmt"
	"github.com/766800551/mopler"
	"github.com/766800551/mopler/impl/authorization"
)

func main() {
	q := authorization.NewQuickTokenImpl(mopler.Token)
	sdk := mopler.New(q)
	info, err := sdk.UserInfo()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", info)
}
```

填入正确的信息，此时就可以看到返回了结果：

```go
&{BaiduName:风屿7010 NetdiskName: AvatarUrl:https://dss0.bdstatic.com/7Ls0a8Sm1A5BphGlnYG/sys/portrait/item/netdisk.1.1986fe57.bLBgo-alQv_C54a0d5WkDA.jpg VipType:2 Uk101935704420 Errno:0 RequestId:9089383694159695404 Errmsg:succ}
```



## 授权

由于百度不允许在非百度官方网站完成账号的授权，因此需要引导用户完成授权。

[示例](test/authentication_test.go)



### 快速授权

快速授权是直接将获取到的token进行传入，目前内置了设备码授权和简化授权两种方式可以获取token



### 设备码授权

设备码授权分两种，第一种为二维码，第二种为设备码，推荐二维码（不用登录账号）

* 设备码：您需要将设备码告知给用户，并且指导用户打开VerificationUrl链接输入设备码和账号密码完成授权
* 二维码：通过QrcodeUrl打开为一张二维码图片，您需要指导用户通过百度类产品扫描完成授权

参数说明：
* ClientId：您应用的AppKey。
* ClientSecret：您应用的SecretKey
* AccessToken：您通过鉴权的方式获取到的access_token， 如果设置了，那么就会跳过获取token的步骤，并且需要您自己实现过期后刷新token的方法。设置此参数后，TokenMode和AuthMethod可以在配置初始化的时候不进行传递
* TokenMode：token验证方式
* AuthMethod： 根据授权模式来引导用户授权。是一个回调函数，参数传递一个response.Code，用来指导用户完成授权，response.Code的参数说明如下：

  - DeviceCode       //设备码，可用于生成单次凭证 Access Token。
  - UserCode         //用户码。 如果选择让用户输入 user code 方式，来引导用户授权，设备需要展示 user code 给用户。
  - VerificationUrl   //用户输入 user code 进行授权的 url。
  - QrcodeUrl        //二维码url，用户用手机等智能终端扫描该二维码完成授权。
  - ExpiresIn           //deviceCode 的过期时间，单位：秒。 到期后 deviceCode 不能换 Access Token。
  - Interval            //deviceCode 换 Access Token 轮询间隔时间，单位：秒。 轮询次数限制小于 expire_in/interval。



### 简化授权

简化授权由于官方并未提交回调，因此需要用户将授权完成后重定向的链接手动拷贝输入，提供了解析链接的函数。





## 接口

| 功能                           | 方法名              | 所属              |
| ------------------------------ | ------------------- | ----------------- |
| 申请token                      | AccessToken         | mopler.SdkContext |
| 复制文件/文件夹                | Copy                | mopler.SdkContext |
| 创建一个空文件夹               | CreateDir           | mopler.SdkContext |
| 删除文件/文件夹                | Delete              | mopler.SdkContext |
| 下载文件/文件夹                | Download            | mopler.SdkContext |
| 获取当前文件夹下的文件列表     | FileList            | mopler.SdkContext |
| 递归获取文件夹下所有的文件列表 | FileListRecursion   | mopler.SdkContext |
| 获取文件信息                   | FileMeta            | mopler.SdkContext |
| 文件搜索                       | FileSearch          | mopler.SdkContext |
| 获取token                      | GetToken            | mopler.SdkContext |
| 获取一个httpclient             | HttpClient          | mopler.SdkContext |
| 移动文件/文件夹                | Move                | mopler.SdkContext |
| 获取网盘容量信息               | Quota               | mopler.SdkContext |
| 刷新token                      | RefreshToken        | mopler.SdkContext |
| 重命名文件/文件夹              | Rename              | mopler.SdkContext |
| 上传文件/文件夹                | Upload              | mopler.SdkContext |
| 获取用户信息                   | UserInfo            | mopler.SdkContext |
| aes加密单文件                  | AesEncrypt          | helper            |
| aes解密单文件                  | AesDecrypt          | helper            |
| aes递归解密文件                | AesDecryptRecursion | helper            |

