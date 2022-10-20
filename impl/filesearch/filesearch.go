// Package filesearch 实现网盘文件搜索
package filesearch

import (
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
)

type FileSearch struct {
	api.Sdker
}

func NewFileSearchImpl(sdk api.Sdker) *FileSearch {
	return &FileSearch{
		Sdker: sdk,
	}
}

func (f *FileSearch) FileSearch(req *request.FileSearchReq) (*response.FileSearchResp, error) {
	var reval response.FileSearchResp
	err := helper.HttpGet(
		f.HttpClient(),
		fmt.Sprintf("http://pan.baidu.com/rest/2.0/xpan/file?access_token=%v&method=search&key=%v"+
			"&web=%v&dir=%v&num=%v&page=%v&recursion=%v",
			f.GetToken(),
			req.Key,
			req.Web,
			req.Dir,
			req.Num,
			req.Page,
			req.Recursion,
		),
		&reval,
	)
	return &reval, err
}
