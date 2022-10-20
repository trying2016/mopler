package request

type FileListReq struct {
	//需要list的目录，以/开头的绝对路径, 默认为/路径包含中文时需要UrlEncode编码给出的示例的路径是/测试目录的UrlEncode编码
	Dir string
	/*
		排序字段：默认为name；
			time表示先按文件类型排序，后按修改时间排序；
			name表示先按文件类型排序，后按文件名称排序；
			size表示先按文件类型排序，后按文件大小排序。
	*/
	Order     string
	Desc      string //  默认为升序，设置为1实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start     string //  起始位置，从0开始
	Limit     string //    查询数目，默认为1000，建议最大不超过1000
	Web       string // 值为1时，返回dir_empty属性和缩略图数据
	Folder    string //  是否只返回文件夹，0 返回所有，1 只返回文件夹，且属性只返回path字段
	Showempty string //   是否返回dir_empty属性，0 不返回，1 返回
}

type FileListRecursionReq struct {
	//需要list的目录，以/开头的绝对路径, 默认为/路径包含中文时需要UrlEncode编码给出的示例的路径是/测试目录的UrlEncode编码
	Path string
	/*
		排序字段：默认为name；
			time表示先按文件类型排序，后按修改时间排序；
			name表示先按文件类型排序，后按文件名称排序；
			size表示先按文件类型排序，后按文件大小排序。
	*/
	Order string
	Desc  string //  默认为升序，设置为1实现降序 （注：排序的对象是当前目录下所有文件，不是当前分页下的文件）
	Start string //  起始位置，从0开始
	Limit string //    查询数目，默认为1000，建议最大不超过1000
	Web   string // 值为1时，返回dir_empty属性和缩略图数据
}
