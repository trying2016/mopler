package api

import (
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
)

// FileMetaer 文件详情
type FileMetaer interface {
	FileMeta(*request.FileMetaReq) (*response.FileMetasResp, error)
}
