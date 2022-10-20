// Package upload 实现网盘文件上传
package upload

import (
	"crypto/md5"
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"github.com/pkg/errors"
	"io"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Upload struct {
	api.Sdker
	finish    atomic.Int32
	SliceSize int64                           //分片大小，只能是4MB，16MB，32MB
	BlockList []string                        //文件各分片MD5数组的json串。
	UR        *request.UploadReq              //上传请求参数
	UPR       *response.UploadPrecreateResp   //预上传返回参数
	USFR      []*response.UploadSuperFileResp //切片上传返回参数
	UCFR      *response.UploadCreateFileResp  //创建网盘文件返回参数
	SFIS      []*request.SliceFileInfo        //分割文件后的信息
	Err       error
}

func NewUploadImpl(sdk api.Sdker) *Upload {
	return &Upload{
		Sdker:     sdk,
		finish:    atomic.Int32{},
		SliceSize: 1024 * 1024 * 4,
	}
}

// BUG(upload) 目前协程并未生效，因为整个方法都被锁住了，待后续修改
func (u *Upload) upload(req *request.UploadReq) error {
	{
		u.UR = req
		u.finish = atomic.Int32{}
		u.BlockList = make([]string, 0)
		u.UPR = nil
		u.USFR = nil
		u.UCFR = nil
		u.SFIS = nil
		u.Err = nil

	}

	info, _ := os.Stat(u.UR.SrcPath)

	if info.IsDir() {
		u.UR.Isdir = "1"
	} else {
		u.UR.Isdir = "0"
		//判断是否需要进行加密，如果需要就加密
		c1 := u.encrypt()
		defer c1()
		//分割文件
		c2 := u.splitFile(req.SrcPath)
		defer c2()
		//预上传
		u.precreate()
		//分片上传
		u.superFile()
	}

	//合并分片文件
	u.create()
	return u.Err
}

// precreate 预上传
func (u *Upload) precreate() {
	vs := url.Values{}
	info, err := os.Stat(u.UR.SrcPath)
	if err != nil {
		u.Err = errors.Wrap(err, "failed to stat SrcPath")
	}
	u.UR.Size = info.Size()
	ct, _, mt := helper.FileTime(info)
	u.UR.LocalCtime = ct
	u.UR.LocalMtime = mt
	u.UR.Autoinit = "1"
	vs.Add("size", fmt.Sprintf("%v", u.UR.Size))
	vs.Add("path", u.UR.RemotePath)
	vs.Add("block_list", strings.Replace(fmt.Sprintf("%q", u.BlockList), " ", ",", -1))
	vs.Add("autoinit", u.UR.Autoinit)
	vs.Add("rtype", u.UR.Rtype)
	vs.Add("uploadid", u.UR.Uploadid)
	vs.Add("content-md5", u.UR.ContentMd5)
	vs.Add("slice-md5", u.UR.SliceMd5)
	vs.Add("local_ctime", u.UR.LocalCtime)
	vs.Add("local_mtime", u.UR.LocalMtime)
	vs.Add("isdir", u.UR.Isdir)
	var resp response.UploadPrecreateResp
	err = helper.HttpPostForm(u.HttpClient(), fmt.Sprintf(
		"http://pan.baidu.com/rest/2.0/xpan/file?method=precreate&access_token=%v", u.GetToken()),
		vs, &resp,
	)
	if err != nil {
		u.Err = errors.Wrap(err, "failed to precreate")
	}
	u.UPR = &resp
}

// superFile 分片上传（每个切片普通用户最大4M，会员16M，超级会员32M）
func (u *Upload) superFile() {
	var resp response.UploadSuperFileResp
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	if len(u.SFIS) <= 0 {
		u.Err = errors.WithMessage(u.Err, "文件切片不存在")
		return
	}
	reqUrl := fmt.Sprintf("https://d.pcs.baidu.com/rest/2.0/pcs/superfile2?method=upload&access_token=%v"+
		"&type=tmpfile&path=%v&uploadid=%v", u.GetToken(), u.UR.RemotePath, u.UPR.Uploadid)
	for i := 0; i < len(u.SFIS); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			s := u.SFIS[i]
			err := helper.HttpPostFile(u.HttpClient(), s.SliceFilePath,
				fmt.Sprintf(reqUrl+"&partseq=%v", i),
				&resp)
			if err != nil {
				u.Err = errors.Wrap(err, "failed to upload file")
			}
			u.USFR = append(u.USFR, &resp)
			u.finish.Add(1)
			if u.UR.SliceUploadFunc != nil {
				u.UR.SliceUploadFunc(len(u.SFIS), int(u.finish.Load()), *s)
			}

		}(i)
	}
	wg.Wait()

}

// create 创建文件
func (u *Upload) create() {
	vs := url.Values{}
	vs.Add("path", u.UR.RemotePath)
	vs.Add("size", fmt.Sprintf("%v", u.UR.Size))
	vs.Add("isdir", u.UR.Isdir)
	vs.Add("block_list", strings.Replace(fmt.Sprintf("%q", u.BlockList), " ", ",", -1))
	if u.UPR != nil {
		vs.Add("uploadid", u.UPR.Uploadid)
	}
	vs.Add("rtype", u.UR.Rtype)
	vs.Add("local_ctime", u.UR.LocalCtime)
	vs.Add("local_mtime", u.UR.LocalMtime)
	var resp *response.UploadCreateFileResp
	err := helper.HttpPostForm(u.HttpClient(),
		fmt.Sprintf("https://pan.baidu.com/rest/2.0/xpan/file?method=create&access_token=%v", u.GetToken()),
		vs,
		&resp)
	if err != nil {
		u.Err = errors.WithStack(err)
	}
}

