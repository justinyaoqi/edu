package main

import (
	_ "edu/docs"
	_ "edu/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:abcqq.com@tcp(localhost:3306)/edu")
	//redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
}

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
