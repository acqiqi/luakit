package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"luakit/utils"
	"time"
)

type ManagerUsers struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)" json:"updated_at"`
	Flag      int       `orm:"default(1)" json:"flag"` //-1删除

	Username      string  `json:"username"`        // 用户名
	Nickname      string  `json:"nickname"`        // 昵称
	Password      string  `json:"password"`        // 密码
	Email         string  `json:"email"`           // 邮箱地址
	Mobile        string  `json:"mobile"`          // 手机号码
	Avatar        string  `json:"avatar"`          // 头像
	Money         float64 `json:"money"`           // 余额
	Score         int     `json:"score"`           // 积分
	SignupIp      int     `json:"signup_ip"`       // 注册ip
	LastLoginTime int     `json:"last_login_time"` // 最后一次登录时间
	LastLoginIp   int     `json:"last_login_ip"`   // 登录ip
	Sort          int     `json:"sort"`            // 排序
	Status        int     `json:"status"`          // 状态：0禁用，1启用
	InfoBind      int     `json:"info_bind"`       // 是否绑定用户昵称和头像
	Gender        string  `json:"gender"`          // 性别
	UserKey       string  `json:"user_key"`        // 用户唯一编码
	SmallOpenid   string  `json:"small_openid"`
	RegFlag       int     `json:"reg_flag"`
	RegTime       int     `json:"reg_time"`
	DeletedAt     string  `json:"deleted_at"`
	AuthType      int     `json:"auth_type"` // 0无认证 1企业认证 2施工队 3施工人员
	IsAuth        int     `json:"is_auth"`   // 是否认证
	Lv            int     `json:"lv"`
	Pingfen       float64 `json:"pingfen"` // 评分
	ManagerName   string  `json:"manager_name"`
	Pid           int     `json:"pid"`
	PidTwo        int     `json:"pid_two"`
	Muid          int     `json:"muid"`
	Sgzl          float64 `json:"sgzl"`    // 施工质量
	Sgjd          float64 `json:"sgjd"`    // 施工进度
	Fwtd          float64 `json:"fwtd"`    // 服务态度
	Gdgl          float64 `json:"gdgl"`    // 工地管理
	Sc            string  `json:"sc"`      // 擅长标签
	Gz            string  `json:"gz"`      // 工种
	SgCity        string  `json:"sg_city"` // 施工区域
	Shifu         int     `json:"shifu"`
	ShareLv1      int     `json:"share_lv1"`
	ShareLv2      int     `json:"share_lv2"`
	Shouyi        float64 `json:"shouyi"`
	IsTopic       int     `json:"is_topic"`
	StShouyi      float64 `json:"st_shouyi"`
	ShareShouyi   float64 `json:"share_shouyi"`
	Latitude      float64 `json:"latitude"`  // 纬度
	Longitude     float64 `json:"longitude"` // 经度
	IsSync        int     `json:"is_sync"`   //是否同步
}

func init() {
	orm.RegisterModelWithPrefix(utils.DataBaseObj.String("prefix"), new(ManagerUsers))
}

//
func GetManagerUserById(id int64) (v *ManagerUsers, err error) {
	o := orm.NewOrm()
	v = &ManagerUsers{}
	if err = o.QueryTable(new(ManagerUsers)).Filter("Id", id).Filter("Flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 获取未同步的数据
func GetManagerUserNotSync() (v *ManagerUsers, err error) {
	o := orm.NewOrm()
	v = &ManagerUsers{}
	if err = o.QueryTable(new(ManagerUsers)).Filter("IsSync", 0).Filter("Flag", 1).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 更新数据
func UpdateManagerUserById(m *ManagerUsers) (err error) {
	o := orm.NewOrm()
	v := ManagerUsers{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}
