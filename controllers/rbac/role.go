/**
 * User: yp
 * Time: 2018/6/4  下午3:19
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type RoleController struct {
	common.CommonController
}

func (this *RoleController) Index() {
	//userinfo := this.GetSession("userinfo")
	//if userinfo == nil {
	//	this.Ctx.Redirect(302,"/public/login")
	//}
	const pageSize = 1
	page,_ := this.GetInt(":page")
	if page == 0 {
		page = 1
	}
	beego.Info(page)
	roles , totalRows := models.Getrolelist(page , pageSize)
	res := models.Paginator(page, pageSize, totalRows)
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["roles"] = &roles
	this.Data["paginator"] = res
	this.Data["totals"] = totalRows
	this.TplName = this.GetTemplatetype() + "/rolelist.html"
}

// role 管理
func (this *RoleController) RoleEdit() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	var roleId int64
	this.Ctx.Input.Bind(&roleId,":id")
	roles := models.GetRoleId(roleId)
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["Rolename"] = roles.Rolename
	this.Data["Status"] = roles.Status
	this.Data["Id"] = roles.Id
	this.Data["Remark"] = roles.Remark
	this.TplName = this.GetTemplatetype() + "/roleedit.html"
}

func (this *RoleController) RoleEdits() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	id := this.GetString("id")
	rolename := this.GetString("rolename")
	remark := this.GetString("remark")
	status := this.GetString("status")
	statuss,_:=strconv.Atoi(status)
	
	roleId ,_ := strconv.ParseInt(id, 10, 64)
	num,err := models.UpdRole(&models.Role{Id:roleId,Remark:remark,Rolename:rolename,Status:statuss})
	if num > 0 && err == nil {
		this.Rsp(true, "Success")
		return
	}else {
		this.Rsp(false, err.Error())
		return
	}
}

func (this *RoleController) RoleAdd() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	
	tree := this.GetTree()
	this.Data["tree"] = &tree
	//this.Data["users"] = userinfo
	this.TplName = this.GetTemplatetype() + "/roleadd.html"
}


func (this *RoleController) RoleAdds() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	remark := this.GetString("remark")
	rolename := this.GetString("rolename")
	status,_ := this.GetInt("status")
	valid := validation.Validation{}
	valid.Required(rolename,"角色名称")
	valid.Required(remark,"角色描述")
	valid.Required(status,"状态")
	valid.MaxSize(rolename,20,"角色名称")
	valid.MaxSize(remark,20,"角色描述")
	valid.MinSize(remark,1,"角色描述")
	valid.MinSize(rolename,1,"角色名称")
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		this.Rsp(false,valid.Errors[0].Key +"    "+ valid.Errors[0].Message)
		return
	}
	r := models.Role{Remark:remark,Status:status,Rolename:rolename}
	
	o := orm.NewOrm()
	
	id, err := o.Insert(&r)
	if err == nil && id >0 {
		this.Rsp(true ,"添加成功")
		return
	}else {
		this.Rsp(false,err.Error())
		return
	}
}