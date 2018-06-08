/**
 * User: yp
 * Time: 2018/6/4  下午3:19
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"reflect"
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
	roleId,_ := this.GetInt(":id")
	//var roleId int
	//this.Ctx.Input.Bind(&roleId,":id")
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
	id ,_ := this.GetInt("id")
	rolename := this.GetString("rolename")
	remark := this.GetString("remark")
	status,_ := this.GetInt("status")
	role := models.Role{Id:id,Remark:remark,Rolename:rolename,Status:status}
	valid := validation.Validation{}
	valid.Valid(role)
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		this.Rsp(false,valid.Errors[0].Key +"    "+ valid.Errors[0].Message)
		return
	}
	num,err := models.UpdRole(&role)
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


// role add
func (this *RoleController) RoleAdds() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	remark := this.GetString("remark")
	rolename := this.GetString("rolename")
	status,_ := this.GetInt("status")
	valid := validation.Validation{}
	valid.Valid(models.Role{Remark:remark,Rolename:rolename,Status:status})
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

func (this *RoleController) RolePower() {
	//userinfo := this.GetSession("userinfo")
	//if userinfo == nil {
	//	this.Ctx.Redirect(302,"/public/login")
	//}
	//roleid,_ := this.GetInt64(":id")
	
	//parents,parent,children := models.GroupList()
	nodes := models.GroupsList()
	
	for i := 0; i < len(nodes); i++ {
		if nodes[i]["Pid"] == int64(0) {
			nodes[i]["ParentsId"] = nodes[i]["Id"]
		} else if  nodes[i]["Pid"] == int64(1)  {
			nodes[i]["ParentId"] = nodes[i]["Id"]
			nodes[i]["ParentState"] = "closed"
		}else {
			nodes[i]["ChildrenId"] = nodes[i]["Id"]
			nodes[i]["ChildrenState"] = "closed"
		}
		
		beego.Info(reflect.TypeOf(nodes))
		beego.Info(reflect.TypeOf(nodes[i]["Level"]))
	}
	//this.Data["json"] = &map[string]interface{}{"total": 1, "rows": &nodes}
	//this.ServeJSON()
	tree := this.GetTree()
	this.Data["tree"] = &tree
	//this.Data["parents"] = &parents
	//this.Data["parent"] = &parent
	//this.Data["children"] = &children
	this.Data["groups"] = &nodes
	this.TplName = this.GetTemplatetype() + "/a.html"
}