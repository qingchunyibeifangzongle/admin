/**
 * User: yp
 * Time: 2018/5/29  下午5:51
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"reflect"
)

type UserController struct {
	common.CommonController
}

func (this *UserController) Index ()  {
	
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	//var pagination []orm.Params
	
	const pageSize = 1
	var page int
	//var offset int
	this.Ctx.Input.Bind(&page,"page")
 	////pageSize ,_  := this.GetInt64("rows")
	//sort := this.GetString("sort")
	//order := this.GetString("order")
	////
	//if len(order) > 0 {
	//	if order == "desc" {
	//		sort = "-" + sort
	//	}
	//} else {
	//	sort = "Id"
	//}
	//
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
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["username"] = userinfo.(models.User).Username
	this.Data["email"] = userinfo.(models.User).Email
	this.Data["id"] = userinfo.(models.User).Id
	
	this.TplName = this.GetTemplatetype() + "/useredit.html"
}

// user edit now
func (this *UserController) UserEdits() {
	email := this.GetString("email")
	username := this.GetString("username")
	id := this.GetString("id")
	intid ,_ := strconv.ParseInt(id, 10, 64)
	u := models.User{Id:intid,Username:username,Email:email}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	beego.Info(u)
	//id, err := models.UpdUser(&u)
	//if err == nil && id >0 {
	//	this.Rsp(true, "Success")
	//	return
	//} else {
	//	this.Rsp(false, err.Error())
	//}
}