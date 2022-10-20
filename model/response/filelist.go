package response

type FileListResp struct {
	Errno     int    `json:"errno"`
	GuidInfo  string `json:"guid_info"`
	RequestId int64  `json:"request_id"`
	Errmsg    string `json:"errmsg"`
	Guid      int64  `json:"guid"`
	HasMore   int    `json:"has_more"`
	Cursor    int    `json:"cursor"`
	List      []struct {
		FsId           int64               `json:"fs_id"`           //文件在云端的唯一标识ID
		Path           string              `json:"path"`            //文件的绝对路径
		ServerFilename string              `json:"server_filename"` //文件名称
		Size           int64               `json:"size"`            //文件大小，单位B
		ServerMtime    int64               `json:"server_mtime"`    //文件在服务器修改时间
		ServerCtime    int64               `json:"server_ctime"`    //文件在服务器创建时间
		LocalMtime     int64               `json:"local_mtime"`     //文件在客户端修改时间
		LocalCtime     int64               `json:"local_ctime"`     //文件在客户端创建时间
		Isdir          int                 `json:"isdir"`           //是否为目录，0 文件、1 目录
		Category       int                 `json:"category"`        //文件类型，1 视频、2 音频、3 图片、4 文档、5 应用、6 其他、7 种子
		Md5            string              `json:"md_5"`            //云端哈希（非文件真实MD5），只有是文件类型时，该字段才存在
		DirEmpty       int                 `json:"dir_empty"`       //该目录是否存在子目录，只有请求参数web=1且该条目为目录时，该字段才存在， 0为存在， 1为不存在
		Thumbs         []map[string]string `json:"thumbs"`          //只有请求参数web=1且该条目分类为图片时，该字段才存在，包含三个尺寸的缩略图URL
		Privacy        string              `json:"privacy"`
		Unlist         int64               `json:"unlist"`
		ServerAtime    int64               `json:"server_atime"`
		Share          int64               `json:"share"`
		Empty          int64               `json:"empty"`
		OperId         int64               `json:"oper_id"`
	} `json:"list"`
}
