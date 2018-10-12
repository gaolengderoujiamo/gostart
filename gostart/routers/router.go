package routers

import (
	"gostart/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/testwrite", &controllers.MainController{}, "get:TestWrite")
	beego.Router("/testjson", &controllers.MainController{}, "get:TestJSON")

	// 正则路由
	beego.Router("/user/add/?:id/?:name", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/update/?:id/?:name", &controllers.UserController{}, "get:UpdateUser")
	beego.Router("/user/list", &controllers.UserController{}, "get:ListUser")
	beego.Router("/user/get/:id", &controllers.UserController{}, "get:FindUserById")
	beego.Router("/user/delete/:id", &controllers.UserController{}, "get:DeleteUserById")

	// 未测通, 不能从RequestBody中获取json
	//beego.Router("/user/addJSON", &controllers.UserController{}, "post:Post")

}
