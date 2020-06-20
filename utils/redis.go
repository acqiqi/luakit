package utils

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func init() {
}

// 入列
func RedisLPush(key string, value string) error {
	conn, err := connectRedis()
	defer conn.Close()
	if err != nil {
		return err
	}
	_, err = conn.Do("lpush", key, value)
	return err
}

// 出列
func RedisLPop(key string) (v string, err error) {
	conn, err := connectRedis()
	defer conn.Close()
	if err != nil {
		return "", err
	}
	str, err := redis.String(conn.Do("lpop", key))
	if err != nil {
		return "", err
	}
	return str, nil
}

func connectRedis() (conn redis.Conn, err error) {
	ip := RedisConfig.String("ip")
	port := RedisConfig.String("port")
	c, err := redis.Dial("tcp", ip+port)
	if err != nil {
		log.Println("connect to redis err", err.Error())
		return nil, err
	}
	return c, nil
}
