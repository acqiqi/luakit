package sm

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *ApiController) Test() {
	var cc = struct {
		Heihei string `json:"he1ihei"`
	}{Heihei: "nmsl"}
	c.Data["json"] = cc
	c.ServeJSON()
}
