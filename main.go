package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	"log"
	_ "luakit/routers"
	_ "luakit/task"
	_ "luakit/utils"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置全局log 打印带行数
	//utils.InitModel()
	for i := 0; i < 100; i++ {

	}
}

func setup() {
	beego.Run()
}

var goodsNum int = 1000                           // 商品可被抢购的个数
var luckyUser = make(map[interface{}]interface{}) // 幸运用户
var mu sync.Mutex
var num = 50

func SecKill(userData interface{}) {
	if num > 0 { // 在并发的场合，这个num的值可能是脏数据，但是有这个判断的话，可以过滤掉当秒杀完商品之后用户再请求就不会再继续造成资源耗费了
		go func() {
			mu.Lock()
			if num > 0 {
				_, e := luckyUser[userData]
				if !e { // 如果该用户已经抢购成功，则不允许继续抢购了
					luckyUser[userData] = userData
					num--
				}
			}
			mu.Unlock()
		}()
	}
}

var lock = false

func eeee() {
	if num >= 10 {

	} else {
		time.Sleep(time.Second * 1)
		num++
		log.Println("heihei")
	}
}
