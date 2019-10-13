package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/linbaozhong/car-care/api/models"
	"github.com/linbaozhong/car-care/api/utils"
	"strconv"
	"time"
)

var (
	auth_cache = utils.NewCache("auth", 24*time.Hour)
)

type Visitor struct {
	Base
}
type Auth struct {
	Token  string `json:"token"`
	Auto   string `json:"auto"`
	Mobile string `json:"mobile"`
}

func NewAuth() *Auth {
	return &Auth{}
}

// @Summary Register
// @Description 注册
// @Param mobile formData int true "手机号码"
// @Param pwd formData string true "密码"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /register [post]
func (c *Visitor) Register() {
	auth := NewAuth()
	c.Result.Data = auth

	mobile, _ := c.GetInt("mobile")
	if !utils.Validate.Mobile(strconv.Itoa(mobile)) {
		c.Result.Status = 0
		c.Result.Code = "mobile"
		c.Result.Message = "手机号码无效"
		return
	}
	pwd := c.GetString("pwd")
	if !govalidator.IsAlphanumeric(pwd) || !govalidator.StringLength(pwd, "6", "12") {
		c.Result.Status = 0
		c.Result.Code = "pwd"
		c.Result.Message = "密码必须是6~12位长度的字母或数字"
		return
	}
	clerk, _ := models.GetClerkByMobile(mobile)
	if clerk != nil {
		c.Result.Status = 0
		c.Result.Code = "mobile"
		c.Result.Message = "手机号码已经存在"
		return
	}
	//新增店员
	clerk = models.NewClerk()
	clerk.Mobile = mobile
	clerk.Pwd = utils.PwdString(pwd, models.SECURITY_CODE)
	clerk.Auto = utils.PwdString(uuid.New().String(), models.SECURITY_CODE)
	clerk.CreateTime = utils.NowMilliSecond()

	ok, e := models.AddClerk(clerk)
	if !ok {
		c.Result.Status = 0
		c.Result.Message = e.Error()
		return
	}
	//返回令牌
	auth.Auto = c.getToken(clerk.Auto, clerk.ID)
	auth.Token = c.getToken(utils.PwdString(uuid.New().String(), models.SECURITY_CODE), clerk.ID)
	c.Result.Data = auth
	//缓存令牌
	auth_cache.Set(clerk.ID, auth.Token)
}

// @Summary Login
// @Description 登录
// @Param mobile formData int true "手机号码"
// @Param pwd formData string true "密码"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /login [post]
func (c *Visitor) Login() {
	auth := NewAuth()
	c.Result.Data = auth

	mobile, _ := c.GetInt("mobile")
	if !utils.Validate.Mobile(strconv.Itoa(mobile)) {
		c.Result.Status = 0
		c.Result.Code = "mobile"
		c.Result.Message = "手机号码无效"
		return
	}
	pwd := c.GetString("pwd")
	if !govalidator.IsAlphanumeric(pwd) || !govalidator.StringLength(pwd, "6", "12") {
		c.Result.Status = 0
		c.Result.Code = "pwd"
		c.Result.Message = "密码必须是6~12位长度的字母或数字"
		return
	}
	clerk, _ := models.GetClerkByMobile(mobile)
	if clerk == nil {
		c.Result.Status = 0
		c.Result.Code = "mobile"
		c.Result.Message = "账户不存在"
		return
	}
	if utils.PwdString(pwd, models.SECURITY_CODE) != clerk.Pwd{
		c.Result.Status = 0
		c.Result.Code = "pwd"
		c.Result.Message = "密码错误"
		return
	}
	//返回令牌
	auth.Auto = c.getToken(clerk.Auto, clerk.ID)
	auth.Token = c.getToken(utils.PwdString(uuid.New().String(), models.SECURITY_CODE), clerk.ID)
	c.Result.Data = auth
	//缓存令牌
	auth_cache.Set(clerk.ID, auth.Token)
}
