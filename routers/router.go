package routers

import (

//	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"microsvc02/controllers"
)

func Router()  {
	namespace:= beego.NewNamespace("/api/svc")
	namespace.Router("/",&controllers.MicroSVC02Controller{},"Get:Get")
	beego.AddNamespace(namespace)
}