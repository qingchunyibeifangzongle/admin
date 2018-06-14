/**
 * User: yp
 * Time: 2018/5/28  上午10:44
*/

package common

import (
	"admin/models"
	"errors"
	"strings"
	"fmt"
	"github.com/astaxie/beego/context"
	"strconv"
	"github.com/astaxie/beego"
)

type AccessNode struct {
	Id        int64
	Controller      string
	Childrens []*AccessNode
}
func AccessRegister(){
	var Check  = func(ctx *context.Context) {
		user_auth_type,_ := strconv.Atoi(beego.AppConfig.String("user_auth_type"))
		rbac_auth_gateway := beego.AppConfig.String("rbac_auth_gateway")
		var accesslist map[string]bool
		if user_auth_type > 0 {
			params := strings.Split(strings.ToLower(strings.Split(ctx.Request.RequestURI,"?")[0]),"/")
			if CheckAccess(params){
				userinfo := ctx.Input.Session("userinfo")
				if userinfo == nil {
					ctx.Redirect(302, rbac_auth_gateway)
					return
				}
				adminuser := beego.AppConfig.String("rbac_admin_user")
				if userinfo.(models.User).Username == adminuser {
					return
				}
				if user_auth_type == 1 {
					listbysession := ctx.Input.Session("accesslist")
					if listbysession != nil {
						accesslist = listbysession.(map[string]bool)
					}
				} else if user_auth_type == 2 {
					
					accesslist, _ = GetAccessList(userinfo.(models.User).Id)
				}
				beego.Info(accesslist)
				ret := AccessDecision(params, accesslist)
				beego.Info(ret)
				if !ret {
					ctx.Output.JSON(&map[string]interface{}{"status": false, "info": "权限不足"}, true, false)
				}
			}
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, Check)
}
//Determine whether need to verify
func CheckAccess(params []string) bool {
	if len(params) < 3 {
		return false
	}
	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
		if params[1] == nap {
			return false
		}
	}
	return true
}

//To test whether permissions
func AccessDecision(params []string, accesslist map[string]bool) bool {
	if CheckAccess(params) {
		s := fmt.Sprintf("%s/%s/%s", params[1], params[2], params[3])
		if len(accesslist) < 1 {
			return false
		}
		_, ok := accesslist[s]
		if ok != false {
			return true
		}
	} else {
		return true
	}
	return false
}

func GetAccessList(id int64)(map[string]bool ,error){
	list,err := models.Accesslist(id)
	if err != nil {
		return nil,err
	}
	alist := make([]*AccessNode, 0)
	for _, l := range list {
		//beego.Info(l["Pid"].(int64)==0)
		//beego.Info(l["Id"].(int64))
		//beego.Info(l["Controller"].(string))
		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 {
			anode := new(AccessNode)
			anode.Id = l["Id"].(int64)
			anode.Controller = l["Controller"].(string)
			alist = append(alist, anode)
		}
	}

	for _, l := range list {
		if l["Level"].(int64) == 2 {
			for _, an := range alist {
				if an.Id == l["Pid"].(int64) {
					anode := new(AccessNode)
					anode.Id = l["Id"].(int64)
					anode.Controller = l["Controller"].(string)
					an.Childrens = append(an.Childrens, anode)
				}
			}
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 3 {
			for _, an := range alist {
				for _, an1 := range an.Childrens {
					if an1.Id == l["Pid"].(int64) {
						anode := new(AccessNode)
						anode.Id = l["Id"].(int64)
						anode.Controller = l["Controller"].(string)
						an1.Childrens = append(an1.Childrens, anode)
					}
				}
				
			}
		}
	}
	
	accesslist := make(map[string]bool)
	for _, v := range alist {
		for _, v1 := range v.Childrens {
			for _, v2 := range v1.Childrens {
				vname := strings.Split(v.Controller, "/")
				v1name := strings.Split(v1.Controller, "/")
				v2name := strings.Split(v2.Controller, "/")
				str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]), strings.ToLower(v2name[0]))
				accesslist[str] = true
			}
		}
	}
	return accesslist, nil
	
}



//检测用户名密码是否正确
func CheckLogin(username string , password string) (user models.User ,err error) {
	user = models.GetUserByName(username)
	if user.Id == 0 {
		return user,errors.New("用户不存在")
	}
	if user.Password != models.Pwdhash(password) {
		return user,errors.New("密码不正确")
	}
	return user ,nil
}