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
	"strings"
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
	pageSize,_ := beego.AppConfig.Int("pagesize")
	
	page,_ := this.GetInt(":page")
	if page == 0 {
		page = 1
	}
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
	roleId,_ := this.GetInt64(":id")
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
	id ,_ := this.GetInt64("id")
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
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	roleid,_ := this.GetInt64(":id")
	powers1,powers2,powers3:= models.GetPowerlistByRoleId(roleid)
	//models.GetPowerlistByRoleId(roleid)
	parents,parent,children := models.GroupLists()
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["parents"] = &parents
	this.Data["parent"] = &parent
	this.Data["children"] = &children
	this.Data["powers1"] = &powers1
	this.Data["powers2"] = &powers2
	this.Data["powers3"] = &powers3
	this.TplName = this.GetTemplatetype() + "/rolepower.html"
}
func (this *RoleController) RolePowers(){
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	powerid := this.GetString("powerid")
	roleid,_ := this.GetInt64("id")
	ids := strings.Split(powerid,",")
	o := orm.NewOrm()
	//qb, _ := orm.NewQueryBuilder("mysql")
	rolepower := new(models.RolePower)
	
	o.QueryTable(rolepower).Filter("Role_id",roleid).Delete()
	for _,v := range ids{
		pid , _ := strconv.ParseInt(v, 10, 64) //转int64
		//qb.InsertInto("role_power","Role_id","Power_id").Values("?","?")
		sql := "INSERT INTO role_power ( Role_id, Power_id ) VALUES (?,?)"
		o.Raw(sql,roleid,pid).Exec()
	}
	this.Rsp(true,"修改成功")
	
}