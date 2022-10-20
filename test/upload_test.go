package test

import (
	"crypto/aes"
	"fmt"
	"github.com/766800551/mopler/model/request"
	"testing"
)

func TestPanUploadRecursion(t *testing.T) {
	err := sdk.Upload(
		&request.UploadReq{
			SrcPath:    "C:\\Users\\76680\\GolandProjects\\mopler\\test\\resoure\\upload",
			RemotePath: "/apps/novel",
			SliceUploadFunc: func(count int, finish int, s request.SliceFileInfo) {
				fmt.Printf("【%v】总共：%v个切片， 已经完成了：%v个，"+
					"正在上传切片编号：%v\n", s.SliceName, count, finish, s.ID)
			},
			Aes: &request.Aes{
				Key: []byte("2pfkl36zyctsc0ki"),
				Iv:  [aes.BlockSize]byte{'c', 's', 'd', 'r', 'g', 'e', 'x', 'y', 's', 'h', 'l', '6', 'u', 'b', '6', 'y'},
			},
		},
		//&request.UploadReq{
		//	SrcPath:    "C:\\Users\\76680\\Desktop\\新建文件夹",
		//	RemotePath: "/apps/novel",
		//	SliceUploadFunc: func(count int, finish int, s request.SliceFileInfo) {
		//		fmt.Printf("【%v】总共：%v个切片， 已经完成了：%v个，"+
		//			"正在上传切片编号：%v\n", s.SliceName, count, finish, s.ID)
		//	},
		//	Aes: &request.Aes{
		//		Key: []byte("2pfkl36zyctsc0ki"),
		//		Iv:  [aes.BlockSize]byte{'c', 's', 'd', 'r', 'g', 'e', 'x', 'y', 's', 'h', 'l', '6', 'u', 'b', '6', 'y'},
		//	},
		//},
		//&request.UploadReq{
		//	SrcPath:    "C:\\Users\\76680\\Desktop\\ccc.xlsx",
		//	RemotePath: "/apps/novel",
		//	SliceUploadFunc: func(count int, finish int, s request.SliceFileInfo) {
		//		fmt.Printf("【%v】总共：%v个切片， 已经完成了：%v个，"+
		//			"正在上传切片编号：%v\n", s.SliceName, count, finish, s.ID)
		//	},
		//	Aes: &request.Aes{
		//		Key: []byte("2pfkl36zyctsc0ki"),
		//		Iv:  [aes.BlockSize]byte{'c', 's', 'd', 'r', 'g', 'e', 'x', 'y', 's', 'h', 'l', '6', 'u', 'b', '6', 'y'},
		//	},
		//},
	)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
