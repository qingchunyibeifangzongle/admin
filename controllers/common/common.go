/**
 * User: yp
 * Time: 2018/5/28  上午10:42
*/

package common

import (
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
	Templatetype string
}

func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}
/*
func (this *CommonController) GetTree(pid int64) []Tree {
	power, _ := m.GetPowerTree(pid)
}*/

func (this *CommonController) GetTemplatetype() string {
	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "easyui"
	}
	return templatetype
}
