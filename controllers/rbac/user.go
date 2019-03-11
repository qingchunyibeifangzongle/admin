/**
 * User: yp
 * Time: 2018/5/29  下午5:51
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
<<<<<<< HEAD
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"strconv"
=======
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"github.com/astaxie/beego/validation"
>>>>>>> 416f0cd1e9133e4ad6ab2023f135e5fd3f1ff301
)

type UserController struct {
	common.CommonController
}

func (this *UserController) Index ()  {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	pageSize,_ := beego.AppConfig.Int("pagesize")
	page, _ := this.GetInt(":page")
	if page == 0 {
		page = 1
	}
<<<<<<< HEAD
=======
 
>>>>>>> 416f0cd1e9133e4ad6ab2023f135e5fd3f1ff301
	users , totalRows := models.Getuserlist(page , pageSize)
	res := models.Paginator(page, pageSize, totalRows)
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["users"] = &users
	this.Data["paginator"] = res
	this.Data["totals"] = totalRows
	this.Data["username"] = userinfo.(models.User).Username
	this.TplName = this.GetTemplatetype() + "/userlist.html"

}

//user add
//user redirect
func (this *UserController) UserAdd() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	roles := models.GetRoleAll()
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["roles"] = &roles
	this.Data["username"] = userinfo.(models.User).Username
	this.TplName = this.GetTemplatetype() + "/useradd.html"
}

// user add  now
func (this *UserController) UserAdds() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	username := this.GetString("username")
	nickname := this.GetString("nickname")
	password := this.GetString("password")
	email := this.GetString("email")
	status,_ := this.GetInt("status")
	roleid,_ := this.GetInt64("roleid")
	
	o := orm.NewOrm()
	user := new(models.User)
	user.Username = username
	user.Email = email
	user.Nickname = nickname
	user.Status = status
	user.Password = password
	valid := validation.Validation{}
	valid.Valid(user)
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		this.Rsp(false,valid.Errors[0].Key +"    "+ valid.Errors[0].Message)
		return
	}
	user.Password = models.Pwdhash(password)
	id,err := o.Insert(user)
	if err != nil {
		this.Rsp(false,"用户添加失败！")
	}
	o.Insert(&models.UserRole{Role_id:roleid,User_id:id})
	this.Rsp(true,"用户添加成功！")
}

// user edit redirect
func (this *UserController) UserEdit() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	roles := models.GetRoleAll()
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["roles"] = &roles
	this.Data["username"] = userinfo.(models.User).Username
	this.Data["email"] = userinfo.(models.User).Email
	this.Data["id"] = userinfo.(models.User).Id
	this.TplName = this.GetTemplatetype() + "/useredit.html"
}

// user edit now
func (this *UserController) UserEdits() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	email := this.GetString("email")
	username := this.GetString("username")
	role_id := this.GetString("rolename")

	id := this.GetString("id")
	intid ,_ := strconv.ParseInt(id, 10, 64)
	introleid ,_ := strconv.ParseInt(role_id, 10, 64)
	o := orm.NewOrm()
	err := o.Begin()
	u := models.User{Id:intid,Username:username,Email:email}
	num, err := models.UpdUser(&u)
	
	r := models.UserRole{User_id:intid,Role_id:introleid}
	if err == nil && num > 0 {
		models.UpdUSerRole(&r)
		user := o.Read(&u)
		this.SetSession("userinfo",user)
		err = o.Commit()
		this.Rsp(true, "Success")
		return
	} else {
		err = o.Rollback()
		this.Rsp(false, err.Error())
	}
}

//开关
func (this *UserController) UserSwitch(){
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	status,_ := this.GetInt("status")
	id,_ := this.GetInt64("id")
	o := orm.NewOrm()
	user := models.User{Id:id,Status:status}
	if _, err := o.Update(&user,"Status"); err == nil {
		this.Rsp(true,"修改成功")
		return
	}
	this.Rsp(false,"修改失败")
	return
}