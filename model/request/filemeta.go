package request

type FileMetaReq struct {
	Fsids     []int64 //	[414244021542671,633507813519281]	文件id数组，数组中元素是uint64类型，数组大小上限是：100
	Dlink     string  //			是否需要下载地址，0为否，1为是，默认为0。获取到dlink后，参考下载文档进行下载操作
	Path      string  //	/123-571234	URL参数	查询共享目录或专属空间内文件时需要。 共享目录格式： /uk-fsid  其中uk为共享目录创建者id， fsid对应共享目录的fsid 专属空间格式：/_pcs_.appdata/xpan/
	Thumb     string  //		是否需要缩略图地址，0为否，1为是，默认为0
	Extra     string  //		图片是否需要拍摄时间、原图分辨率等其他信息，0 否、1 是，默认0
	Needmedia string  //	视频是否需要展示时长信息，0 否、1 是，默认0
}
