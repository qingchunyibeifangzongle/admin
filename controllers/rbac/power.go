/**
 * User: yp
 * Time: 2018/6/11  下午4:58
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PowerController struct {
	common.CommonController
}

func (this *PowerController) Index() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	//const pageSize = 1
	//page,_ := this.GetInt(":page")
	//if page == 0 {
	//	page = 1
	//}
	//powers , totalRows := models.Getpowerlist(page , pageSize)
	//res := models.Paginator(page, pageSize, totalRows)
	powers1,powers2,powers3 := models.GroupList()
	beego.Info(powers1)
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["powers1"] = &powers1
	this.Data["powers2"] = &powers2
	this.Data["powers3"] = &powers3
	//this.Data["paginator"] = res
	//this.Data["totals"] = totalRows
	this.TplName = this.GetTemplatetype() + "/powerlist.html"
}

func (this *PowerController) PowerEdit(){
	//userinfo := this.GetSession("userinfo")
	//if userinfo == nil {
	//	this.Ctx.Redirect(302,"/public/login")
	//}
	powerId,_ := this.GetInt64(":id")
	powers := models.GetPowerId(powerId)
	powers1,powers2,powers3 := models.GroupList()
	beego.Info(powers)
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["powers1"] = &powers1
	this.Data["powers2"] = &powers2
	this.Data["powers3"] = &powers3
	this.Data["powers"] = &powers
	this.TplName = this.GetTemplatetype() + "/poweredit.html"
}
func (this *PowerController) PowerEdits(){
	powername := this.GetString("powername")
	controller := this.GetString("controller")
	action := this.GetString("action")
	pid,_ := this.GetInt("pid")
	//var lpid int64
	//lpid = int64(pid)
	id,_ := this.GetInt64("id")
	powerid,_ := this.GetInt64("powerid")
	status,_ := this.GetInt("status")
	o := orm.NewOrm()
	power := models.Power{}
	power.Id = id
	o.Read(&power)
	if pid == 0 { //顶级
		power.Level = 2
		power.Pid = int(power.Id)
	}else{
		if power.Level == 2 {
			power.Level = 3
			power.Pid = int(power.Id)
		}
	}
	power.Controller = controller
	power.Action = action
	power.Status = status
	power.Powername = powername
	power.Id = powerid
	beego.Info(power)
	if _, err := o.Update(&power); err == nil {
		this.Rsp(true,"修改成功")
		return
	}
	this.Rsp(false,"修改失败")
	return
}