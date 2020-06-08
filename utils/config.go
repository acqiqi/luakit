package utils

import (
	"github.com/astaxie/beego/config"
)

var DataBaseObj config.Configer
var ServiceConfig config.Configer
var RedisConfig config.Configer

func init() {
	DataBaseObj, _ = config.NewConfig("ini", "conf/database.conf")
	ServiceConfig, _ = config.NewConfig("ini", "conf/service.conf")
	RedisConfig, _ = config.NewConfig("ini", "conf/redis.conf")
}
