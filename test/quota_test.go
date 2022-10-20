package test

import (
	"fmt"
	"github.com/766800551/mopler/model/request"
	"testing"
)

func TestQuota(t *testing.T) {
	quota, err := sdk.Quota(&request.QuotaReq{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", quota)
}
