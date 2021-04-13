package model

const (
	MEDIA_TYPE_TEXT = 1  //文本样式  {id:1,userid:2,dstid:3,cmd:10,media:1,content:"hello"}
	MEDIA_TYPE_NEWS = 2 //图文样式  {id:1,userid:2,dstid:3,cmd:10,media:2,content:"标题",pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/dsturl","memo":"这是描述"}
	MEDIA_TYPE_VOICE = 3 //语音样式 amount单位秒 {id:1,userid:2,dstid:3,cmd:10,media:3,url:"http://www.a,com/dsturl.mp3",anount:40}
	MEDIA_TYPE_IMG = 4 //图片样式  {id:1,userid:2,dstid:3,cmd:10,media:4,url:"http://www.baidu.com/a/log,jpg"}
	MEDIA_TYPE_RED_PACKAGR = 5 //红包样式 红包amount 单位分 {id:1,userid:2,dstid:3,cmd:10,media:5,url:"http://www.baidu.com/a/b/c/redpackageaddress?id=100000","amount":300,"memo":"恭喜发财"}
	MEDIA_TYPE_EMOJ = 6 //emoj表情 {id:1,userid:2,dstid:3,cmd:10,media:6,"content":"cry"}
	MEDIA_TYPE_LINK = 7 //链接 {id:1,userid:2,dstid:3,cmd:10,media:7,"url":"http://www.a,com/dsturl.html"}
	MEDIA_TYPE_VIDEO = 8 //视频 {id:1,userid:2,dstid:3,cmd:10,media:8,pic:"http://www.baidu.com/a/log,jpg",url:"http://www.a,com/a.mp4"}
	MEDIA_TYPE_CONTACT = 9 //名片 {id:1,userid:2,dstid:3,cmd:10,media:9,"content":"10086","pic":"http://www.baidu.com/a/avatar,jpg","memo":"胡大力"}
)

//消息体
type Message struct {
	Id      int64  `xorm:"pk autoincr bigint(20)" json:"id,omitempty" form:"id"`           //消息ID
	Userid  int64  `xorm:"bigint(20)" json:"userid,omitempty" form:"userid"`   //谁发的
	Cmd     int    `xorm:"int(11)" json:"cmd,omitempty" form:"cmd"`         //群聊还是私聊
	Dstid   int64  `xorm:"bigint(20)" json:"dstid,omitempty" form:"dstid"`     //对端用户ID/群ID
	Media   int    `xorm:"int(11)" json:"media,omitempty" form:"media"`     //消息按照什么样式展示
	Content string `xorm:"varchar(1000)" json:"content,omitempty" form:"content"` //消息的内容
	Pic     string `xorm:"varchar(255)" json:"pic,omitempty" form:"pic"`         //预览图片
	Url     string `xorm:"varchar(255)" json:"url,omitempty" form:"url"`         //服务的URL
	Memo    string `xorm:"varchar(255)" json:"memo,omitempty" form:"memo"`       //简单描述
	Amount  int    `xorm:"int(11)" json:"amount,omitempty" form:"amount"`   //其他和数字相关的
}