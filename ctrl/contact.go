package ctrl

import (
	"../args"
	"../model"
	"../service"
	"../util"
	"net/http"
)

var ContactService *service.ContactService

//查找好友
func LoadFriend(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//绑定参数
	err := util.Bind(request, &arg)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		users := ContactService.SearchFriend(arg.Userid)
		util.RespOk(writer, users, "ok")
	}
}

//添加好友
func AddFriend(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//绑定参数
	err := util.Bind(request, &arg)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		err = ContactService.AddFriend(arg.Userid, arg.Dstid)
		if err != nil {
			util.RespFail(writer, err.Error())
		} else {
			util.RespOk(writer, "", "ok")
		}
	}
}

//搜索群列表
func LoadCommunity(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//绑定参数
	err := util.Bind(request, &arg)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		users := ContactService.SearchCommunity(arg.Userid)
		util.RespOk(writer, users, "ok")
	}

}

//创建群
func CreateCommunity(writer http.ResponseWriter, request *http.Request) {
	var arg model.Community
	//绑定参数
	util.Bind(request, &arg)
	users, err := ContactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(writer, err.Error())
	}
	util.RespOk(writer, users, "ok")
}

//加入群
func JoinCommunity(writer http.ResponseWriter, request *http.Request) {
	var arg args.ContactArg
	//绑定参数
	util.Bind(request, &arg)
	users := ContactService.JoinCommunity(arg.Userid, arg.Dstid)
	//刷新用户的群组信息
	AddGroupId(arg.Userid, arg.Dstid)
	util.RespOk(writer, users, "ok")
}

func AddGroupId(userId, gid int64) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	rwLocker.RUnlock()
}
