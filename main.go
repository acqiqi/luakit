package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"log"
	"luakit/mqtt"
	_ "luakit/routers"
	_ "luakit/task"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()

	go mqtt.Init()
	log.Println("Init Project")
	setup()
	//reptile.GetDYRSNewsAll(14150,14300)
}

func setup() {
	beego.Run()
}
