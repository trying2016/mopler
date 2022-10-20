// Package filemeta 实现获取网盘文件详情
package filemeta

import (
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"strings"
)

type FileMeta struct {
	api.Sdker
}

func NewFileMetaImpl(sdk api.Sdker) *FileMeta {
	return &FileMeta{
		Sdker: sdk,
	}
}

func (f *FileMeta) FileMeta(req *request.FileMetaReq) (*response.FileMetasResp, error) {
	var reval response.FileMetasResp
	err := helper.HttpGet(
		f.HttpClient(),
		fmt.Sprintf("http://pan.baidu.com/rest/2.0/xpan/multimedia?access_token=%v&method=filemetas&dlink=1&fsids=%v",
			f.GetToken(),
			fmt.Sprintf("[%v]", strings.Join(helper.Int64SliceToStringSlice(req.Fsids), ","))),
		&reval,
	)
	return &reval, err
}
