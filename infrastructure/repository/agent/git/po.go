package git

import "time"

type PO struct {
	Id int `gorm:"primaryKey"`
	// Git名称
	Name string
	// 托管地址
	Hub string
	// Git分支
	Branch string
	// 账户名称
	UserName string
	// 账户密码
	UserPwd string
	// 拉取时间
	PullAt time.Time
}
