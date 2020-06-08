package task

import (
	"log"
	"luakit/utils"
)

// 优惠券自动生成任务
func GenerCouponTask() error {
	log.Println("Loop GenerCouponTask")
	cb := new(utils.CallBackStrcut)
	if err := utils.HttpPostJson(utils.ServiceConfig.String("webapi_url")+
		"api/task/conponQueueTask", utils.GetEmptyStruct, &cb); err != nil {
		return err
	}
	log.Println(cb)
	return nil
}
