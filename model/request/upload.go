package request

import "crypto/aes"

type UploadReq struct {
	//	需要上传到的网盘路径	RequestBody参数	上传后使用路径（不包含文件名或新的目录），需要urlencode
	//例如需要上传 c:/desktop/a.txt 到网盘/apps/novel目录下，那么:
	//  SrcPath=c:/desktop/a.txt RemotePath=/apps/novel
	RemotePath string
	SrcPath    string //本地文件的路径
	BlockSize  int64

	//	4096	RequestBody参数	文件和目录两种情况：上传文件时，表示文件的大小，单位B；上传目录时，表示目录的大小，目录的话大小默认为0
	Size int64
	//	0	RequestBody参数	是否为目录，0 文件，1 目录
	Isdir string
	/*
		["98d02a0f54781a93e354b1fc85caf488", "ca5273571daefb8ea01a42bfa5d02220"]
		RequestBody参数
		文件各分片MD5数组的json串。block_list的含义如下，如果上传的文件小于4MB，其md5值（32位小写）即为block_list字符串数组的唯一元素；
		如果上传的文件大于4MB，需要将上传的文件按照4MB大小在本地切分成分片，
		不足4MB的分片自动成为最后一个分片，所有分片的md5值（32位小写）组成的字符串数组即为block_list。
	*/
	BlockList string
	//	1	RequestBody参数	固定值1
	Autoinit string
	//	1	RequestBody参数	文件命名策略。
	/*
		1 表示当path冲突时，进行重命名
		2 表示当path冲突且block_list不同时，进行重命名
		3 当云端存在同名文件时，对该文件进行覆盖
	*/
	Rtype string

	//    P1-MTAuMjI4LjQzLjMxOjE1OTU4NTg==    RequestBody参数    上传ID
	Uploadid string
	//	content-md5  b20f8ac80063505f264e5f6fc187e69a    RequestBody参数    文件MD5，32位小写
	ContentMd5 string
	// 9aa0aa691s5c0257c5ab04dd7eddaa47    RequestBody参数    文件校验段的MD5，32位小写，校验段对应文件前256KB
	SliceMd5 string
	//1595919297    RequestBody参数    客户端创建时间， 默认为当前时间戳
	LocalCtime string
	// 1595919297    RequestBody参数    客户端修改时间，默认为当前时间戳
	LocalMtime string
	//切片上传回调，总共有多少个切片，完成了多少个，当前切片信息
	SliceUploadFunc func(countNum int, finishNum int, s SliceFileInfo)
	//文件进行Aes加密
	Aes *Aes
}

type Aes struct {
	Key []byte
	Iv  [aes.BlockSize]byte
}

type SliceFileInfo struct {
	SliceName     string
	FileName      string
	Size          int64
	MD5           string
	SliceFilePath string
	ID            int64
}
