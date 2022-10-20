package test

import (
	"fmt"
	"github.com/766800551/mopler/model/request"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := sdk.CreateDir("vvvv", "/apps/novel")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileCopy(t *testing.T) {
	resp, err := sdk.Copy(&request.FileManagerReq{
		Async: "1",
		Filelist: []request.FileManagerInfo{
			{
				Path:    "/apps/novel/aa",
				Dest:    "/apps/novel",
				Newname: "bbb",
				Ondup:   "fail",
			},
		},
		Ondup: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)
}

func TestFileRename(t *testing.T) {
	resp, err := sdk.Rename(&request.FileManagerReq{
		Async: "1",
		Filelist: []request.FileManagerInfo{
			{
				Path:    "/apps/novel/aa",
				Newname: "ccc",
			},
		},
		Ondup: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)
}

func TestFileMove(t *testing.T) {
	resp, err := sdk.Move(&request.FileManagerReq{
		Async: "1",
		Filelist: []request.FileManagerInfo{
			{
				Path:    "/apps/novel/ccc",
				Dest:    "/apps/novel/bbb",
				Newname: "bbb",
				Ondup:   "fail",
			},
		},
		Ondup: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)
}

func TestFileDelete(t *testing.T) {
	resp, err := sdk.Delete(&request.FileManagerReq{
		Async: "1",
		Filelist: []request.FileManagerInfo{
			{
				Path: "/apps/novel/aaa",
			},
		},
		Ondup: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", resp)
}
