package request

type FileManagerInfo struct {
	Path    string `json:"path,omitempty"`
	Dest    string `json:"dest,omitempty"`
	Newname string `json:"newname,omitempty"`
	Ondup   string `json:"ondup,omitempty"`
}

type FileManagerReq struct {
	Async    string            `json:"async,omitempty"`
	Filelist []FileManagerInfo `json:"filelist,omitempty"`
	Ondup    string            `json:"ondup,omitempty"`
}
