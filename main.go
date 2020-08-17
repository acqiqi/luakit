package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"log"
	"luakit/common"
	_ "luakit/routers"
	_ "luakit/task"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()

	//go mqtt.Init()
	log.Println("Init Project")
	//setup()
	u := common.ExcelUtils{}
	u.Test()

}

func setup() {
	beego.Run()
}
