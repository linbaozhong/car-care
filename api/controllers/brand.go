package controllers

import (
	"github.com/astaxie/beego"
	"github.com/linbaozhong/car-care/api/models"
)

type Brand struct {
	Base
}

// @Summary Gets
// @Description 读取全部品牌列表
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /gets [get]
func (c *Brand) Gets() {
	beans, e := models.GetAllBrands()
	if e != nil {
		beego.Error(e)
		c.Result.Status = -200
		c.Result.Message = e.Error()
		return
	}
	c.Result.Data = beans
}
