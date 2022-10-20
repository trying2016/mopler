package response

type UploadPrecreateResp struct {
	Errno      int    `json:"errno"`       //错误码
	Path       string `json:"path"`        //文件的绝对路径
	Uploadid   string `json:"uploadid"`    //上传唯一ID标识此上传任务
	ReturnType int    `json:"return_type"` //返回类型，系统内部状态字段
	BlockList  []int  `json:"block_list"`  //需要上传的分片序号列表，索引从0开始
}

type UploadSuperFileResp struct {
	Errno     int    `json:"errno"`
	RequestId int    `json:"request_id"`
	Errmsg    string `json:"errmsg"`
	Md5       string `json:"md5"`
}

type UploadCreateFileResp struct {
	Errno          int    `json:"errno"`           //错误码
	FsId           int64  `json:"fs_id"`           //文件在云端的唯一标识ID
	Md5            string `json:"md_5"`            //文件的MD5，只有提交文件时才返回，提交目录时没有该值
	ServerFilename string `json:"server_filename"` //文件名
	Category       int    `json:"category"`        //分类类型, 1 视频 2 音频 3 图片 4 文档 5 应用 6 其他 7 种子
	Path           string `json:"path"`            //上传后使用的文件绝对路径
	Size           int64  `json:"size"`            //文件大小，单位B
	Ctime          int64  `json:"ctime"`           //文件创建时间
	Mtime          int64  `json:"mtime"`           //文件修改时间
	Isdir          int    `json:"isdir"`           //是否目录，0 文件、1 目录
	Name           string `json:"name"`
	FromType       int    `json:"from_type"`
}
