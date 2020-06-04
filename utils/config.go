package utils

import (
	"github.com/astaxie/beego/config"
)

var DataBaseObj config.Configer

func init() {
	DataBaseObj, _ = config.NewConfig("ini", "conf/database.conf")
}
