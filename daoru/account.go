package daoru

type Accounts struct {
	ID          string `json:"id"`
	UID         string `json:"uid"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
	Flag        string `json:"flag"`
	UserType    string `json:"user_type"`
	ProjectType string `json:"project_type"`
	ProjectID   string `json:"project_id"`
	Type        string `json:"type"`
	Lv          string `json:"lv"`
	Content     string `json:"content"`
	Price       string `json:"price"`
	IsDz        string `json:"is_dz"`
	ProductName string `json:"product_name"`
	ProjectUID  string `json:"project_uid"`
	CardName    string `json:"card_name"`
	CardID      string `json:"card_id"`
	CardAdr     string `json:"card_adr"`
	CardAddress string `json:"card_address"`
	IsTixian    string `json:"is_tixian"`
	OrderNo     string `json:"order_no"`
	Desc        string `json:"desc"`
	ClientUID   string `json:"client_uid"`
	Muid        string `json:"muid"`
	ShareUID    string `json:"share_uid"`
	Lv1UID      string `json:"lv1_uid"`
}
