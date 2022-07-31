package controller

import "github.com/beego/beego/v2/server/web"

type UserController struct {
	web.Controller
}

func (u UserController) GetUserById() {
	u.Ctx.WriteString("GetUserById")
}
