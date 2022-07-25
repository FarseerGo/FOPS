package repository

import (
	"fops/domain/metaData/dockerHub"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() dockerHub.Repository { return &dockerHubRepository{data.Init[context.MysqlContext]().DockerHub} })
}

type dockerHubRepository struct {
	data.TableSet[model.DockerHubPO]
}

// ToList DockerHub列表
func (repository dockerHubRepository) ToList() []dockerHub.DomainObject {
	lstPO := repository.Desc("Id").ToList()
	return mapper.Array[dockerHub.DomainObject](lstPO)
}

// ToInfo DockerHub信息
func (repository dockerHubRepository) ToInfo(id int) dockerHub.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	return mapper.Single[dockerHub.DomainObject](po)
}

//// Count DockerHub数量
//func (repository dockerHubRepository) Count() int64 {
//	return repository.Count()
//}

// Add 添加仓库
func (repository dockerHubRepository) Add(do dockerHub.DomainObject) {
	po := mapper.Single[model.DockerHubPO](do)
	repository.Insert(&po)
}

// Update 修改仓库
func (repository dockerHubRepository) Update(id int, do dockerHub.DomainObject) {
	po := mapper.Single[model.DockerHubPO](do)
	repository.Where("Id = ?", id).Update(po)
}

// Delete 删除仓库
func (repository dockerHubRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
