/**
 * User: yp
 * Time: 2018/5/25  下午2:36
*/

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type RolePower struct {
	Id          int           `orm:"auto"`
	Role_id     int
	Power_id    int
	//Role        []*Role        `orm:"rel(m2m)"`
	//Power       []*Power       `orm:"rel(m2m)"`
	
}
func (u *RolePower) TableName() string {
	return beego.AppConfig.String("rbac_role_power_table")
}
func init() {
	orm.RegisterModel(new(RolePower))
}