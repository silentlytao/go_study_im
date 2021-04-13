package service

import (
	"../model"
	"../util"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

//用户服务
type UserService struct {

}
//注册
func (u *UserService) Register(mobile,password,nickName,avatar,sex string) (user model.User,err error) {
	tmp := model.User{}
	//查询手机号对应用户
	_,err = DbEngine.Where("mobile=?",mobile).Get(&tmp)
	if err != nil{
		log.Fatal(err.Error())
		return tmp,err
	}
	//如果存在则返回提示已经注册
	if tmp.Id > 0 {
		return tmp, errors.New("该手机号已存在")
	}
	//插入数据库
	tmp.Mobile = mobile
	tmp.Avatar = avatar
	tmp.Nickname = nickName
	if sex == "" {
		sex = model.SexNukown
	}
	tmp.Sex = sex
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tmp.Passwd = util.MakePasswd(password, tmp.Salt)
	tmp.Createat = time.Now()
	//token 可以是一个随机数
	tmp.Token = fmt.Sprintf("%08d", rand.Int31())

	_,err = DbEngine.InsertOne(&tmp)
	return tmp,nil
}
//登录
func (u *UserService) Login(mobile,password string) (user model.User,err error) {
	tmp := model.User{}
	//查询手机号对应用户
	_,err = DbEngine.Where("mobile=?",mobile).Get(&tmp)
	if err != nil{
		log.Fatal(err.Error())
		return tmp,err
	}
	//验证密码
	if util.ValidatePasswd(password,tmp.Salt,tmp.Passwd) != true{
		return tmp,errors.New("账号或密码错误")
	}
	//刷新token
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.MD5Encode(str)
	tmp.Token = token
	//返回数据
	DbEngine.ID(tmp.Id).Cols("token").Update(&tmp)
	return tmp,nil
}
//查找某个用户
func (u *UserService) Find(userId int64) model.User {
	//首先通过手机号查询用户
	tmp := model.User{}
	DbEngine.ID(userId).Get(&tmp)
	return tmp
}
