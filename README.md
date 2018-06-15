# admin
beego admin ，权限rbac
自动生成表

#app.conf

appname = admin
httpAddr = 127.0.0.1
httpport = 8080
runmode = dev

#mysql

db_host = localhost
db_port = 3306
db_user = root
db_pass = root
db_name = beego_admin
db_type = mysql

#session
sessionon = true

#table_names
rbac_user_table = user
rbac_role_table = role
rbac_power_table = power
rbac_role_power_table = role_power
rbac_user_role_table = user_role
rbac_login_log_table = login_log
#admin用户名 此用户登录不用认证
rbac_admin_user = admin

#默认不需要认证模块
not_auth_package = public,static
#默认认证类型 0 不认证 1 登录认证 2 实时认证
user_auth_type = 1
#默认登录网关
rbac_auth_gateway = /public/login
#默认模版
template_type=amz


#配置文件，仿照beego admin写的。里面的代码自动生成表仿照的