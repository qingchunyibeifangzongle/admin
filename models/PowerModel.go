/**
 * User: yp
 * Time: 2018/5/25  下午2:20
*/

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Power struct {
	Id              int64           `orm:"auto"`
	//Power_id        int64
	Controller       string         `orm:"size(100)" form:"Controller" valid:"Required;MaxSize(20);MinSize(1)"`
	Action           string         `orm:"size(100)" form:"Action"  valid:"Required"`
	Powername       string          `orm:"size(100)" form:"Powername" valid:"Required;MaxSize(20);MinSize(1)"`
	Pid             int             `form:"Pid"  valid:"Required"`
	Level           int             `orm:"default(2)" form:"Level" valid:"Required"`
	Status          int             `orm:"default(2)" form:"Status" valid:"Required"`
	Createtime      time.Time       `orm:"type(datetime);auto_now_add"`
	Updatetime      time.Time       `orm:"null;type(datetime)" form:"-"`
}

func (u *Power) TableName() string {
	return beego.AppConfig.String("rbac_power_table")
}
func init() {
	orm.RegisterModel(new(Power))
}

func GetPowerTree(pid int64 , level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	power := new(Power)
	var powers []orm.Params
	_, err := o.QueryTable(power).Filter("Pid", pid).Filter("Level" , level).Values(&powers)
	if err != nil {
		return powers, err
	}
	return powers, nil
}


//role 分配权限 父节点
func GroupList() (parents []orm.Params,parent []orm.Params,children []orm.Params) {
	o := orm.NewOrm()
	power := new(Power)
	qs1 := o.QueryTable(power).Filter("Level",1)
	qs2 := o.QueryTable(power).Filter("Level",2)
	qs3 := o.QueryTable(power).Filter("Level",3)
	//var groups []orm.Params
	qs1.Values(&parents,"Id","Controller","Action","Powername","Pid","Level","Createtime","Status")
	qs2.Values(&parent,"Id","Controller","Action","Powername","Pid","Level","Createtime","Status")
	qs3.Values(&children,"Id","Controller","Action","Powername","Pid","Level","Createtime","Status")
	return parents,parent,children
}
func GroupLists() (parents []orm.Params,parent []orm.Params,children []orm.Params) {
	o := orm.NewOrm()
	power := new(Power)
	qs1 := o.QueryTable(power).Filter("Level",1).Filter("Status",2)
	qs2 := o.QueryTable(power).Filter("Level",2).Filter("Status",2)
	qs3 := o.QueryTable(power).Filter("Level",3).Filter("Status",2)
	//var groups []orm.Params
	qs1.Values(&parents,"Id","Controller","Action","Powername","Pid","Level","Createtime","Status")
	qs2.Values(&parent,"Id","Controller","Action","Powername","Pid","Level","Createtime","Status")
	qs3.Values(&children,"Id","Controller","Action","Powername","Pid","Level","Createtime","Status")
	return parents,parent,children
}

func GroupsList()(groups []orm.Params) {
	o := orm.NewOrm()
	power := new(Power)
	o.QueryTable(power).Values(&groups,"Id","Controller","Action","Powername","Pid","Level")
	return groups
}
func Getpowerlist(page int, pageSize int) (powers []orm.Params , count int64){
	qb, _ := orm.NewQueryBuilder("mysql")
	var offset int
	if page <= 1 {
		page = 0
	} else {
		offset = (page - 1) * pageSize
	}
	o  := orm.NewOrm()
	count, err := o.QueryTable(new(Power)).Count()
	beego.Info(err)
	
	qb.Select("Id","Controller","Action","Level","Pid","Status","Createtime","Updatetime","Powername").From("power").Where("Status").In("1,2").Limit(pageSize).Offset(offset)
	
	sql := qb.String()
	o.Raw(sql).Values(&powers)
	return powers , count
}

//根据role去读取role_power，获取power数据
//分配权限是否选中
func GetPowerlistByRoleId(id int64) (power1 []orm.Params,power2 []orm.Params,power3 []orm.Params) {
	o := orm.NewOrm()
	var powerIds []orm.Params
	var powerId1 []orm.Params
	var powerId2 []orm.Params
	var powerId3 []orm.Params
	o.QueryTable(new(RolePower)).Filter("Role_id",id).Values(&powerIds,"power_id")
	for _,v := range powerIds {
		o.QueryTable(new(Power)).Filter("id",v["Power_id"]).Filter("level",1).Values(&powerId1,"Id","Controller","Action","Pid","Level","Powername")
		o.QueryTable(new(Power)).Filter("id",v["Power_id"]).Filter("level",2).Values(&powerId2,"Id","Controller","Action","Pid","Level","Powername")
		o.QueryTable(new(Power)).Filter("id",v["Power_id"]).Filter("level",3).Values(&powerId3,"Id","Controller","Action","Pid","Level","Powername")
		for _,n := range powerId1 {
			power1 = append(power1, n)
		}
		for _,n := range powerId2 {
			power2 = append(power2, n)
		}
		for _,n := range powerId3 {
			power3 = append(power3, n)
		}
	}
	return power1,power2,power3
}

func GetPowerId(id int64) (powers []orm.Params) {
	o := orm.NewOrm()
	o.QueryTable(new(Power)).Filter("Id",id).Values(&powers)
	return powers
}