/**
 * User: yp
 * Time: 2018/5/29  下午5:51
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	common.CommonController
}

func (this *UserController) Index ()  {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,"/public/login")
	}
	page , _ := this.GetInt64("page")
	pageSize ,_  := this.GetInt64("rows")
	sort := this.GetString("sort")
	order := this.GetString("order")
	
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "Id"
	}
	
	users , count := models.Getuserlist(page , pageSize ,sort)
	
	if this.IsAjax() {
		this.Data["json"] = &map[string]interface{}{"total":count , "rows": &users}
		this.ServeJSON()
		return
	} else  {
		tree := this.GetTree()
		this.Data["tree"] = &tree
		this.Data["users"] = &users
		
		this.TplName = this.GetTemplatetype() + "/userlist.html"
	}
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
