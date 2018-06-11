/**
 * User: yp
 * Time: 2018/6/11  下午4:58
*/

package rbac

import (
	"admin/controllers/common"
	"admin/models"
	"github.com/astaxie/beego"
)

type PowerController struct {
	common.CommonController
}

func (this *PowerController) Index() {
	//userinfo := this.GetSession("userinfo")
	//if userinfo == nil {
	//	this.Ctx.Redirect(302,"/public/login")
	//}
	//const pageSize = 1
	//page,_ := this.GetInt(":page")
	//if page == 0 {
	//	page = 1
	//}
	//powers , totalRows := models.Getpowerlist(page , pageSize)
	//res := models.Paginator(page, pageSize, totalRows)
	powers1,powers2,powers3 := models.GroupList()
	beego.Info(powers1)
	tree := this.GetTree()
	this.Data["tree"] = &tree
	this.Data["powers1"] = &powers1
	this.Data["powers2"] = &powers2
	this.Data["powers3"] = &powers3
	//this.Data["paginator"] = res
	//this.Data["totals"] = totalRows
	this.TplName = this.GetTemplatetype() + "/powerlist.html"
}