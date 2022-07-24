package admin

import "time"

type PO struct {
	Id int `gorm:"primaryKey"`
	// 管理员名称
	UserName string
	// 管理员密码
	UserPwd string
	// 是否启用
	IsEnable bool
	// 上次登陆时间
	LastLoginAt time.Time
	// 上次登陆IP
	LastLoginIp string
	// 创建时间
	CreateAt time.Time
	// 创建人
	CreateUser string
	// 创建人ID
	CreateId string
}

func (PO) TableName() string { return "admin" }
