package vo

import "time"

// UserVO 用户表
type UserVO struct {
	Id          int32     `json:"id" gorm:"primary_key"` // 主键
	UserName    string    `json:"user_name"`             // 名称
	UserAddress string    `json:"user_address"`          // 地址
	CreateTime  time.Time `json:"create_time"`           // 创建时间
	UpdateTime  time.Time `json:"update_time"`           // 更新时间
	IsDeleted   int32     `json:"is_deleted"`            // 是否删除
	Year        time.Time `json:"year"`
	Time        time.Time `json:"time"`
	Timestamp   time.Time `json:"timestamp"`
	Float       float64   `json:"float"`
	Double      float64   `json:"double"`
	Decimal     float64   `json:"decimal"`
	Char        string    `json:"char"`
	Text        string    `json:"text"`
	Longtext    string    `json:"longtext"`
	Tinytext    string    `json:"tinytext"`
	MediumT     string    `json:"medium_t"`
	SmallInt    int32     `json:"small_int"`
	Bigint      int64     `json:"bigint"`
	Tinyint     int32     `json:"tinyint"`
}
