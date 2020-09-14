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

	//查询列表

	go mqtt.Init()
	log.Println("Init Project 2020-09-05 09:18:56")
	setup()
	//reptile.GetDYRSNewsAll(14150,14300)
}

func setup() {
	beego.Run()
}
