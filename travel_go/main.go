package main

import (
	_ "travel/boot"
	_ "travel/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	//server := g.Server()
	//server.Use(MiddlewareCORS)
	g.Server().Run()
}
