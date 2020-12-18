package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"log"
	"luakit/mqtt"
	_ "luakit/routers"
	_ "luakit/task"
	//"luakit/utils"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()
	//return
	//查询列表
	go mqtt.Init()
	setup()
	//reptile.GetDYRSNewsAll(14150,14300)
}

func setup() {
	beego.Run()
}
