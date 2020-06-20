package boot

import (
	"github.com/gogf/gf/frame/g"
	"travel/app/middleware"
	_ "travel/packed"
)

func init() {
	server := g.Server()
	server.Use(middleware.MiddlewareCORS)
}
