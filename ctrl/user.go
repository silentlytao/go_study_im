package ctrl

import (
	"../service"
	"../util"
	"net/http"
	"strconv"
)

var UserService *service.UserService

//登录
func UserLogin(writer http.ResponseWriter, request *http.Request) {
	//解析参数
	request.ParseForm()
	mobile := request.Form.Get("mobile")
	pwd := request.Form.Get("password")

	user, err := UserService.Login(mobile, pwd)
	if err != nil {
		util.RespFail(writer, "账号或密码错误")
	} else {
		util.RespOk(writer, user, "登录成功")
	}
}

//注册
func UserRegister(writer http.ResponseWriter, request *http.Request) {
	//解析参数
	request.ParseForm()
	mobile := request.Form.Get("mobile")
	pwd := request.Form.Get("password")
	avatar := request.Form.Get("avatar")
	sex := request.Form.Get("sex")
	nickName := request.Form.Get("nick_name")
	_, err := UserService.Register(mobile, pwd, nickName, avatar, sex)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, "", "注册成功")
	}
}

//查找用户信息
func Find(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	id := request.Form.Get("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		util.RespFail(writer, err.Error())
	}
	user := UserService.Find(userId)
	util.RespOk(writer, user, "")
}
