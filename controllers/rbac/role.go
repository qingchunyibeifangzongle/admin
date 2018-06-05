/**
 * User: yp
 * Time: 2018/6/4  下午3:19
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego"
	"strconv"
)

type RoleController struct {
	common.CommonController
}

func (this *RoleController) Index() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	const pageSize = 1
	var page int
	var sort string
	sort = "Id"
	roles , totalRows := models.Getrolelist(page , pageSize ,sort)
	res := models.Paginator(page, pageSize, totalRows)
	beego.Info(roles)
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
	beego.Info(roles)
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
	beego.Info(num)
	if num > 0 && err == nil {
		this.Rsp(true, "Success")
		return
	}else {
		this.Rsp(false, err.Error())
		return
	}
}

func (this *RoleController) RoleAdd() {

}


func (this *RoleController) RoleAdds() {

}