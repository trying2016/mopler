package api

import (
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
)

// FileManager 文件管理，这里的文件是指网盘里的文件
type FileManager interface {
	CreateDir(dirname string, remotePath string) error
	Copy(req *request.FileManagerReq) (resp *response.FileManagerResp, err error)
	Rename(req *request.FileManagerReq) (resp *response.FileManagerResp, err error)
	Move(req *request.FileManagerReq) (resp *response.FileManagerResp, err error)
	Delete(req *request.FileManagerReq) (resp *response.FileManagerResp, err error)
}
