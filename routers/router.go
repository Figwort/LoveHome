package routers

import (
	"github.com/astaxie/beego"
	"test/LoveHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/v1.0/areas", &controllers.AreaControllers{}, "Get:GetAreaInfo")

	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "Get:GetSessionName;Delete:DelSessionName")

	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "Get:GetHousesIndex")

	beego.Router("/api/v1.0/users", &controllers.UserControllers{}, "post:Reg")

	beego.Router("/api/v1.0/sessions", &controllers.UserControllers{}, "post:Login")
}
