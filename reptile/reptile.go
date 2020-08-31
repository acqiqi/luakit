package reptile

import (
	"fmt"
	"log"
	"luakit/utils"
	"os"
	"strconv"
)

// 获取东易日盛 装修攻略
func GetDYRSNewsAll(begin_id int, end_id int) error {
	url := "https://app2020ad.dyrs.com.cn/api/app/detail/getnews?token=&id="

	for i := begin_id; i < end_id; i++ {
		callback, err := utils.HttpGetString(url + strconv.Itoa(i))
		log.Println(callback)
		if err == nil {
			f, err := os.Create("./reptile/database/dyrs/news/" + strconv.Itoa(i) + ".json")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				_, err = f.Write([]byte(callback))
			}
			f.Close()
		}
	}
	return nil
}
