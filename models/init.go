package models

import (
	"github.com/astaxie/beego/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"go_frame/conf"
	"strconv"
)
var (
	Orm orm.Ormer
)

func init(){
	//	注册业务数据表
	orm.RegisterModel()

	host, _ := beego.AppConfig.String("DB_HOST")
	port, _:= beego.AppConfig.String("DB_PORT")
	driver, _ := beego.AppConfig.String("Driver_Name")
	dbname, _ := beego.AppConfig.String("DB_NAME")
	user, _ := beego.AppConfig.String("DB_USERNAME")
	password, _ := beego.AppConfig.String("DB_PASSWORD")

	//	连接数据库
	orm.RegisterDriver(driver, orm.DRMySQL)
	orm.RegisterDataBase("default", driver, user+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8")
	debug, _ := strconv.ParseBool(conf.GetConf("debug"))
	orm.Debug = debug
	Orm = orm.NewOrm()
	Orm.Using(dbname)
}