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
	service.ParseValidateAndSave(registerForm, r)
}

// 登录接口
func (*Account) Login(r *ghttp.Request) {
	var loginForm *form.LoginForm
	service.ParseValidateAndSave(loginForm, r)
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
	service.ParseValidateAndSave(sendMailForm, r)
}
