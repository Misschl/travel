package router

import (
	"github.com/gogf/gf/frame/g"
	"travel/app/api/user"
)

func init() {
	s := g.Server()
	s.BindObject("/account", new(user.Account))
}
