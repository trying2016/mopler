// Package filelist 实现获取网盘文件列表
package filelist

import (
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"net/url"
	"strings"
)

type FileList struct {
	api.Sdker
}

func NewFileListImpl(sdk api.Sdker) *FileList {
	return &FileList{
		Sdker: sdk,
	}
}

func (f *FileList) FileList(req *request.FileListReq) (*response.FileListResp, error) {
	var reval response.FileListResp
	err := helper.HttpGet(
		f.HttpClient(),
		fmt.Sprintf("https://pan.baidu.com/rest/2.0/xpan/file?method=list&access_token=%v"+
			"&folder=%v&order=%v&showempty=%v&desc=%v&dir=%v&limit=%v&start=%v&web=%v",
			f.GetToken(), req.Folder, req.Order, req.Showempty, req.Desc,
			url.QueryEscape(req.Dir), req.Limit, req.Start, req.Web),
		&reval,
	)
	//错误码-9为文件不存在。这里忽略！
	if err != nil && strings.Contains(err.Error(), "错误码：-9") {
		return &reval, nil
	}
	return &reval, err
}

func (f *FileList) FileListRecursion(req *request.FileListRecursionReq) (*response.FileListResp, error) {
	var reval response.FileListResp
	err := helper.HttpGet(
		f.HttpClient(),
		fmt.Sprintf("https://pan.baidu.com/rest/2.0/xpan/file?method=listall&access_token=%v&recursion=1"+
			"&order=%v&desc=%v&path=%v&limit=%v&start=%v&web=%v",
			f.GetToken(), req.Order, req.Desc,
			url.QueryEscape(req.Path), req.Limit, req.Start, req.Web),
		&reval,
	)
	return &reval, err
}
