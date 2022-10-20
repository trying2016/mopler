// Package api 定义百度网盘SDK的接口
package api

import (
	"net/http"
)

// Sdker 集成所有的接口，这样在实例化实现了Sdker接口的结构体后，便能够直接调用所有的接口方法
type Sdker interface {
	UserInfoer
	Quotaer
	FileLister
	Uploader
	FileSearcher
	FileManager
	FileMetaer
	Downloader
	Authorizationer
	HttpClient() *http.Client
}
