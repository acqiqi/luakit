package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"log"
	"luakit/utils"
	"time"
)

func init() {
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local",
		utils.DataBaseObj.String("username"),
		utils.DataBaseObj.String("password"),
		utils.DataBaseObj.String("host"),
		utils.DataBaseObj.String("port"),
		utils.DataBaseObj.String("database"))
	log.Printf("Init Orm")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connect_str, 30)
	orm.DefaultTimeLoc = time.UTC
}
