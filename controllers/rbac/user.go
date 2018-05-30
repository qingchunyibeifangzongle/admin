/**
 * User: yp
 * Time: 2018/5/29  下午5:51
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
)

type UserController struct {
	common.CommonController
}

func (this *UserController) Index ()  {
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
		
		if this.GetTemplatetype() != "easyui" {
			this.Layout = this.GetTemplatetype() + "/public/layout.tpl"
		}
		this.TplName = this.GetTemplatetype() + "/rbac/user.tpl"
	}
}
