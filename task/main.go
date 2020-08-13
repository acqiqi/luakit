package task

import (
	"github.com/astaxie/beego/toolbox"
	"log"
)

type Task struct {
	TaskName string
	Spec     string
	Func     toolbox.TaskFunc
}

var Tk []Task

func init() {
	var tk = make([]Task, 4)
	tk[0] = Task{TaskName: "message_queue", Spec: "0/1 * * * * *", Func: MessagePushQueue}
	tk[1] = Task{TaskName: "service_task", Spec: "1 * * * * *", Func: ServiceTask}
	tk[2] = Task{TaskName: "coupon_task", Spec: "1 * * * * *", Func: GenerCouponTask}
	tk[3] = Task{TaskName: "packet_end", Spec: "1 * * * * *", Func: PacketEnd}
	Tk = tk
	run()
}

func run() {
	log.Println("Run Task")
	for _, v := range Tk {
		tkobj := toolbox.NewTask(v.TaskName, v.Spec, v.Func)
		toolbox.AddTask(v.TaskName, tkobj)
	}
	toolbox.StartTask()
}