// splitFile 分割文件
func (u *Upload) splitFile(path string) (clean func()) {
	var err error
	path, err = filepath.Abs(path)
	if err != nil {
		u.Err = errors.WithStack(err)
	}
	t, err := os.MkdirTemp(".", "")
	if err != nil {
		u.Err = errors.WithStack(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		u.Err = errors.WithStack(err)
	}

	f, err := os.Open(path)
	if err != nil {
		u.Err = errors.WithStack(err)
	}
	defer f.Close()
	buf := make([]byte, u.SliceSize)
	for i := int64(0); i <= info.Size()/u.SliceSize; i++ {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			u.Err = errors.WithStack(err)
		}
		sp := fmt.Sprintf("%s/%s_%d", t, info.Name(), i+1)
		err = os.WriteFile(sp, buf[:n], os.ModePerm)
		if err != nil {
			u.Err = errors.WithStack(err)
		}
		m5 := fmt.Sprintf("%x", md5.Sum(buf[:n]))
		u.SFIS = append(u.SFIS, &request.SliceFileInfo{
			SliceName:     fmt.Sprintf("%s_%d", info.Name(), i+1),
			FileName:      info.Name(),
			Size:          int64(n),
			MD5:           m5,
			SliceFilePath: sp,
			ID:            i,
		})
		u.BlockList = append(u.BlockList, m5)
	}

	return func() {
		err := os.RemoveAll(t)
		if err != nil {
			u.Err = errors.WithStack(err)
		}
	}
}

func (u *Upload) encrypt() (clean func()) {
	if u.UR.Aes != nil && len(u.UR.Aes.Key) > 0 {
		prefix := "[aes_ofb]"
		err := helper.AesEncrypt(u.UR.SrcPath, u.UR.Aes.Key, u.UR.Aes.Iv)
		if err != nil {
			u.Err = errors.WithStack(err)
		}
		u.UR.SrcPath = filepath.Dir(u.UR.SrcPath) +
			string(filepath.Separator) +
			prefix + filepath.Base(u.UR.SrcPath)
		u.UR.RemotePath = helper.PathSeparatorFormat(filepath.Dir(u.UR.RemotePath)) +
			"/" + prefix + filepath.Base(u.UR.RemotePath)

		//删除加密的临时文件
		return func() {
			err := os.Remove(u.UR.SrcPath)
			u.Err = errors.WithStack(err)
		}
	}
	return func() {}
}

// Upload 递归上传，这个方法有待优化，考虑了重复文件重命名的情况
// BUG(upload) 目前协程并未生效，因为整个方法都被锁住了，待后续修改
func (u *Upload) Upload(reqs ...*request.UploadReq) (err error) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for _, req := range reqs {
		wg.Add(1)
		go func(req *request.UploadReq) {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			req.SrcPath, err = filepath.Abs(req.SrcPath)
			if err != nil {
				u.Err = errors.WithMessage(err, "获取绝对路径失败")
			}
			oldRemotePath := ""
			newRemotePath := ""
			searchDir := true
			var list *response.FileListResp
			err := filepath.Walk(req.SrcPath, func(path string, info fs.FileInfo, err error) error {
				//将当前路径格式化为/，例如：c:/desktop/aaa
				path = helper.PathSeparatorFormat(path)
				//将本地文件路径格式化为/，例如：c:/desktop/
				req.SrcPath = helper.PathSeparatorFormat(req.SrcPath)
				//将远程文件路径格式化为/，例如：/apps/novel
				req.RemotePath = helper.PathSeparatorFormat(req.RemotePath)
				//保留用户传入的远程和本地的根路径，防止在递归搜索的时候将其改动了无法恢复。
				_remote := req.RemotePath
				_path := req.SrcPath
				//防止一个目录下多次调用查询接口，因此一个目录只调用一次
				if searchDir {
					//查看远程根目录下面的文件
					list, err = u.FileList(&request.FileListReq{Dir: req.RemotePath})
					if err != nil {
						u.Err = errors.WithMessage(err, "查看远程根目录下面的文件失败")
					}
					searchDir = false
				}
				//设置搜索到的路径，路径为远程根路径+本地去掉根路径
				//例如远程路径为：/apps/novel ，本地根路径为c:/desktop/  本地文件路径为：c:/desktop/filefolder，
				//那么新的远程路径就是/apps/novel/filefolder
				req.RemotePath = req.RemotePath +
					strings.Replace(path, helper.PathSeparatorFormat(filepath.Dir(req.SrcPath)), "", 1)
				//本地的路径变更为walk递归的路径，确保准确找到文件
				req.SrcPath = path
				//判断当前文件是否是文件夹
				if info.IsDir() {
					//迭代远程根目录下的所有文件的路径
					for _, v := range list.List {
						//如果远程文件的路径有和设置的远程路径一致时，则说明文件夹已经存在，那么就将冲突的文件夹进行临时存储，
						//并且根据时间戳创建一个新的文件夹
						if v.Path == req.RemotePath {
							oldRemotePath = req.RemotePath
							newRemotePath = req.RemotePath + fmt.Sprintf("_%v", time.Now().Unix())
						}
					}
				}
				//将设置的远程冲突的文件夹进行替换为新的文件夹
				req.RemotePath = strings.Replace(req.RemotePath, oldRemotePath, newRemotePath, -1)
				err = u.upload(req)
				if err != nil {
					u.Err = errors.WithStack(err)
				}
				//恢复用户的根路径
				req.RemotePath = _remote
				req.SrcPath = _path
				return nil
			})
			if err != nil {
				u.Err = errors.WithMessage(err, "路径搜索失败")
			}
		}(req)
	}
	wg.Wait()
	return u.Err
}
