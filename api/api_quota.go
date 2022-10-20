package api

import (
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
)

// Quotaer 网盘容量
type Quotaer interface {
	Quota(*request.QuotaReq) (*response.QuotaResp, error)
}
