package main

import (
	_ "myapp/beeblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

