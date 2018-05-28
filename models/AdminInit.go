package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
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
	insertRole()
	insertUserRole()
	insertPower()
	insertRolePower()
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
//创建数据库
func createdb() {
	
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	
	var dsn string
	var sqlstring string
	switch db_type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
		break
	case "postgres":
		dsn = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_host, db_user, db_pass, db_port, db_sslmode)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
		break
	case "sqlite3":
		if db_path == "" {
			db_path = "./"
		}
		dsn = fmt.Sprintf("%s%s.db", db_path, db_name)
		os.Remove(dsn)
		sqlstring = "create table init (n varchar(32));drop table init;"
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	db, err := sql.Open(db_type, dsn)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()
	
}

//insert into user
func insertUser() {
	fmt.Println("insert into user ...")
	u := new(User)
	u.Username = "admin"
	u.Password = Pwdhash("admin")
	u.Nickname = "admin"
	u.Status  = 2
	u.Email = "995645653@qq.com"
	o := orm.NewOrm()
	o.Insert(u)
	fmt.Println("insert into user end")
}
//insert into user_role 派生表
func insertUserRole() {
	fmt.Println("insert into user_role ...")
	u := new(UserRole)
	u.User_id = 1
	u.Role_id = 1
	o := orm.NewOrm()
	o.Insert(u)
	fmt.Println("insert into user_role end")
}

// insert into role
func insertRole() {
	fmt.Println("insert into role ...")
	u := new(Role)
	u.Rolename = "超级管理员"
	o := orm.NewOrm()
	o.Insert(u)
	fmt.Println("insert into role end")
}

//insert into role_power 派生表
func insertRolePower() {
	fmt.Println("insert into role_power ...")
	//u := new(RolePower)
	//u.Role_id = 1
	//u.Power_id = 1
	//o := orm.NewOrm()
	//o.Insert(u)
	fmt.Println("insert into role_power end")
}

//insert into power
func insertPower() {
	fmt.Println("insert into power ...")
	power := [16]Power{
		{Controller:"admin",Action:"index",Powername:"后台管理员管理",Pid:0,Level:1,Status:2}, //id 1
		{Controller:"admin/user",Action:"user",Powername:"用户列表",Pid:1,Level:1,Status:2},  //id 2
		{Controller:"admin/role",Action:"role",Powername:"角色列表",Pid:1,Level:1,Status:2},  //id 3
		{Controller:"admin/power",Action:"power",Powername:"权限列表",Pid:1,Level:1,Status:2}, //id 4
		{Controller:"admin/user",Action:"edit",Powername:"用户列表",Pid:2,Level:2,Status:2},
		{Controller:"admin/user",Action:"add",Powername:"用户列表",Pid:2,Level:2,Status:2},
		{Controller:"admin/user",Action:"upd",Powername:"用户列表",Pid:2,Level:2,Status:2},
		{Controller:"admin/user",Action:"del",Powername:"用户列表",Pid:2,Level:2,Status:2},
		{Controller:"admin/role",Action:"edit",Powername:"角色列表",Pid:3,Level:2,Status:2},
		{Controller:"admin/role",Action:"add",Powername:"角色列表",Pid:3,Level:2,Status:2},
		{Controller:"admin/role",Action:"upd",Powername:"角色列表",Pid:3,Level:2,Status:2},
		{Controller:"admin/role",Action:"del",Powername:"角色列表",Pid:3,Level:2,Status:2},
		{Controller:"admin/power",Action:"edit",Powername:"权限列表",Pid:4,Level:2,Status:2},
		{Controller:"admin/power",Action:"add",Powername:"权限列表",Pid:4,Level:2,Status:2},
		{Controller:"admin/power",Action:"upd",Powername:"权限列表",Pid:4,Level:2,Status:2},
		{Controller:"admin/power",Action:"del",Powername:"权限列表",Pid:4,Level:2,Status:2},
	}
	
	for _,v := range power  {
		n := new(Power)
		n.Powername = v.Powername
		n.Controller = v.Controller
		n.Action = v.Action
		n.Pid = v.Pid
		n.Status = v.Status
		n.Level = v.Level
		o.Insert(n)
		
		u := new(RolePower)
		u.Role_id = 1
		u.Power_id = v.Id
		o.Insert(u)
	}
	fmt.Println("insert into power end")
}


