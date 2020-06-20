package marketing

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
	"strconv"
	"sync"
)

func (this *SeckillController) Prepare() {
	log.Println("Prepare")
	platform := apibase.Auth2PlatformBase(this.Controller)
	this.Platform = *platform
}

type SeckillController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

var killlock sync.Mutex

func (this *SeckillController) Kill() {
	c := struct {
		Cuid int64 `json:"cuid"`
	}{}
	utils.GetPostJson(this.Controller, &c)
	killlock.Lock()

	user, err := models.GetUcenterUsersById(c.Cuid)
	if err != nil {
		killlock.Unlock()
		utils.ApiErr(this.Controller, "用户不存在")
	}
	log.Println(user)

	seckill, err := models.GetMarketingSeckillPacketOrBegin()
	if err != nil {
		killlock.Unlock()
		utils.ApiErr(this.Controller, "本次红包已经抢完")
	}

	// 检测是否抢过
	_, err = models.GetMarketingSeckillPacketUsersOrUser(user.Id, seckill.Id)
	if err == nil {
		killlock.Unlock()
		utils.ApiErr(this.Controller, "已经抢过，不可重复抢")
	}

	maps := make(map[string]string)
	maps["seckill_packet_id"] = strconv.FormatInt(seckill.Id, 10)
	maps["is_end"] = "0"
	maps["flag"] = "1"
	list, err := models.GetAllMarketingSeckillPacketLink(maps, []string{}, []string{"id"}, []string{"desc"}, 0, 1000)

	if len(list) == 0 {
		seckill.Status = 2
		models.UpdateMarketingSeckillPacketById(seckill)
		killlock.Unlock()
		utils.ApiErr(this.Controller, "本次红包已经抢完")
	}

	log.Println(len(list))
	ran_list := make([]int, len(list))
	for k, v := range list {
		ran_list[k] = v.(models.MarketingSeckillPacketLink).Num - v.(models.MarketingSeckillPacketLink).UseNum
	}

	i := 0
	if len(ran_list) <= 1 {
		i = 0
	} else {
		i = utils.GetRand(ran_list)
	}

	if list[i].(models.MarketingSeckillPacketLink).UseNum >= list[i].(models.MarketingSeckillPacketLink).Num {
		killlock.Unlock()
		utils.ApiErr(this.Controller, "当前红包已经抢完")
	}
	if list[i].(models.MarketingSeckillPacketLink).IsEnd == 1 {
		killlock.Unlock()
		utils.ApiErr(this.Controller, "当前红包已经抢完1")
	}

	//处理红包
	link := new(models.MarketingSeckillPacketLink)
	link.Id = list[i].(models.MarketingSeckillPacketLink).Id
	link.CreatedAt = list[i].(models.MarketingSeckillPacketLink).CreatedAt
	link.UpdatedAt = list[i].(models.MarketingSeckillPacketLink).UpdatedAt
	link.Flag = list[i].(models.MarketingSeckillPacketLink).Flag
	link.SeckillPacketId = list[i].(models.MarketingSeckillPacketLink).SeckillPacketId
	link.Price = list[i].(models.MarketingSeckillPacketLink).Price
	link.Num = list[i].(models.MarketingSeckillPacketLink).Num
	link.UseNum = list[i].(models.MarketingSeckillPacketLink).UseNum + 1
	log.Println("3333")

	if link.UseNum >= link.Num {
		link.IsEnd = 1
	} else {
		link.IsEnd = 0
	}
	err = models.UpdateMarketingSeckillPacketLinkById(link)
	if err != nil {
		killlock.Unlock()
		utils.ApiErr(this.Controller, "操作有误")
	}

	save_data := new(models.MarketingSeckillPacketUsers)
	save_data.Cuid = int(user.Id)
	save_data.PlatformKey = this.Platform.PlatformKey
	save_data.SeckillPacketId = int(seckill.Id)
	save_data.SeckillPacketLinkId = int(link.Id)
	save_data.Flag = 1
	models.AddMarketingSeckillPacketUsers(save_data)

	killlock.Unlock()
	utils.ApiOk(this.Controller, "获取成功", struct {
		Seckill  models.MarketingSeckillPacket      `json:"seckill"`
		Link     models.MarketingSeckillPacketLink  `json:"link"`
		UserLink models.MarketingSeckillPacketUsers `json:"user_link"`
	}{
		Seckill:  *seckill,
		Link:     *link,
		UserLink: *save_data,
	})
	//platform.PlatformKey
}
