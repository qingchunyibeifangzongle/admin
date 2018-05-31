/**
 * User: yp
 * Time: 2018/5/28  上午10:44
*/

package common

import (
	"github.com/astaxie/beego"
	"admin/models"
)

type MainController struct {
	CommonController
}

type Tree struct {
	Id         int64      `json:"id"`
	Text       string     `json:"text"`
	IconCls    string     `json:"iconCls"`
	Checked    string     `json:"checked"`
	State      string     `json:"state"`
	Children   []Tree     `json:"children"`
	Attributes Attributes `json:"attributes"`
}
type Attributes struct {
	Url   string `json:"url"`
	Price int64  `json:"price"`
}

func (this *MainController) Admin(){
	this.TplName = this.GetTemplatetype() + "/login.html"
}
//首页
//session不存在，跳到网关登陆页面
func (this *MainController) Index() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	tree := this.GetTree()
	if this.IsAjax() {
		this.Data["json"] = &tree
		this.ServeJSON()
	} else {
		//groups := models.GroupList()
		this.Data["tree"] = &tree
		this.Data["userinfo"] = userinfo
		this.TplName = this.GetTemplatetype() + "/index.html"
	}
	
	
}

//登陆
func (this *MainController) Login()  {
	isajax := this.GetString("isajax")
	this.Data["json"] = isajax
	if isajax == "1" {
		username := this.GetString("username")
		password := this.GetString("password")
		user ,err := CheckLogin(username,password)
		if err == nil  {
			this.SetSession("userinfo",user)
			accessList, _ := GetAccessList(user.Id)
			this.SetSession("accesslist", accessList)
			this.Rsp(true, "登录成功")
			return
		} else {
			this.Rsp(false, err.Error())
			return
		}
	}
	userinfo := this.GetSession("userinfo")
	beego.Info(userinfo)
	if userinfo != nil {
		this.Ctx.Redirect(302, "/public/index")
	}
	this.TplName = this.GetTemplatetype() + "/login.html"
}

//退出
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.Ctx.Redirect(302,"/public/login")
}

//修改密码
func (this *MainController) Changepwd() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302,beego.AppConfig.String("rbac_auth_gateway"))
	}
	
	oldPwd := this.GetString("oldpassword")
	newPwd := this.GetString("newpassword")
	repeatPwd := this.GetString("repeatpassword")
	
	if newPwd != repeatPwd {
		this.Rsp(false,"两次输入的密码不一致！")
	}
	user , err := CheckLogin(userinfo.(models.User).Username,oldPwd)
	if err == nil {
		var u models.User
		u.Password = newPwd
		u.Id  = user.Id
		num , err := models.UpdPwd(&u)
		
		if err == nil && num > 0 {
			this.Rsp(true , "密码修改成功！")
			return
		} else {
			this.Rsp(false , err.Error())
			return
		}
	}
	this.Rsp(false , "密码有误！")
}



