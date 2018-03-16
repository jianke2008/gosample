package main

import (
	_ "waw99.com/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"waw99.com/models"
)

func init() {
	models.RegisterDB()

}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	beego.Run()
}

