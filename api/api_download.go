package api

import "github.com/766800551/mopler/model/response"

// Downloader 文件下载
type Downloader interface {
	Download(downDir string, path string, downloadStatuFunc func(fms response.FileMetaResp)) (err error)
}
