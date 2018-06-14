package routers

import (
	"github.com/astaxie/beego"
	"admin/controllers/common"
	//"admin/controllers"
	"admin/controllers/rbac"
)

func init() {
	
	//beego.Router("/",&controllers.MainController{})
	beego.Router("/",&common.MainController{},"*:Admin")
	beego.Router("public/index",&common.MainController{},"*:Index")
	beego.Router("public/login",&common.MainController{},"*:Login")
	beego.Router("public/logout",&common.MainController{},"*:Logout")
	beego.Router("public/changepwd",&common.MainController{},"*:Changepwd")
	
	
	
	beego.Router("admin/user/?:page",&rbac.UserController{},"*:Index")
	beego.Router("admin/user/add",&rbac.UserController{},"*:UserAdd")
	beego.Router("admin/user/adds",&rbac.UserController{},"*:UserAdds")
	beego.Router("admin/user/edit",&rbac.UserController{},"*:UserEdit")
	beego.Router("admin/user/edits",&rbac.UserController{},"*:UserEdits")
	beego.Router("admin/user/switch",&rbac.UserController{},"*:UserSwitch")
	
	
	beego.Router("admin/role/?:page",&rbac.RoleController{},"*:Index")
	veego.Router("admin/role/edit/:id",&rbac.RoleController{},"*:RoleEdit")
	beego.Router("admin/role/edits",&rbac.RoleController{},"*:RoleEdits")
	beego.Router("admin/role/add",&rbac.RoleController{},"*:RoleAdd")
	beego.Router("admin/role/adds",&rbac.RoleController{},"*:RoleAdds")
	beego.Router("admin/role/rolepower/:id",&rbac.RoleController{},"*:RolePower")
	beego.Router("admin/role/rolepowers",&rbac.RoleController{},"*:RolePowers")
	beego.Router("admin/role/switch",&rbac.RoleController{},"*:RoleSwitch")
	
	
	beego.Router("admin/power/?:page",&rbac.PowerController{},"*:Index")
	beego.Router("admin/power/edit/:id",&rbac.PowerController{},"*:PowerEdit")
	beego.Router("admin/power/edits",&rbac.PowerController{},"*:PowerEdits")
	beego.Router("admin/power/add",&rbac.PowerController{},"*:PowerAdd")
	beego.Router("admin/power/adds",&rbac.PowerController{},"*:PowerAdds")
	//开关
	beego.Router("admin/power/switch",&rbac.PowerController{},"*:PowerSwitch")
}

