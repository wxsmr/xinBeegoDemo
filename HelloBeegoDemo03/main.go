package main

import (
	"HelloBeegoDemo03/db_mysql"
	_ "HelloBeegoDemo03/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.Connect()
	beego.Run()
}