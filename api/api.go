// Package api 定义所有在使用百度网盘SDK时，可能使用到的方法。
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
