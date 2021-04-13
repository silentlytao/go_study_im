package ctrl

import (
	"../util"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func init()  {
	os.MkdirAll("./mnt",os.ModePerm)
}

//上传文件
func Upload(writer http.ResponseWriter,request *http.Request)  {
	UploadLocal(writer,request)
	//因为本人没有oss配置,所以并没有实现oss的上传。
	//请想用oss的同学，自行进行补充😀
}

func UploadLocal(writer http.ResponseWriter,request *http.Request)  {
	//获取上传的源文件
	srcFile,head,err := request.FormFile("file")
	if err != nil {
		util.RespFail(writer,err.Error())
		return
	}
	//创建一个新文件
	suffix := ".png"
	ofileName := head.Filename
	tmp := strings.Split(ofileName,".")
	if len(tmp) >1{
		suffix = "."+tmp[len(tmp) -1]
	}
	fileType := request.FormValue("filetype")
	if len(fileType) > 0{
		suffix = fileType
	}
	fileName := fmt.Sprintf("%d%04d%s",time.Now().Unix(),rand.Int31(),suffix)
	dsFile,err := os.Create("./mnt/"+fileName)
	if err != nil{
		util.RespFail(writer,err.Error())
		return
	}
	//将源文件内容copy到新文件
	_,err = io.Copy(dsFile,srcFile)
	if err != nil {
		util.RespFail(writer,err.Error())
		return
	}
	//将新文件路径转换成url地址
	url := "/mnt/"+fileName
	util.RespOk(writer,url,"")
}
