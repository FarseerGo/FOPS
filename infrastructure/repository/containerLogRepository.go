package repository

import (
	"fops/domain/appLog/containerLog"
	"fs/core/container"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() containerLog.Repository { return &containerLogRepository{} })
}

type containerLogRepository struct {
}

// ToList 读取前500条日志
func (repository containerLogRepository) ToList(top int) []containerLog.DomainObject {
	//ContainerLogAgent.ToList(top).MapAsync<ContainerLogDO, ContainerLogPO>()
	return nil
}
