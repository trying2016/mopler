package test

import (
	"fmt"
	"github.com/766800551/mopler/model/request"
	"testing"
)

func TestFileSearch(t *testing.T) {
	search, err := sdk.FileSearch(&request.FileSearchReq{
		Key: "新建",
		Dir: "/apps/novel/a",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", search)
}
