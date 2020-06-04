package main

import (
	"github.com/astaxie/beego"
	"log"
	_ "luakit/routers"
	_ "luakit/task"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()

	//o ,err:= models.GetMessageTplById(34)
	//if err != nil {
	//	 log.Println(err.Error())
	//}
	beego.Run()
}
