package args

type ContactArg struct {
	Userid int64 	`json:"userid" form:"userid"`
	Dstid int64 	`json:"dstid" form:"dstid"`
}
