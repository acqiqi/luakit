package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"log"
	"luakit/daoru"
	_ "luakit/routers"
	_ "luakit/task"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()
	log.Println("Init Project")
	daoru.DaoruManagerAccount()
	//setup()
	//i := utils.GetRand([]int{100000, 100, 200, 300, 5000})
	//log.Println(i)
}

func setup() {
	beego.Run()
}
