package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"log"
	"luakit/models"
	"luakit/mqtt"
	_ "luakit/routers"
	_ "luakit/task"
	"luakit/utils"

	//"luakit/utils"
	_ "luakit/utils"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()
	//return
	//查询列表
	go mqtt.Init()
	log.Println("Init Project 2020-09-05 09:18:56")

	row, lists, err := models.GetMarketingCardMoneyAllList(1)
	if err != nil {
		log.Println("mdzz Daochu失败")
		return
	}
	list := make([]map[string]interface{}, row)
	for i, v := range *lists {
		db := make(map[string]interface{})
		db["id"] = v.Id
		db["name"] = v.Name
		db["desc"] = v.Desc
		db["price"] = v.Price
		db["card_no"] = v.CardNo
		db["card_password"] = v.CardPassword
		// 获取二维码
		//wechatUtils := new(utils.WechatUtils)
		//wechatUtils.SmallQrcodeData.Path = "card/money/detail?card_no=" + v.CardNo
		//wechatUtils.SmallQrcodeData.Width = 430
		//wechatUtils.Init("wx70618bbce8b722ab", "b4b911ea1ed1757a9b674c7091390f9c")
		//wechatUtils.GetAccessToken()
		//b, err := wechatUtils.GetSmallQrcode()
		//if err != nil {
		//	log.Println(err.Error())
		//}
		//db["qrcode"] = b
		//直接使用Url 模式
		db["qrcode"] = "http://tapi.ddgongjiang.com/statics/qrcode/getSmallQrcode?platform_key=DDSM_CLIENT&path=card/money/detail?card_no=" + v.CardNo
		list[i] = db

	}

	keys := make([]utils.ExcelKey, 7)
	keys[0] = utils.ExcelKey{
		Key:  "id",
		Name: "Id",
	}
	keys[1] = utils.ExcelKey{
		Key:  "name",
		Name: "卡券名称",
	}
	keys[2] = utils.ExcelKey{
		Key:  "desc",
		Name: "描述",
	}
	keys[3] = utils.ExcelKey{
		Key:  "price",
		Name: "金额",
	}
	keys[4] = utils.ExcelKey{
		Key:  "card_no",
		Name: "卡号",
	}
	keys[5] = utils.ExcelKey{
		Key:  "card_password",
		Name: "卡密",
	}
	keys[6] = utils.ExcelKey{
		Key:  "qrcode",
		Name: "二维码",
		//Type:  "Image",
		//Width: 70,
	}
	utilsExcel := new(utils.ExcelUtils)
	utilsExcel.Title = "导出表格"
	utilsExcel.Keys = keys
	utilsExcel.DataBase = list
	utilsExcel.FilePath = "./excelout"
	utilsExcel.FileName = "mdzz"
	//utilsExcel.ExcelStyle.Height = 430
	if err := utilsExcel.Export(); err != nil {
		log.Println("导出失败", err.Error())
	}
	//setup()
	//reptile.GetDYRSNewsAll(14150,14300)
}

func setup() {
	beego.Run()
}
