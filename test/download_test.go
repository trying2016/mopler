package test

import (
	"fmt"
	"github.com/766800551/mopler/model/response"
	"testing"
)

func TestDownload(t *testing.T) {
	err := sdk.Download("dddd", "/apps/novel/upload", func(fms response.FileMetaResp) {
		fmt.Printf("正在下载：%v，文件大小：%v\n", fms.Filename, fms.Size)
	})
	if err != nil {
		t.Fatal(err)
	}
	err = sdk.Download("eee", "/apps/novel/upload/[aes_ofb]chromedriver_win32.zip", func(fms response.FileMetaResp) {
		fmt.Printf("正在下载：%v，文件大小：%v\n", fms.Filename, fms.Size)
	})
	if err != nil {
		t.Fatal(err)
	}
}
