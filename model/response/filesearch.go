package response

type FileSearchResp struct {
	Errno       int      `json:"errno"`
	RequestId   any      `json:"request_id"`
	Errmsg      any      `json:"errmsg"`
	Contentlist []string `json:"contentlist"`
	HasMore     int      `json:"has_more"`
	List        []struct {
		FsID           int64               `json:"fs_id"`
		Path           string              `json:"path"`
		ServerFilename string              `json:"server_filename"`
		Size           int                 `json:"size"`
		ServerMtime    int                 `json:"server_mtime"`
		ServerCtime    int                 `json:"server_ctime"`
		LocalMtime     int                 `json:"local_mtime"`
		LocalCtime     int                 `json:"local_ctime"`
		Isdir          int                 `json:"isdir"`
		Category       int                 `json:"category"`
		Share          int                 `json:"share"`
		OperID         int                 `json:"oper_id"`
		ExtentTinyint1 int                 `json:"extent_tinyint_1"`
		Md5            string              `json:"md_5"`
		Thumbs         []map[string]string `json:"thumbs"`
		DeleteType     int                 `json:"delete_type"`
		OwnerId        int                 `json:"owner_id"`
		Wpfile         int                 `json:"wpfile"`
	} `json:"list"`
}
