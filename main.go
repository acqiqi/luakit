package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	_ "luakit/models"
	_ "luakit/routers"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数

	//tk1 := toolbox.NewTask("tk1", "1 * * * * *", onTask)
	//toolbox.AddTask("tk11",tk1)
	//toolbox.StartTask()
	//defer toolbox.StopTask()
	//o ,err:= models.GetMessageTplById(34)
	//if err != nil {
	//	 log.Println(err.Error())
	//}
	beego.Run()
}

func onTask() error {
	fmt.Print("heihei")
	return nil
}
