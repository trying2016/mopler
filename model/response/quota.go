package response

type QuotaResp struct {
	Total     int    `json:"total"`  //总空间大小，单位B
	Expire    bool   `json:"expire"` //7天内是否有容量到期
	Used      int    `json:"used"`   //已使用大小，单位B
	Free      int    `json:"free"`   //剩余大小，单位B
	Errno     int    `json:"errno"`
	RequestId int    `json:"request_id"`
	Errmsg    string `json:"errmsg"`
}
