package response

type FileMetasResp struct {
	Errmsg string         `json:"errmsg"`
	Errno  int            `json:"errno"`
	List   []FileMetaResp `json:"list"`
	Names  struct {
	} `json:"names"`
	RequestId string `json:"request_id"`
}

type FileMetaResp struct {
	Category    int    `json:"category"`
	Filename    string `json:"filename"`
	FsId        int64  `json:"fs_id"`
	Isdir       int    `json:"isdir"`
	Md5         string `json:"md5"`
	OperId      int    `json:"oper_id"`
	Path        string `json:"path"`
	RverCtime   int    `json:"rver_ctime"`
	ServerMtime int    `json:"server_mtime"`
	Size        int    `json:"size"`
	Dlink       string `json:"dlink"`
}
