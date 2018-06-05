package main

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"fmt"
	"mime"
	"os"
	"admin/models"
	_ "admin/routers"
)

const VERSION = "0.0.1"



func initialize() {
	mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	initArgs()

	models.Connect()

}
func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}


func main() {
	//初始化
	initialize()

	fmt.Println("Starting....")

	fmt.Println("Start ok")
	// 开启 ORM 调试模式
	//orm.Debug = true
	// 自动建表
	//orm.RunSyncdb("default", false, true)
	beego.AddFuncMap("stringsToJson", models.StringsToJson)
	//
	// 运行时
	beego.Run()
}