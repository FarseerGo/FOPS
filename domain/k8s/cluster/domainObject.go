package cluster

import (
	"fops/domain/_/eumRuntimeEnv"
)

type DomainObject struct {
	// 主键
	Id int
	// 集群名称
	Name string
	// 本地kubectl配置地址
	Config string
	// 排序（越小越前）
	Sort int
	// 集群环境类型
	RuntimeEnvType eumRuntimeEnv.Enum
}

func New() DomainObject {
	return DomainObject{}
}
