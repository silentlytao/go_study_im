package service

import (
	"../model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var DbEngine *xorm.Engine

func init()  {
	var err error
	DbEngine,err = xorm.NewEngine("mysql","root:root@(127.0.0.1:3306)/go_im?charset=utf8")
	if err != nil{
		log.Fatal(err.Error())
	}
	//显示打印的SQL
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	fmt.Println("xorm init success")
	DbEngine.Sync2(model.User{},model.Community{},model.Contact{})
}

