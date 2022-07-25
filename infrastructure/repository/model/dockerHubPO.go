package model

type DockerHubPO struct {
	Id int `gorm:"primaryKey"`
	// 仓库名称
	Name string
	// 托管地址
	Hub string
	// 账户名称
	UserName string
	// 账户密码
	UserPwd string
}
