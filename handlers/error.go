package handlers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error(err error) {
	c.Data["content"] = err
	c.TplName = "error.tpl"
}