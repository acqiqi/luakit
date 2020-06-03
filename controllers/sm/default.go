package sm

import (
	"github.com/astaxie/beego"
)

type SmController struct {
	beego.Controller
}

func (c *SmController) Test() {
	var cc = struct {
		Heihei string `json:"he1ihei"`
	}{Heihei: "nmsl"}
	c.Data["json"] = cc
	c.ServeJSON()
}
