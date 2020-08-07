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
	//i := utils.GetRand([]int{100000, 100, 200, 300, 5000})
	//log.Println(i)
	//utils.SendAdminSmsLuosimao("中间件测试短信")

	//accounts_utils := new(common.Accounts)
	//accounts_utils.SmUsersServiceId = 277
	//if err := accounts_utils.InitSmUsersService();err != nil {
	//	log.Println("mdzz?"+err.Error())
	//}
	//if err := accounts_utils.AccountsSmServiceOrders() ;err != nil {
	//	log.Println(err)
	//}

}

func setup() {
	beego.Run()
}
