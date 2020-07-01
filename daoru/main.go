package daoru

import (
	"fmt"
	"io/ioutil"
	"log"
	"luakit/models"
	"luakit/utils"
	"strconv"
	"time"
)

func DaoruManagerUsers() {
	b, err := ioutil.ReadFile("daoru/files/vhake_manager_users.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	var list []ManagerUsers
	utils.JsonDecode(str, &list)
	log.Println(list)

	for i, v := range list {
		if v.Mobile == "" {
			//直接跳出当前循环
			continue
		}
		m_id, _ := strconv.Atoi(v.ID)
		m_score, _ := strconv.Atoi(v.Score)
		money, _ := strconv.ParseFloat(v.Money, 32)

		lv1, _ := strconv.Atoi(v.ShareLv1)
		lv2, _ := strconv.Atoi(v.ShareLv2)

		u, err := models.GetUcenterUsersByMobile(v.Mobile)
		if err == nil {
			log.Println("有数据")
			log.Println("对比1")
			log.Println(u.Mobile)
			log.Println("对比2")
			log.Println(v.Mobile)
			return
		} else {
			log.Println("无数据")
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
			//models.AddUcenterUsers(&model)
			log.Println(model)
		}
		log.Println(i)
	}
}

func daoruComment() {
	b, err := ioutil.ReadFile("daoru/files/vhake_comment.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	var list []Comment
	utils.JsonDecode(str, &list)

	for i, v := range list {
		c_at, _ := time.ParseInLocation("2006/1/2 15:04:05", v.CreatedAt, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
		u_at, _ := time.ParseInLocation("2006/1/2 15:04:05", v.UpdatedAt, time.Local)

		quck_id, _ := strconv.Atoi(v.QuckID)

		cuid, _ := strconv.Atoi(v.UID)

		if quck_id == 2 && v.Content != "" {
			usr, err := models.GetUcenterUsersById(int64(cuid))
			if err == nil {
				model := models.SmUsersComment{
					Id:             0,
					CreatedAt:      c_at,
					UpdatedAt:      u_at,
					Flag:           1,
					ServiceId:      1,
					ApptId:         0,
					ServiceSkuId:   0,
					OrderId:        0,
					UsersServiceId: 0,
					ServiceAreaId:  0,
					Cuid:           cuid,
					Type:           0,
					Content:        v.Content,
					Pics:           v.Pics,
					VideoUrl:       "",
					Avatar:         usr.Avatar,
					Nickname:       usr.Nickname,
					Star:           5,
					Tags:           "",
					IsOld:          1,
				}
				models.AddSmUsersComment(&model)
				log.Println(i)
			}
		}

	}
}

func daoruAccount() {
	b, err := ioutil.ReadFile("daoru/files/vhake_accounts.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	var list []Accounts
	utils.JsonDecode(str, &list)
	//log.Println(list)

	for i, v := range list {
		user_type, _ := strconv.Atoi(v.UserType)

		if user_type == 0 {
			c_at, _ := time.ParseInLocation("2006/1/2 15:04:05", v.CreatedAt, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
			u_at, _ := time.ParseInLocation("2006/1/2 15:04:05", v.UpdatedAt, time.Local)
			ctype, _ := strconv.Atoi(v.ProjectType)
			cuid, _ := strconv.Atoi(v.UID)

			lv, _ := strconv.Atoi(v.Lv)
			price, _ := strconv.ParseFloat(v.Price, 32)
			share_uid, _ := strconv.Atoi(v.ClientUID)

			if ctype == 0 || ctype == 1 || ctype == 2 {
				log.Println("ojbk")

				model := models.UcenterAccounts{
					Id:          0,
					CreatedAt:   c_at,
					UpdatedAt:   u_at,
					Flag:        1,
					Cuid:        cuid,
					PlatformKey: "DDSM_CLIENT",
					Type:        ctype,
					Level:       lv,
					Content:     v.Content,
					Describe:    v.Desc,
					ProjectId:   0,
					OrderId:     0,
					OrderNo:     0,
					Price:       price,
					IsDz:        1,
					SourceCuid:  share_uid,
					ProjectName: "",
					Title:       v.Desc,
					AccountNo:   "DDSMOLD" + v.ID,
					IsOld:       1,
				}
				models.AddUcenterAccounts(&model)
				log.Println(i)
			}

		}

	}

}

func daoruClientUsers() {
	b, err := ioutil.ReadFile("daoru/files/vhake_client_users.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'
	var list []ClientUsers
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
