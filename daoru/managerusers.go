package daoru

type ManagerUsers struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Avatar        string `json:"avatar"`
	Money         string `json:"money"`
	Score         string `json:"score"`
	SignupIP      string `json:"signup_ip"`
	LastLoginTime string `json:"last_login_time"`
	LastLoginIP   string `json:"last_login_ip"`
	Sort          string `json:"sort"`
	Status        string `json:"status"`
	InfoBind      string `json:"info_bind"`
	Gender        string `json:"gender"`
	UserKey       string `json:"user_key"`
	SmallOpenid   string `json:"small_openid"`
	RegFlag       string `json:"reg_flag"`
	RegTime       string `json:"reg_time"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
	Flag          string `json:"flag"`
	AuthType      string `json:"auth_type"`
	IsAuth        string `json:"is_auth"`
	Lv            string `json:"lv"`
	Pingfen       string `json:"pingfen"`
	ManagerName   string `json:"manager_name"`
	Pid           string `json:"pid"`
	PidTwo        string `json:"pid_two"`
	Muid          string `json:"muid"`
	Sgzl          string `json:"sgzl"`
	Sgjd          string `json:"sgjd"`
	Fwtd          string `json:"fwtd"`
	Gdgl          string `json:"gdgl"`
	Sc            string `json:"sc"`
	Gz            string `json:"gz"`
	SgCity        string `json:"sg_city"`
	Shifu         string `json:"shifu"`
	ShareLv1      string `json:"share_lv1"`
	ShareLv2      string `json:"share_lv2"`
	Shouyi        string `json:"shouyi"`
	IsTopic       string `json:"is_topic"`
	StShouyi      string `json:"st_shouyi"`
	ShareShouyi   string `json:"share_shouyi"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
}
