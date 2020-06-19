package main

import (
	_ "travel/boot"
	_ "travel/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
