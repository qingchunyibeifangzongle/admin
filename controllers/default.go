package controllers

import (
	"github.com/astaxie/beego"
	//"admin/models"
	//"github.com/astaxie/beego/orm"
	
	"time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	const DATE_LAYOUT = "2006-01-02 15:04:05"
	now := time.Now()
	now_str := now.Format(DATE_LAYOUT)
	beego.Info(now_str)
	this.ServeJSON()
	this.TplName = "amz/login.html"
	//models.Connect()
	//o := orm.NewOrm()
	//var powers []orm.Params
	//var aaaaa []orm.Params
	//power := new(models.Power)
	//names := [5]models.RolePower{{Power_id:5},{Power_id:1},{Power_id:2},{Power_id:3},{Power_id:4}}
	//for _,n := range names {
	//	o.QueryTable(power).Filter("id",n.Power_id).Values(&powers)
	//	for _,v := range powers {
	//		aaaaa = append(aaaaa, v)
	//	}
	//}
	//
	//
	//
	//beego.Info(aaaaa)
	//this.ServeJSON()
}
