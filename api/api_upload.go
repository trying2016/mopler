package api

import "github.com/766800551/mopler/model/request"

// Uploader 文件上传
type Uploader interface {
	Upload(...*request.UploadReq) error
}
