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
)

type AccessNode struct {
	Id        int64
	Controller      string
	Childrens []*AccessNode
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