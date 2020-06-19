package form

type RegisterForm struct {
	Email     string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
	Password  string `p:"password" v:"required|length:6,18 #请输入密码|账号长度为:min到:max位"`
	Rpassword string `p:"rpassword" v:"required|length:6,18|same:password #请再次输入密码|账号长度为:min到:max位|两次密码输入不一致"`
}

type LoginForm struct {
	Email    string `p:"email" v:"required|email #请输入邮箱|邮箱格式有误"`
	Password string `p:"password" v:"required|length:6,18 #请输入密码|账号长度为:min到:max位"`
}
