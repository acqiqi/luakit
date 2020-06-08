package task

import (
	"log"
	"luakit/utils"
	"time"
)

// 自动服务任务
func ServiceTask() error {
	log.Println("Loop Service")
	cb := new(utils.CallBackStrcut)
	if err := utils.HttpPostJson(
		utils.ServiceConfig.String("webapi_url")+
			"api/task/smServiceTask", utils.GetEmptyStruct, &cb); err != nil {
		return err
	}
	log.Println(cb)
	time.Sleep(60 * time.Second)
	return nil
}
