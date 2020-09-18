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
	//
	//b, err := ioutil.ReadFile("./excelout/1.jpg") // just pass the file name
	//if err != nil {
	//	fmt.Print(err)
	//}
	//
	//list := make([]map[string]interface{}, 10)
	//
	//it := make(map[string]interface{})
	//it["name"] = "name!"
	//it["nickname"] = "nickname!"
	//it["desc"] = b
	//for i := 0; i < 10; i++ {
	//	list[i] = it
	//
	//}
	//
	//keys := make([]utils.ExcelKey, 3)
	//keys[0] = utils.ExcelKey{
	//	Key:  "name",
	//	Name: "Mdzz",
	//}
	//keys[1] = utils.ExcelKey{
	//	Key:  "nickname",
	//	Name: "Haha",
	//}
	//keys[2] = utils.ExcelKey{
	//	Key:  "desc",
	//	Name: "Desc",
	//	Type: "Image",
	//}
	//utilsExcel := new(utils.ExcelUtils)
	//utilsExcel.Title = "导出表格"
	//utilsExcel.Keys = keys
	//utilsExcel.DataBase = list
	//utilsExcel.FilePath = "./excelout"
	//utilsExcel.FileName = "mdzz"
	//utilsExcel.ExcelStyle.Height = 100
	//if err := utilsExcel.Export(); err != nil {
	//	log.Println("导出失败", err.Error())
	//}
	//setup()
	//reptile.GetDYRSNewsAll(14150,14300)
}

func setup() {
	beego.Run()
}
