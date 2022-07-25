package repository

import (
	"fops/domain/appLog/containerLog"
	"fops/infrastructure/repository/model"
	"fs/core/container"
	"fs/data"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() containerLog.Repository { return &containerLogRepository{data.TableSet[model.ContainerLogPO]{}} })
}

type containerLogRepository struct {
	data.TableSet[model.ContainerLogPO]
}

// ToList 读取前500条日志
func (repository containerLogRepository) ToList(top int) []containerLog.DomainObject {
	//ContainerLogAgent.ToList(top).MapAsync<ContainerLogDO, ContainerLogPO>()
	return nil
}
