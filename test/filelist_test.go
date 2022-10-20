package test

import (
	"fmt"
	"github.com/766800551/mopler/model/request"
	"testing"
)

func TestFileList(t *testing.T) {
	list, err := sdk.FileList(&request.FileListReq{Dir: "/apps/novel"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", list)
}

func TestFileListRecursion(t *testing.T) {
	list, err := sdk.FileListRecursion(&request.FileListRecursionReq{Path: "/apps/novel"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", list)
}
