// Package quota 实现获取网盘容量
package quota

import (
	"fmt"
	"github.com/766800551/mopler/api"
	"github.com/766800551/mopler/helper"
	"github.com/766800551/mopler/model/request"
	"github.com/766800551/mopler/model/response"
	"github.com/pkg/errors"
)

type Quota struct {
	api.Sdker
}

func NewQuotaImpl(sdk api.Sdker) *Quota {
	return &Quota{
		Sdker: sdk,
	}
}

// Quota 获取网盘信息
func (q *Quota) Quota(req *request.QuotaReq) (*response.QuotaResp, error) {
	var resp response.QuotaResp
	err := helper.HttpGet(q.HttpClient(),
		fmt.Sprintf("https://pan.baidu.com/api/quota?access_token=%vcheckfree=%v&checkexpire=%v",
			q.GetToken(), req.Checkfree, req.Checkexpire),
		&resp)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting quota information")
	}
	return &resp, nil
}
