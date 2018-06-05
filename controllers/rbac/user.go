/**
 * User: yp
 * Time: 2018/5/29  下午5:51
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type UserController struct {
	common.CommonController
}

func (this *UserController) Index ()  {
	
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	const pageSize = 1
	var page int
	//var offset int
	this.Ctx.Input.Bind(&page,"page")
 
	var sort string
	sort = "Id"
	users , totalRows := models.Getuserlist(page , pageSize ,sort)
	res := models.Paginator(page, pageSize, totalRows)
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["users"] = &users
	this.Data["paginator"] = res
	this.Data["totals"] = totalRows
	this.TplName = this.GetTemplatetype() + "/userlist.html"

}

//user add
//user redirect
func (this *UserController) UserAdd() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["users"] = userinfo
	this.TplName = this.GetTemplatetype() + "/useradd.html"
}

// user add  now
func (this *UserController) UserAdds() {
	username := this.GetString("username")
	password := this.GetString("password")
	email := this.GetString("email")
	
	o := orm.NewOrm()
	user := new(models.User)
	user.Username = username
	user.Password = models.Pwdhash(password)
	user.Email = email
	num,_ := o.Insert(user)
	if num == 0 {
		this.Rsp(false,"用户添加失败！")
	}
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
	//this.Data["password"] = userinfo.(models.User).Password
	
	this.TplName = this.GetTemplatetype() + "/useredit.html"
}

// user edit now
func (this *UserController) UserEdits() {
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