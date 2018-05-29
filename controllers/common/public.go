/**
 * User: yp
 * Time: 2018/5/28  上午10:44
*/

package common

import "github.com/astaxie/beego"

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
	this.TplName = this.GetTemplatetype() + "/public/login.tpl"
}
//首页
//session不存在，跳到网关登陆页面
func (this *MainController) Index() {
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
	}
	this.TplName = this.GetTemplatetype() + "/public/index.tpl"
	//tree := this.GetTree()
	
	
	
}

//登陆
func (this *MainController) Login()  {
	isajax := this.GetString("isajax")
	if isajax == "1" {
		username := this.GetString("username")
		password := this.GetString("password")
		user ,err := CheckLogin(username,password)
		beego.Debug(user)
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
	if userinfo != nil {
		this.Ctx.Redirect(302, "/public/index")
	}
	this.TplName = this.GetTemplatetype() + "/public/login.tpl"
}



