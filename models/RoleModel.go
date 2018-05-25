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