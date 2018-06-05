/**
 * User: yp
 * Time: 2018/5/25  下午2:36
*/

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserRole struct {
	Id          int64     `orm:"auto"`
	User_id     int64
	Role_id     int64
	//Role        []*Role   `orm:"rel(m2m)"`
	//User        []*User   `orm:"rel(m2m)"`
}
func (u *UserRole) TableName() string {
	return beego.AppConfig.String("rbac_user_role_table")
}
func init() {
	orm.RegisterModel(new(UserRole))
}

func UpdUSerRole(r *UserRole) (int64 ,error){
	o := orm.NewOrm()
	userRole := new(UserRole)
	num, err := o.QueryTable(userRole).Filter("user_id",r.User_id).Update(orm.Params{"role_id":r.Role_id})
	return num,err
}