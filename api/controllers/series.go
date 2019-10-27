package controllers

import (
	"github.com/astaxie/beego"
	"github.com/linbaozhong/car-care/api/models"
)

type Series struct {
	Base
}
// @Summary 读取全部车系列表
// @Description 读取全部车系列表
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /gets [get]
func (c *Series) Gets() {
	beans, e := models.GetAllSeries()
	if e != nil {
		beego.Error(e)
		c.Result.Status = -200
		c.Result.Message = e.Error()
		return
	}
	c.Result.Data = beans
}