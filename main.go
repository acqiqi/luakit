package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"log"
	_ "luakit/routers"
	_ "luakit/task"
	"luakit/utils"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()
	utils.RedisLPush("vhake", "heihei")
	str, _ := utils.RedisLPop("vhake")
	log.Println(str)
	setup()
}

func setup() {
	beego.Run()
}
