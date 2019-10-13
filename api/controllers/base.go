package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	jsoniter "github.com/json-iterator/go"
	"github.com/linbaozhong/car-care/api/models"
	"github.com/linbaozhong/car-care/api/utils"
	"time"
)

var (
	json         = jsoniter.ConfigCompatibleWithStandardLibrary
	submit_cache = utils.NewCache("submit", 5*time.Second)
)

type Base struct {
	beego.Controller
	CurrentID int64
	Result    *models.Result
}

func (c *Base) Prepare() {
	c.Result = models.NewResult()
}

// 中断请求并响应错误
func (c *Base) about() {
	c.Result.Status = 0
	c.Finish()
	c.StopRun()
}

// 生成登录token
func (b *Base) getToken(tk string, id int64) string {
	return fmt.Sprintf("%s_%v", tk, utils.ToHex(id))
}
func (c *Base) Finish() {
	c.Data["json"] = c.Result
	c.ServeJSON(true)
	c.Result.Release()
}

type ClerkAuthorized struct {
	Base
}

func (c *ClerkAuthorized) Prepare() {
	c.Base.Prepare()
	tk := c.GetString("tk")
	id, _ := utils.ParseIdTokenByTk(tk)
	if id < 1 {
		c.Result.Status = 0
		c.Result.Message = "无效的tk"
		c.about()
		return
	}

	// 防范重复提交
	if c.Ctx.Input.IsPost() {
		c.reSubmit()
	}
	token,ok:=auth_cache.Get(id)
	if !ok || tk != token{
		c.Result.Status = 0
		c.Result.Message = "错误的tk"
		c.about()
		return
	}
	c.CurrentID = id
}

// 阻止10秒内重复提交
func (c *ClerkAuthorized) reSubmit() {
	key := fmt.Sprintf("request:%s:%s", c.Ctx.Input.URI(), json.Get(c.Ctx.Input.RequestBody).ToString())

	if _, ok := submit_cache.Get(key); ok {
		c.about()
		c.Result.Message = "重复请求"
		return
	}
	submit_cache.Set(key, 1)
}
func (c *ClerkAuthorized) Finish() {
	c.Base.Finish()
}