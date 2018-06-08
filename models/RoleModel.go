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
	Remark          string      `orm:"size(120)" form:"Remark" valid:"Required;Match(/^[\u4e00-\u9fa5]+$/);MaxSize(20);MinSize(1)"`
	//Rule            string      `orm:"size(120)" form:"Rule"`
	Status          int         `orm:"default(2)" form:"Status" valid:"Range(1,2);"`
	Rolename        string      `orm:"size(32)" form:"Rolename" valid:"Required;MaxSize(20);MinSize(1);Match(/^[\u4e00-\u9fa5]+$/)"`
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

//func Getrolelist(page int, pageSize int , sort string) (users []orm.Params , count int64){
//	qb, _ := orm.NewQueryBuilder("mysql")
//
//	//var offset int
//	//if page <= 1 {
//	//	offset = 0
//	//} else {
//	//	offset = (page - 1) * pageSize
//	//}
//	qb.Select("rolename","role.status","powername","role.id as role_id","power.id as power_id").From("role_power").LeftJoin("role").On("role.id = role_power.role_id").RightJoin("power").On("power.id = role_power.power_id").Where("role.status").In("1,2")
//	//qs.Limit(pageSize , offset).OrderBy(sort).Values(&users)
//	//count , _ = qs.Count()
//
//	sql := qb.String()
//	o := orm.NewOrm()
//	beego.Info(sql)
//	var maps []orm.Params
//	num, _ := o.Raw(sql).Values(&maps)
//	return maps,num
//}

func Getrolelist(page int, pageSize int) (roles []orm.Params , count int64){
	qb, _ := orm.NewQueryBuilder("mysql")
	var offset int
	if page <= 1 {
		page = 0
	} else {
		offset = (page - 1) * pageSize
	}
	o  := orm.NewOrm()
	count, err := o.QueryTable(new(Role)).Count()
	beego.Info(err)
	
	qb.Select("Id","Remark","Rolename","Createtime","Updatetime","Status").From("role").Where("Status").In("1,2").Limit(pageSize).Offset(offset)
	
	sql := qb.String()
	o.Raw(sql).Values(&roles)
	return roles , count
}

func GetRoleAll()(roles []orm.Params) {
	o := orm.NewOrm()
	role := new(Role)
	o.QueryTable(role).Values(&roles)
	return roles
}


func GetRoleId(roleId int64) (role Role ) {
	role = Role{Id:roleId}
	o := orm.NewOrm()
	o.Read(&role,"Id")
	return role
}

func UpdRole(role *Role) (int64 , error) {
	o := orm.NewOrm()
	updatetime := time.Now()
	role.Updatetime = updatetime
	num , err := o.Update(role,"Id","Remark","Rolename","Updatetime","Status")
	return  num,err
}