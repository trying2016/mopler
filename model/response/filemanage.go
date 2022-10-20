package response

type FileManagerResp struct {
	Errno     int `json:"errno"`
	RequestId any `json:"request_id"`
	Errmsg    any `json:"errmsg"`
	Info      []struct {
		Errno int    `json:"errno"`
		Path  string `json:"path"`
	} //文件信息
	Taskid uint64 //异步任务id, 当async=2时返回
}
