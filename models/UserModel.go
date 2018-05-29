package models

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"errors"
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

//根据名字读取
func GetUserByName(username string) (user User)  {
	user = User{Username:username}
	o := orm.NewOrm()
	o.Read(&user,"Username")
	return user
}

//修改密码
func UpdPwd(u *User)(int64, error) {
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	if len(u.Password) > 0 {
		user["Password"] = Strtomd5(u.Password)
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}

	var table  User
	num,err := o.QueryTable(table).Filter("Id",u.Id).Update(user)
	return num ,err
}