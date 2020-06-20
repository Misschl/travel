package service

import (
	"github.com/gogf/gf/net/ghttp"
	"travel/app/response"
)

/*
封装一下模型的接口
*/

type Modeler interface {
	Validate(*ghttp.Request) // 校验模块函数
	Save(*ghttp.Request)     // 校验完成后,执行的相应操作
}

// 验证并保存
func ValidateAndSave(m Modeler, r *ghttp.Request) {
	m.Validate(r)
	m.Save(r)
}

// 解析验证并保存
func ParseValidateAndSave(m Modeler, r *ghttp.Request) {
	if err := r.Parse(&m); err != nil {
		_ = r.Response.WriteJsonExit(response.RequestBadResponse(err.Error()))
	}
	ValidateAndSave(m, r)
}
