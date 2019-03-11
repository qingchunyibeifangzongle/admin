/**
 * User: yp
 * Time: 2018/5/25  下午2:36
*/

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
)

type LoginLog struct {
	Id          int64       `orm:"auto"`
	Username    string      `orm:"size(100)" form:"Username" valid:"Required;MaxSize(20);MinSize(1)"`
	Createtime  time.Time   `orm:"type(datetime);auto_now_add"`

}
func (u *LoginLog) TableName() string {
	return beego.AppConfig.String("rbac_login_log_table")
}
func init() {
	orm.RegisterModel(new(LoginLog))
}