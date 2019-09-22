package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/linbaozhong/car-care/api/models"
)

type Brand struct {
	Base
}

// @Title Gets
// @Description 读取全部品牌列表
// @Success 200 {object} models.Result
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

// @Title Test
// @Description 读取全部品牌列表
// @Success 200 {object} models.Result
// @router /test [get]
func (c *Brand) Test() {
	beans, e := models.GetAllBrands()
	if e != nil {
		beego.Error(e)
		return
	}
	for _, brand := range beans {
		//url:=fmt.Sprintf("https://item.tuhu.cn/Car/SelOneBrand?callback=__GetCarBrands__&Brand=%s&_=1569119617781","")
		//req:=httplib.Get()
		url := fmt.Sprintf("https://item.tuhu.cn/Car/SelOneBrand?Brand=%s&_=1569119617781", brand.First+" - "+brand.Name)
		fmt.Println(url)
	}
}
