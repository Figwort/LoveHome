package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
	_ "test/LoveHome/models"
	_ "test/LoveHome/routers"
)

func main() {

	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath() {
	beego.InsertFilter("/", beego.BeforeRouter, TransparentStaic)
	beego.InsertFilter("/*", beego.BeforeRouter, TransparentStaic)
}

func TransparentStaic(ctx *context.Context) {
	orpath := ctx.Request.URL.Path

	beego.Debug("requst url: ", orpath)

	if strings.Index(orpath, "api") >= 0 {
		return
	}

	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/html/"+ctx.Request.URL.Path)
}
