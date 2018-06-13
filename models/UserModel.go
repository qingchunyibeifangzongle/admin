package models

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"errors"
)

type User struct {
	Id       int64          `orm:"auto"`
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(1)"`
	Password      string    `orm:"size(32)" form:"Password"  valid:"Required;MaxSize(20);MinSize(1)"`
	Nickname      string    `orm:"-" form:"Nickname" valid:"Required;MaxSize(20);MinSize(1)"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Required;"`
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

func Getuserlist(page int, pageSize int) (users []orm.Params , count int64){
	//o := orm.NewOrm()
	//user := new(User)
	//qs := o.QueryTable(user)
	//var offset int
	//if page <= 1 {
	//	offset = 0
	//} else {
	//	offset = (page - 1) * pageSize
	//}
	//
	//qs.Limit(pageSize , offset).OrderBy(sort).Values(&users)
	//count , _ = qs.Count()
	//return users , count
	qb, _ := orm.NewQueryBuilder("mysql")
	var offset int
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	o  := orm.NewOrm()
	count, err := o.QueryTable(new(UserRole)).Count()
	beego.Info(err)
	qb.Select("user.id","password","email","user.createtime","user.updatetime","user.status","username","role.id as role_id","remark","role.status as role_status","rolename").From("user_role").LeftJoin("user").On("user.id = user_role.user_id").RightJoin("role").On("role.id = user_role.role_id").Where("user.status").In("1,2").Limit(pageSize).Offset(offset)

	sql := qb.String()
	beego.Info(sql)
	o.Raw(sql).Values(&users)
	return users,count
}

//修改用户信息
func UpdUser(u *User) (int64 ,error) {
	//if err := checkUser(u); err != nil {
	//	return 0, err
	//}
	o := orm.NewOrm()
	
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Password) > 0 {
		user["Password"] = Strtomd5(u.Password)
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	user["Updatetime"] = time.Now()
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
	
}

