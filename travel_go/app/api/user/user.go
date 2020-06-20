package user

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"travel/app/form"
	"travel/app/middleware"
	"travel/app/model/tbl_user"
	"travel/app/response"
	"travel/app/service"
)

type Account struct{}

// 注册接口
func (*Account) Register(r *ghttp.Request) {
	var registerForm *form.RegisterForm

	if err := r.Parse(&registerForm); err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}

	service.ValidateAndSave(registerForm, r)
}

// 登录接口
func (*Account) Login(r *ghttp.Request) {
	var loginForm *form.LoginForm

	if err := r.Parse(&loginForm); err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}

	service.ValidateAndSave(loginForm, r)
}

// 用户信息接口
func (*Account) Info(r *ghttp.Request) {
	middleware.LoginRequired(r)

	uid := r.GetCtxVar("uid")

	record, _ := g.DB().Table(tbl_user.Table).Where("id=?", uid).FieldsEx(tbl_user.Columns.Password).One()

	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "获取用户信息成功!",
		Result:  record,
	})
}

// 发送邮件接口
func (*Account) SendMail(r *ghttp.Request) {
	var sendMailForm *form.SendRegisterMailForm

	if err := r.Parse(&sendMailForm); err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}

	service.ValidateAndSave(sendMailForm, r)
}
