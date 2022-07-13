package git

import "time"

type DomainObject struct {
	// 主键
	Id int
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

func New() DomainObject {
	return DomainObject{
		PullAt: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}
