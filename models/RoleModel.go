/**
 * User: yp
 * Time: 2018/5/25  上午11:48
*/

package models


import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Role struct {
	Id              int64       `orm:"auto"`
	//Role_id        int64
	Rolename      string        `orm:"size(32)" form:"Rolename" valid:"Required;MaxSize(20);MinSize(6)"`
	Createtime     time.Time    `orm:"type(datetime);auto_now_add"`
	Updatetime     time.Time    `orm:"null;type(datetime)" form:"-"`
	//User           []*User      `orm:"rel(m2m)"`
}

func (u *Role) TableName() string {
	return beego.AppConfig.String("rbac_role_table")
}
func init() {
	orm.RegisterModel(new(Role))
}

func Accesslist(id int64) (list []orm.Params, err error) {
	var users []orm.Params
	o := orm.NewOrm()
	userRole := new(UserRole)
	_, err = o.QueryTable(userRole).Filter("User_id",id).Values(&users)
	if err != nil {
		return nil,err
	}
	var role_id interface{}
	for _,v := range users {
		role_id  = v["Role_id"]
	}
	var roles []orm.Params
	role := new(Role)
	_,err = o.QueryTable(role).Filter("id",role_id).Values(&roles)
	if err != nil {
		return nil,err
	}
	
	var powerId []orm.Params
	role_power := new(RolePower)
	for _,v := range roles {
		_,err = o.QueryTable(role_power).Filter("Role_id",v["Id"]).Values(&powerId)
	}
	
	if err != nil {
		return nil,err
	}
	
	
	//这个里面的roles是多维数组，有n个power——id，so raqnge去select
	var powers []orm.Params
	power := new(Power)
	for _,v := range powerId {
		_,err = o.QueryTable(power).Filter("id",v["Power_id"]).Values(&powers)
		if err != nil  {
			return nil,err
		}
		
		for _,n := range powers {
			list = append(list, n)
			
		}
	}
	
	return list,nil
}