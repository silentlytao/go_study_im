package util

import (
	"encoding/json"
	"log"
	"net/http"
)

//定义返回结构体
type Response struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
//失败
func RespFail(w http.ResponseWriter ,msg string)  {
	Resp(w,0,"",msg)
}
//成功
func RespOk(w http.ResponseWriter ,data interface{},msg string)  {
	Resp(w,1,data,msg)
}


//返回内容
func Resp(w http.ResponseWriter ,code int,data interface{},msg string)  {
	//返回json
	w.Header().Set("Content-Type","application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	response := Response{
		Code : code,
		Msg : msg,
		Data : data,
	}
	//转成json
	response_json,error := json.Marshal(response)
	if error != nil{
		log.Printf("返回错误%s",error.Error())
	}
	w.Write(response_json)
}
