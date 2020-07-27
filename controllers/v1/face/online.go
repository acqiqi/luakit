package face

import (
	"github.com/astaxie/beego"
	"luakit/mqtt"
	"luakit/utils"
	"time"
)

type FaceController struct {
	beego.Controller
}

func (this *FaceController) Mate() {
	c := struct {
		Type      string `json:"type"`
		Url       string `json:"url"`
		MessageId string `json:"message_id"`
		OutUrl    string `json:"out_url"`
		Id        string `json:"id"`
		Timeout   int    `json:"timeout"`
	}{}
	utils.GetPostJson(this.Controller, &c)

	t := struct {
		Type      string
		Url       string
		MessageId string
		OutUrl    string
		Id        string
		Timeout   int
	}{
		Type:      c.Type,
		Url:       c.Url,
		MessageId: c.MessageId,
		OutUrl:    c.OutUrl,
		Id:        c.Id,
		Timeout:   c.Timeout,
	}

	mqtt.Publish("/vhake/hub/sub/face", 2, utils.JsonEncode(t))
	utils.DeleteCacheString("/vhake/hub/pub/face")

	if c.Timeout == 0 {
		c.Timeout = 10
	}

	for i := 0; i < c.Timeout*10; i++ {
		time.Sleep(time.Millisecond * 100)

		callback := utils.GetCacheString("/vhake/hub/pub/face")
		if callback != "" {

			utils.DeleteCacheString("/vhake/hub/pub/face")

			//c.JSON(http.StatusOK, common.ApiJsonSuccess("获取成功", string(decodeBytes)))
			s := struct {
				Value float32
			}{}
			utils.JsonDecode(callback, &s)
			utils.ApiOk(this.Controller, "获取成功", s)
		}
	}
	utils.ApiErr(this.Controller, "超时")
}
