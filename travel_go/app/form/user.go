package form

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/grand"
	"time"
	"travel/app/model/tbl_user"
	"travel/app/response"
	"travel/app/service"
	"travel/utils"
)

type RegisterForm struct {
	Email     string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
	Password  string `p:"password" v:"required|length:6,18 #请输入密码|账号长度为:min到:max位"`
	Rpassword string `p:"rpassword" v:"required|length:6,18|same:password #请再次输入密码|账号长度为:min到:max位|两次密码输入不一致"`
	Code      string `p:"code"  v:"required#验证码为必填字段"`
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

	code := utils.Uint8SliceString(reply.([]uint8))
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

func (this *RegisterForm) writeResponse(r *ghttp.Request) {
	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "注册成功!",
	})
}

func (this *RegisterForm) Save(r *ghttp.Request) {
	this.createUser(r)
	this.writeResponse(r)
}

type LoginForm struct {
	Email    string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
	Password string `p:"password" v:"required|length:6,18 #请输入密码|账号长度为:min到:max位"`
}

func (this *LoginForm) authenticate(r *ghttp.Request) {
	user, ok := tbl_user.Authenticate(this.Email, this.Password)

	if !ok {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse("用户名或密码错误!"))
	}

	this.update(user)

	r.SetCtxVar("userId", user.Id)
}

func (this *LoginForm) getJwt(r *ghttp.Request) {
	// todo add jwt here

	userId := r.GetCtxVar("user").Int()

	token, err := service.CreateToken(userId, tbl_user.Salt, time.Now().Add(time.Minute*60*24*2).Unix())

	if err != nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("登录失败!"))
	}

	r.SetCtxVar("token", token)

}

func (this *LoginForm) update(user *tbl_user.Entity) {
	// todo 更新登录时间
	user.LastLoginTime = gtime.Now()
	user.Update()
}

func (this *LoginForm) writeResponse(r *ghttp.Request) {
	token := r.GetCtxVar("token").String()
	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "登录成功!",
		Result:  token,
	})
}

func (this *LoginForm) Validate(r *ghttp.Request) {
	this.authenticate(r)
}

func (this *LoginForm) Save(r *ghttp.Request) {
	this.getJwt(r)
	this.writeResponse(r)
}

type SendRegisterMailForm struct {
	Email string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
}

// do nothing, just implement Modeler
func (this *SendRegisterMailForm) Validate(request *ghttp.Request) {}


func (this *SendRegisterMailForm) sendMail(r *ghttp.Request) {
	randInt := r.GetCtxVar("randInt").String()
	mailList := []string{this.Email}
	title := "欢迎注册travel"
	body := fmt.Sprintf(`
	<p>尊敬的travel用户，您好：</p>
	<p>欢迎注册travel！</p>
	<p>这是来自travel官方的验证邮件，来验证您的邮箱是否真实有效。</p>
	<p>您的注册验证码为【%s】，有效期为10分钟。</p>
	<p>如若误发，请忽略该邮件!</p>
`, randInt)

	// 开启协程发送邮件
	go service.SendMail(mailList, title, body)
}

func (this *SendRegisterMailForm) saveCode(r *ghttp.Request) {
	randInt := r.GetCtxVar("randInt").String()
	coon := g.Redis().Conn()
	defer coon.Close()

	// 将code存入redis  设置10分钟的有效期
	if _, err := coon.Do("SET", this.Email, randInt); err != nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("服务器开小差了!"))
	}

	if _, err := coon.Do("EXPIRE", this.Email, 60*10); err != nil {
		_ = r.Response.WriteJsonExit(response.SeverErrorResponse("服务器开小差了!"))
	}
}

func (this *SendRegisterMailForm) writeResponse(r *ghttp.Request) {
	_ = r.Response.WriteJsonExit(response.Response{
		Success: true,
		Message: "发送邮件成功!",
	})
}

func (this *SendRegisterMailForm) Save(r *ghttp.Request) {
	// 随机获取code
	randInt := grand.Digits(6)
	r.SetCtxVar("randInt", randInt)
	this.sendMail(r)
	this.saveCode(r)
	this.writeResponse(r)
}
