package controllers

import (
	"github.com/astaxie/beego"
	"gostart/models"
	"strconv"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) AddUser()  {
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	name := this.Ctx.Input.Param(":name")
	u := models.User{Id: id, Name: name}
	// insert
	_, err := orm.NewOrm().Insert(&u)
	if err == nil {
		this.Data["json"] = "{\"msg\": \"success\"}"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

func (this *UserController) UpdateUser() {
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	name := this.Ctx.Input.Param(":name")
	u := models.User{Id: id, Name: name}
	// update
	_, err := orm.NewOrm().Update(&u)
	if err == nil {
		this.Data["json"] = "{\"msg\": \"success\"}"
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

func (this *UserController) ListUser()  {
	var maps []orm.Params
	_, err := orm.NewOrm().Raw("SELECT * FROM user").Values(&maps)
	if err == nil {
		this.Data["json"] = maps
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

func (this *UserController) FindUserById() {
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// read one
	var u models.User
	err := orm.NewOrm().QueryTable("user").Filter("id", id).One(&u);
	if err == nil {
		this.Data["json"] = u
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

func (this *UserController) DeleteUserById() {
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	//delete
	var u models.User
	_, err := orm.NewOrm().QueryTable("user").Filter("id", id).Delete()
	if err != nil {
		this.Data["json"] = u
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

func (this *UserController) Post()  {
	var u models.User
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &u); err == nil {
		orm.NewOrm().Insert(&u)
		this.Data["json"] = "{\"msg\": \"success\"}"
	} else {
		panic(err)
	}
	this.ServeJSON()
}