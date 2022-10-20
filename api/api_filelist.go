package api

import (
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
)

// FileLister 文件列表
type FileLister interface {
	FileList(*request.FileListReq) (*response.FileListResp, error)
	FileListRecursion(*request.FileListRecursionReq) (*response.FileListResp, error)
}
