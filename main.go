package main

import (
	_ "fops/infrastructure/repository"
	"fops/interfaces/controller"
	"github.com/beego/beego/v2/server/web"
	"github.com/farseernet/farseer.go/configure"
	"github.com/farseernet/farseer.go/fsApp"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	fsApp.Initialize("FOPS")
	web.BConfig.WebConfig.Session.SessionOn = true
	configure.SetDefault("FSSServer", "http://fss:888")

	go func() {
		for {
			log.Println("当前routine数量:", runtime.NumGoroutine())
			time.Sleep(time.Second)
		}
	}()
	// get http://localhost:8080/api/user/123
	web.CtrlGet("api/user/:id", controller.UserController.GetUserById)
	web.Run()
}

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("C语言中文网"))
}

func index2(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("C语言中文网")
}
