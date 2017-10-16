package main

import (
	_ "webdemo/routers"
	_"webdemo/models/user_db"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

