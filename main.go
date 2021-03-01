package main

import (
	"github.com/astaxie/beego"
	_ "microsvc02/conf"
	_ "microsvc02/filter"
	"microsvc02/routers"
)

func main() {
	routers.Router()
	beego.Run()

}
