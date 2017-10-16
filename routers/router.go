package routers

import (
	"webdemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.IndexCotroller{})
    beego.Router("/", &controllers.MainController{})
}
