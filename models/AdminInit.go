package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
)

var o orm.Ormer

func Syncdb()  {
	createdb()
	Connect()
	o = orm.NewOrm()
	
	name := "default"
	
	force := true
	
	verbose := true
	
	err := orm.RunSyncdb(name,force,verbose)
	
	if err == nil {
		fmt.Println(err)
	}
	
	insertUser()
	insertGroup()
	insertRole()
	insertNodes()
	fmt.Println("database init is complete.\nPlease restart the application")
}

//数据库链接
func Connect() {
	var dsn string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dsn = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
	case "sqlite3":
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if db_path == "" {
			db_path = "./"
		}
		dsn = fmt.Sprintf("%s%s.db", db_path, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	orm.RegisterDataBase("default", db_type, dsn)
}
//创建表
func createdb() {
	
}

func insertUser() {
	
}

func insertGroup() {
	
}

func insertRole() {
	
}

func insertNodes()  {
	
}


