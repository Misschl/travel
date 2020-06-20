package form

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"travel/app/model/tbl_user"
	"travel/app/response"
	"travel/utils"
)

type RegisterForm struct {
	Email     string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
	Password  string `p:"password" v:"required|length:6,18 #请输入密码|账号长度为:min到:max位"`
	Rpassword string `p:"rpassword" v:"required|length:6,18|same:password #请再次输入密码|账号长度为:min到:max位|两次密码输入不一致"`
	Code      int    `p:"code"  v:"required#验证码为必填字段"`
}

func (this *RegisterForm) existEmail(r *ghttp.Request) {
	ok, err := tbl_user.ExistEmail(this.Email)
	if err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}
	if ok {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse("邮箱已存在!"))
	}
}

func (this *RegisterForm) createUser(r *ghttp.Request) {
	result, err := tbl_user.CreateUser(this.Email, this.Password)
	if err != nil || result == nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("注册失败!"))
	}
}

func (this *RegisterForm) checkCode(r *ghttp.Request) {
	// 校验邮箱的验证码
	coon := g.Redis().Conn()
	defer coon.Close()

	reply, err := coon.Do("GET", this.Email)
	if err != nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("服务器开小差了!"))
	}

	if reply == nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse("请先获取验证码!"))
	}

	// 将redis里的数据转换成int
	code, err := utils.Uint8SliceToInt(reply.([]uint8))

	if err != nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("服务器开小差了!"))
	}

	if this.Code != code {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse("邮箱验证码错误!"))
	}
}

func (this *RegisterForm) Validate(r *ghttp.Request) {

	// 校验邮箱的验证码
	this.checkCode(r)
	// 校验email
	this.existEmail(r)
}

func (this *RegisterForm) Save(r *ghttp.Request) {
	this.createUser(r)
}

type LoginForm struct {
	Email    string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
	Password string `p:"password" v:"required|length:6,18 #请输入密码|账号长度为:min到:max位"`
}

type SendRegisterMailForm struct {
	Email string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
}
