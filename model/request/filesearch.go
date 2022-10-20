package request

type FileSearchReq struct {
	Key       string `json:"key"`       //	搜索关键字
	Dir       string `json:"dir"`       //	搜索目录，默认根目录
	Page      string `json:"page"`      //	页数，从1开始，缺省则返回所有条目
	Num       string `json:"num"`       //	默认为500，不能修改
	Recursion string `json:"recursion"` //	是否递归搜索子目录 1:是，0:否（默认）
	Web       string `json:"web"`       //	默认0，为1时返回缩略图信息
}
