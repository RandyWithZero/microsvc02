package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	microSvcDb01Url = "http://svcdb:9091/api/svcdb"
)

type MicroSVC02Controller struct {
	beego.Controller
}

func (c *MicroSVC02Controller) Get() {
	response, _ := http.Get(microSvcDb01Url)
	defer response.Body.Close()
	all, _ := ioutil.ReadAll(response.Body)
	msg2 := string(all)
	getEnv := os.Getenv("HOSTNAME")
	msg := "--->调用svc02" + "[" + getEnv + "]"
	c.Ctx.WriteString(msg + msg2)
}
