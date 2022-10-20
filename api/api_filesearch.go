package api

import (
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
)

// FileSearcher 文件搜索
type FileSearcher interface {
	FileSearch(req *request.FileSearchReq) (*response.FileSearchResp, error)
}
