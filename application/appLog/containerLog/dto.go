package containerLog

import (
	"fops/domain/_/eumLogLevel"
	"time"
)

// DTO 容器日志记录
type DTO struct {
	// 主键
	Id string
	// 应用名称
	AppName string
	// 容器名称
	ContainerName string
	// 镜像名称
	ContainerImage string
	// 容器IP
	ContainerIp string
	// 环境变量
	ContainerEnv map[string]string
	// 节点名称
	NodeName string
	// 节点IP
	NodeIp string
	// 日志级别
	LogLevel eumLogLevel.Enum
	// 日志内容
	Content string
	// 日志时间
	CreateAt time.Time
}
