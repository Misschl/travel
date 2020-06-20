package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"travel/app/model/tbl_user"
	"travel/app/response"
	"travel/app/service"
)

func LoginRequired(r *ghttp.Request) {
	token := r.Request.Header.Get("Authorization")
	uid, err := service.ParseToken(token, tbl_user.Salt)
	if err != nil {
		r.Response.Status = 403
		_ = r.Response.WriteJsonExit(response.ForbiddenResponse("无效的登录信息!"))
	}
	r.SetCtxVar("uid", uid)
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
