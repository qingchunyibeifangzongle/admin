package models

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64          `orm:"auto"`
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
	Password      string    `orm:"size(32)" form:"Password"  valid:"Required;MaxSize(20);MinSize(6)"`
	Nickname      string    `orm:"-" form:"Nickname" valid:"Required;MaxSize(20);MinSize(6)"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2);"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
	Createtime    time.Time `orm:"type(datetime);auto_now_add" `
	Updatetime    time.Time `orm:"null;type(datetime)" form:"-"`
	//Role          []*Role   `orm:"rel(m2m)"`
}

func (u *User) TableName() string {
	return beego.AppConfig.String("rbac_user_table")
}
func init() {
	orm.RegisterModel(new(User))
}
