package user_db

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"

)


func init(){
	//connect user info db
	err:=orm.RegisterDriver("mysql",orm.DRMySQL)
	if err!=nil{
		beego.Error(err.Error())
		return
	}
	username:=beego.AppConfig.String("usrdbuser")
	pwd:=beego.AppConfig.String("usrdbpwd")
	addr:=beego.AppConfig.String("usrdbhost")
	dbname:=beego.AppConfig.String("usrdbname")
	conn:=username+":"+pwd+"@tcp(" +addr + ")/" + dbname + "?charset=utf8"

	err=orm.RegisterDataBase("default","mysql",conn)
	if err!=nil{
		beego.Error(err.Error())
	}
}
