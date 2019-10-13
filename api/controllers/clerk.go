package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/linbaozhong/car-care/api/models"
)

type Clerk struct {
	ClerkAuthorized
}

// @Summary UpdateInfo
// @Description 修改店员信息
// @Param tk query string true "token"
// @Param name formData string true "姓名"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /updateinfo [post]
func (c *Clerk) UpdateInfo() {
	name := c.GetString("name")
	if !govalidator.StringLength(name, "2", "10") {
		c.Result.Status = 0
		c.Result.Code = "name"
		c.Result.Message = "姓名必须2~10位字符长度"
		return
	}
	ok, e := models.UpdateClerkInfo(c.CurrentID, name)
	if ok {
		return
	}

	c.Result.Status = 0
	if e != nil {
		c.Result.Message = e.Error()
	}
}

// @Summary 注销
// @Description 店员注销
// @Param tk query string true "token"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /logout [post]
func (c *Clerk) Logout() {
	auth_cache.Delete(c.CurrentID)
}
