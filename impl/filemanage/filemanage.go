package filemanage

import (
	"encoding/json"
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"github.com/pkg/errors"
	"net/url"
	"os"
	"path/filepath"
)

type FileManage struct {
	api.Sdker
}

func NewFileManageImpl(sdk api.Sdker) *FileManage {
	return &FileManage{
		Sdker: sdk,
	}
}

func (f *FileManage) CreateDir(dirname string, remotePath string) error {
	temp, _ := os.MkdirTemp(".", "")
	defer os.RemoveAll(temp)
	err := os.Mkdir(temp+string(filepath.Separator)+dirname, os.ModePerm)
	if err != nil {
		return errors.WithStack(err)
	}
	absPath, err := filepath.Abs(temp + string(filepath.Separator) + dirname)
	if err != nil {
		return errors.WithMessage(err, "获取绝对路径失败")
	}
	err = f.Upload(&request.UploadReq{SrcPath: absPath, RemotePath: remotePath})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (f *FileManage) manage(req *request.FileManagerReq, opera string) (resp *response.FileManagerResp, err error) {
	vs := url.Values{}
	vs.Add("async", req.Async)
	vs.Add("ondup", req.Ondup)
	if len(req.Filelist) > 0 {
		b, err := json.Marshal(req.Filelist)
		if err != nil {
			return nil, errors.WithMessage(err, "Failed to marshal")
		}
		vs.Add("filelist", string(b))
	}
	err = helper.HttpPostForm(f.HttpClient(),
		fmt.Sprintf("http://pan.baidu.com/rest/2.0/xpan/file?method=filemanager&access_token=%v&opera=%v",
			f.GetToken(), opera),
		vs, &resp)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return
}

func (f *FileManage) Copy(req *request.FileManagerReq) (resp *response.FileManagerResp, err error) {
	return f.manage(req, "copy")
}

func (f *FileManage) Rename(req *request.FileManagerReq) (resp *response.FileManagerResp, err error) {
	return f.manage(req, "rename")
}

func (f *FileManage) Move(req *request.FileManagerReq) (resp *response.FileManagerResp, err error) {
	return f.manage(req, "move")
}

func (f *FileManage) Delete(req *request.FileManagerReq) (resp *response.FileManagerResp, err error) {
	return f.manage(req, "delete")
}
