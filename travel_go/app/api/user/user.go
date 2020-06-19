package user

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"travel/app/form"
	"travel/app/model/tbl_user"
	"travel/app/response"
)

type Account struct{}

// 注册接口
func (*Account) Register(r *ghttp.Request) {
	var registerForm *form.RegisterForm
	err := r.Parse(&registerForm)
	if err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}
	ok, err := tbl_user.ExistEmail(registerForm.Email)
	if err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}
	if ok {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse("邮箱已存在!"))
	}
	result, err := tbl_user.CreateUser(registerForm.Email, registerForm.Password)
	if err != nil || result == nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("注册失败!"))
	}
	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "注册成功!",
	})
}

// 登录接口
func (*Account) Login(r *ghttp.Request) {
	var loginForm *form.LoginForm
	err := r.Parse(&loginForm)
	if err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}
	user, ok := tbl_user.Authenticate(loginForm.Email, loginForm.Password)
	if !ok {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse("用户名或密码错误!"))
	}
	// todo add jwt here
	// todo 更新登录时间
	fmt.Println(user)
	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "登录成功!",
	})
}
