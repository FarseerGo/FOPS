package model

import (
	"fops/domain/_/eumBuildStatus"
	"time"
)

type BuildPO struct {
	Id int `gorm:"primaryKey"`
	// 项目ID
	ProjectId int
	// 项目名称
	ProjectName string
	// 集群ID
	ClusterId int
	// 构建号
	BuildNumber int
	// 状态
	Status eumBuildStatus.Enum
	// 是否成功
	IsSuccess bool
	// 开始时间
	CreateAt time.Time
	// 完成时间
	FinishAt time.Time
	// 构建的服务端id
	BuildServerId int64
}
