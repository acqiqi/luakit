package models

import (
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type BaseModel struct {
	Id        int64
	CreatedAt string `orm:"auto_now_add;type(datetime)";json:"created_at"`
	UpdatedAt string `orm:"auto_now;type(datetime)";json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
