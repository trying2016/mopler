package mopler

import (
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/impl/download"
	"github.com/766800551/mopler/impl/filelist"
	"github.com/766800551/mopler/impl/filemanage"
	"github.com/766800551/mopler/impl/filemeta"
	"github.com/766800551/mopler/impl/filesearch"
	"github.com/766800551/mopler/impl/quota"
	"github.com/766800551/mopler/impl/upload"
	"github.com/766800551/mopler/impl/userinfo"
	"net/http"
)

type SdkContext struct {
	api.UserInfoer
	api.Quotaer
	api.FileLister
	api.Uploader
	api.FileSearcher
	api.FileManager
	api.FileMetaer
	api.Downloader
	api.Authorizationer
	client *http.Client
}

func New(a api.Authorizationer) *SdkContext {
	s := &SdkContext{
		client: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
			},
		},
	}
	s.Authorizationer = a
	s.UserInfoer = userinfo.NewUserInfoImpl(s)
	s.Quotaer = quota.NewQuotaImpl(s)
	s.FileLister = filelist.NewFileListImpl(s)
	s.FileMetaer = filemeta.NewFileMetaImpl(s)
	s.FileSearcher = filesearch.NewFileSearchImpl(s)
	s.Uploader = upload.NewUploadImpl(s)
	s.FileManager = filemanage.NewFileManageImpl(s)
	s.Downloader = download.NewDownloadImpl(s)

	return s
}

func (s *SdkContext) HttpClient() *http.Client {
	return s.client
}
