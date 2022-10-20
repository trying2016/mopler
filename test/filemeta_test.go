package test

import (
	"fmt"
	"github.com/766800551/mopler/model/request"
	"testing"
)

func TestFileMeta(t *testing.T) {
	meta, err := sdk.FileMeta(&request.FileMetaReq{Fsids: []int64{95307883004145}})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", meta)
}
