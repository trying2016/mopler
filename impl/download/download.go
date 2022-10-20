// Package download 实现网盘文件下载
package download

import (
	"bufio"
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Download struct {
	api.Sdker
	Err error
}

func NewDownloadImpl(sdk api.Sdker) *Download {
	return &Download{
		Sdker: sdk,
	}
}

func (d *Download) Download(downDir string, path string, downloadStatuFunc func(fms response.FileMetaResp)) (err error) {
	//存储需要下载的文件id
	var fsIds []int64
	//先获取上层目录，根据列表接口，然后将文件id加入需要进行下载的里面
	s, _ := d.FileList(&request.FileListReq{Dir: helper.PathSeparatorFormat(filepath.Dir(path))})
	for _, v := range s.List {
		if v.Path == path {
			fsIds = append(fsIds, v.FsId)
		}
	}

	list, err := d.FileListRecursion(&request.FileListRecursionReq{Path: path})
	if err != nil {
		return errors.WithStack(err)
	}
	for _, f := range list.List {
		fsIds = append(fsIds, f.FsId)
	}
	err = d.download(downDir, path, fsIds, downloadStatuFunc)
	if err != nil {
		return errors.WithMessage(err, "文件下载失败")
	}
	return
}

func (d *Download) download(downDir string, basePath string, fsIds []int64, downloadStatuFunc func(fms response.FileMetaResp)) (err error) {
	basePath = helper.PathSeparatorFormat(filepath.Dir(basePath))
	//列举当前目录下的所有文件
	_fm, err := d.FileMeta(&request.FileMetaReq{Fsids: fsIds})
	if err != nil {
		return errors.WithStack(err)
	}
	fms := _fm.List
	if len(fms) == 0 {
		return errors.WithMessage(err, "没有需要下载的文件，请检查fsId是否正确！")
	}

	wg := &sync.WaitGroup{}
	_ = os.MkdirAll(downDir, os.ModePerm)
	for _, f := range fms {
		wg.Add(1)
		go func(f response.FileMetaResp) {
			defer wg.Done()
			//如果是目录直接创建，不进行下载操作
			if f.Isdir == 1 {
				err = os.MkdirAll(downDir+strings.Replace(f.Path, basePath, "", 1), os.ModePerm)
				if err != nil {
					d.Err = errors.WithStack(err)
				}
				return
			}
			req, err := http.NewRequest("GET",
				fmt.Sprintf("%v&access_token=%v", f.Dlink, d.GetToken()), nil)
			if err != nil {
				d.Err = errors.WithStack(err)
				return
			}
			req.Header.Set("User-Agent", "pan.baidu.com")
			resp, err := d.HttpClient().Do(req)
			if err != nil {
				d.Err = errors.WithStack(err)
				return
			}
			defer resp.Body.Close()
			var file *os.File
			file, err = os.OpenFile(helper.PathSeparatorFormat(downDir+strings.Replace(f.Path, basePath, "", 1)),
				os.O_CREATE|os.O_RDWR,
				os.ModePerm)
			if err != nil {
				d.Err = errors.WithStack(err)
				return
			}
			defer file.Close()
			w := bufio.NewWriterSize(file, 1024*1024*4)
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				d.Err = errors.WithStack(err)
				return
			}
			_ = w.Flush()
			return
		}(f)
	}
	for i := 0; i < len(fms); i++ {
		if d.Err != nil {
			return d.Err
		}
		if downloadStatuFunc != nil {
			downloadStatuFunc(fms[i])
		}
	}
	wg.Wait()
	return nil
}
