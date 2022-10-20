package request

type QuotaReq struct {
	//	是否检查免费信息，0为不查，1为查，默认为0
	Checkfree string `json:"checkfree"`
	//		是否检查过期信息，0为不查，1为查，默认为0
	Checkexpire string `json:"checkexpire"`
}
