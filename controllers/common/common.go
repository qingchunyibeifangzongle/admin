/**
 * User: yp
 * Time: 2018/5/28  上午10:42
*/

package common

import (
	"github.com/astaxie/beego"
	"admin/models"
)

type CommonController struct {
	beego.Controller
	Templatetype string
}

func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
}
func (this *CommonController) GetTree() []Tree {
	powers, _ := models.GetPowerTree(0,1)
	tree := make([]Tree , len(powers))
	
	for k,v := range powers {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Powername"].(string)
		
		children, _ := models.GetPowerTree(v["Id"].(int64), 2)
		tree[k].Children = make([]Tree , len(children))
		for k1,v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Powername"].(string)
			tree[k].Children[k1].Attributes.Url = "/" + v["Controller"].(string)+ "/" + v1["Action"].(string)
		}
	}
	return tree
	
	
	
	
}

func (this *CommonController) GetTemplatetype() string {
	templatetype := beego.AppConfig.String("template_type")
	if templatetype == "" {
		templatetype = "amz"
	}
	return templatetype
}
