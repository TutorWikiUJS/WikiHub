package routers

import (
	"github.com/astaxie/beego"
	"github.com/xuzhenglun/WikiHub/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
