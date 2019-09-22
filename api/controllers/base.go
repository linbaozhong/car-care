package controllers

import (
	"github.com/astaxie/beego"
	"github.com/linbaozhong/car-care/api/models"
)
var(

)
type Base struct {
	beego.Controller
	Result models.Result
}

func (c *Base)Prepare() {
	c.Result.Status=200
}
func (c *Base)Finish() {
	c.Data["json"]=c.Result
	c.ServeJSON(true)
}