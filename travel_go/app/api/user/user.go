package user

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"time"
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
	// 校验参数
	if err := r.Parse(&registerForm); err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}
	registerForm.Validate(r)

	registerForm.Save(r)

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
	token, err := service.CreateToken(user.Id, tbl_user.Salt, time.Now().Add(time.Minute*60*24*2).Unix())

	if err != nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("登录失败!"))
	}

	// todo 更新登录时间
	user.LastLoginTime = gtime.Now()
	user.Update()

	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "登录成功!",
		Result:  token,
	})
}

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

func (*Account) SendMail(r *ghttp.Request) {
	var sendMailForm *form.SendRegisterMailForm
	if err := r.Parse(&sendMailForm); err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}
	// 随机获取code
	randInt := grand.Digits(6)

	mailList := []string{sendMailForm.Email}
	title := "欢迎注册travel"
	body := fmt.Sprintf(`
	<p>尊敬的travel用户，您好：</p>
	<p>欢迎注册travel！</p>
	<p>这是来自travel官方的验证邮件，来验证您的邮箱是否真实有效。</p>
	<p>您的注册code为【%s】，有效期为10分钟。</p>
	<p>如若误发，请忽略该邮件!</p>
`, randInt)

	// 开启协程发送邮件
	go service.SendMail(mailList, title, body)

	coon := g.Redis().Conn()
	defer coon.Close()

	// 将code存入redis  设置10分钟的有效期
	if _, err := coon.Do("SET", sendMailForm.Email, randInt); err != nil {
		fmt.Println(err)
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("服务器开小差了!"))
	}

	expireIn := 60 * 10

	if _, err := coon.Do("EXPIRE", sendMailForm.Email, expireIn); err != nil {
		fmt.Println(err)
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("服务器开小差了!"))
	}

	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "发送邮件成功!",
		Result:  map[string]interface{}{"expire_in": expireIn},
	})
}
