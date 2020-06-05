package utils

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
	"time"
)

var vc cache.Cache

func init() {
	initCache()
}

func initCache() {
	log.Println("Init Cache")
	c, err := cache.NewCache("redis", `{"key":"vhake","conn":"127.0.0.1:6379","dbNum":"0","password":""}`)
	if err != nil {
		log.Println("Cache Error" + err.Error())
	}
	vc = c
	if err != nil {
		log.Println("Cache Error" + err.Error())
	}
}

// 设置字符串型焕春
func SetCacheString(key string, val string, timeout time.Duration) error {
	return vc.Put(key, val, timeout)
}

func GetCacheString(key string) string {
	cb := vc.Get(key)
	if cb == nil {
		return ""
	} else {
		return string(cb.([]byte))
	}
}
