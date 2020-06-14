package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"luakit/daoru"
	"luakit/models"
	_ "luakit/routers"
	_ "luakit/task"
	"luakit/utils"
	_ "luakit/utils"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	utils.InitModel()
	log.Println("Init Project")

	//daoruClientUsers() //导入客户端用户

	//setup()
}

func daoruAccount() {
	//b, err := ioutil.ReadFile("daoru/files/vhake_accounts.json") // just pass the file name
	//if err != nil {
	//	fmt.Print(err)
	//}
	//str := string(b) // convert content to a 'string'
	//

}

func daoruClientUsers() {
	b, err := ioutil.ReadFile("daoru/files/vhake_client_users.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	var list []daoru.ClientUsers
	utils.JsonDecode(str, &list)
	for i, v := range list {
		m_id, _ := strconv.Atoi(v.ID)
		m_score, _ := strconv.Atoi(v.Score)
		money, _ := strconv.ParseFloat(v.Money, 32)

		lv1, _ := strconv.Atoi(v.ShareLv1)
		lv2, _ := strconv.Atoi(v.ShareLv2)

		model := models.UcenterUsers{
			Id:             int64(m_id),
			CreatedAt:      time.Time{},
			UpdatedAt:      time.Time{},
			Flag:           1,
			Username:       "",
			Password:       "",
			Mobile:         v.Mobile,
			Nickname:       v.Nickname,
			Email:          v.Email,
			Avatar:         v.Avatar,
			Gender:         v.Gender,
			Status:         1,
			RoleType:       0,
			Score:          m_score,
			Money:          money,
			OkMoney:        money,
			NoMoney:        0,
			LastLoginIp:    "",
			LastLoginTime:  0,
			LastLongitude:  0,
			LastLatitude:   0,
			IsAuth:         0,
			IdcardTop:      "",
			IdcardBom:      "",
			IdcardId:       "",
			ShareOne:       lv1,
			ShareTwo:       lv2,
			StOne:          0,
			StTwo:          0,
			UserKey:        "",
			WechatUnionid:  "",
			RegType:        0,
			RegSource:      "",
			RegPlatformKey: "",
			BindUserinfo:   0,
			IsVip:          0,
			VipEndTime:     "",
			IsPayPassword:  0,
			PayPassword:    "",
		}
		models.AddUcenterUsers(&model)
		log.Println(i)
	}

}

func setup() {
	beego.Run()
}
