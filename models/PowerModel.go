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
	Controller       string         `orm:"size(100)" form:"Controller" valid:"Required;MaxSize(20);MinSize(6)"`
	Action           string         `orm:"size(100)" form:"Action"  valid:"Required"`
	Powername       string          `orm:"size(100)" form:"Powername" valid:"Required;MaxSize(20);MinSize(6)"`
	Pid             int             `form:"Pid"  valid:"Required"`
	Level           int             `orm:"default(2)" form:"Level" valid:"Range(1,2);"`
	Status          int             `orm:"default(2)" form:"Status" valid:"Range(1,2);"`
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



func GroupList() (parents []orm.Params,parent []orm.Params,children []orm.Params) {
	o := orm.NewOrm()
	power := new(Power)
	qs1 := o.QueryTable(power).Filter("Level",1)
	qs2 := o.QueryTable(power).Filter("Level",2)
	qs3 := o.QueryTable(power).Filter("Level",3)
	//var groups []orm.Params
	qs1.Values(&parents,"Id","Controller","Action","Powername","Pid","Level")
	qs2.Values(&parent,"Id","Controller","Action","Powername","Pid","Level")
	qs3.Values(&children,"Id","Controller","Action","Powername","Pid","Level")
	//beego.Info(parents)
	//beego.Info(parent)
	//beego.Info(children)
	//for _,v := range groups{
	//	if v["Pid"].(int64) == 0 && v["Level"].(int64) == 1 {
	//		maps = append(maps, v)
	//	}else if v["Level"].(int64) == 2 && v["Id"].(int64) == v["Pid"].(int64) {
	//		maps = append(maps, v)
	//	}else if  v["Level"].(int64) == 3 && v["Id"].(int64) == v["Pid"].(int64) {
	//		maps = append(maps, v)
	//	}
	//}
	return parents,parent,children
}

func GroupsList()(groups []orm.Params) {
	o := orm.NewOrm()
	power := new(Power)
	o.QueryTable(power).Values(&groups,"Id","Controller","Action","Powername","Pid","Level")
	return groups
}
func Getpowerlist(page int, pageSize int , sort string) (users []orm.Params , count int64){
	qb, _ := orm.NewQueryBuilder("mysql")
	
	//var offset int
	//if page <= 1 {
	//	offset = 0
	//} else {
	//	offset = (page - 1) * pageSize
	//}
	qb.Select("rolename","role.status","powername","role.id as role_id","power.id as power_id").From("role_power").LeftJoin("role").On("role.id = role_power.role_id").RightJoin("power").On("power.id = role_power.power_id").Where("role.status").In("1,2")
	//qs.Limit(pageSize , offset).OrderBy(sort).Values(&users)
	//count , _ = qs.Count()
	
	sql := qb.String()
	o := orm.NewOrm()
	beego.Info(sql)
	var maps []orm.Params
	num, _ := o.Raw(sql).Values(&maps)
	return maps,num
}

//func GetPowerlistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
//	o := orm.NewOrm()
//	power := new(Power)
//	count, _ = o.QueryTable(power).Filter(, Groupid).RelatedSel().Values(&nodes)
//	return nodes, 0
//}
//根据role去读取role_power，获取power数据
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
	//beego.Info(power1)
	//beego.Info(power2)
	//beego.Info(power3)
	
	
	//qb1, _ := orm.NewQueryBuilder("mysql")
	//qb2, _ := orm.NewQueryBuilder("mysql")
	//qb3, _ := orm.NewQueryBuilder("mysql")
	//
	//qb1.Select("power.Id","Controller","Action","Powername","Pid","power.Level").From("role_power").RightJoin("power").On("power.id = role_power.power_id").Where("role_id = ?" ).And("level = ?")
	//qb2.Select("power.Id","Controller","Action","Powername","Pid","Level").From("role_power").RightJoin("power").On("power.id = role_power.power_id").Where("role_id = ?" ).And("level = ?")
	//qb3.Select("power.Id","Controller","Action","Powername","Pid","Level").From("role_power").RightJoin("power").On("power.id = role_power.power_id").Where("role_id = ?" ).And("level = ?")
	//sql1 := qb1.String()
	//sql2 := qb2.String()
	//sql3 := qb3.String()
	//o.Raw(sql1,id,1).Values(&power1)
	//o.Raw(sql2,id,1).Values(&power2)
	//o.Raw(sql3,id,1).Values(&power3)
	return power1,power2,power3
}