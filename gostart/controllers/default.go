package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (this *MainController) TestWrite() {
	this.Ctx.WriteString("balabalbababa...")
}

func (this *MainController) TestJSON() {
	m := make(map[string]interface{})
	m["a"] = "ppt"
	m["b"] = "cdf"
	this.Data["json"] = m
	this.ServeJSON()
}