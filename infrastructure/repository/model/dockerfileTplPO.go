package model

type DockerfileTplPO struct {
	Id int `gorm:"primaryKey"`
	// 模板名称
	Name string
	// 模板内容
	Template string
}
