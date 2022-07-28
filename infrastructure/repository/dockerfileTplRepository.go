package repository

import (
	"fops/domain/metaData/dockerfileTpl"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/data"
	"github.com/farseernet/farseer.go/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() dockerfileTpl.Repository {
		return &dockerfileTplRepository{data.Init[context.MysqlContext]("fops").DockerfileTpl}
	})
}

type dockerfileTplRepository struct {
	data.TableSet[model.DockerfileTplPO]
}

// ToList Dockerfile模板列表
func (repository dockerfileTplRepository) ToList() []dockerfileTpl.DomainObject {
	lstPO := repository.ToList()
	return mapper.Array[dockerfileTpl.DomainObject](lstPO)
}

// ToInfo Dockerfile模板信息
func (repository dockerfileTplRepository) ToInfo(id int) dockerfileTpl.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	return mapper.Single[dockerfileTpl.DomainObject](po)
}

//// Count Dockerfile模板数量
//func (repository dockerfileTplRepository) Count() int64 {
//	return repository.Count()
//}

// Add 添加Dockerfile模板
func (repository dockerfileTplRepository) Add(do dockerfileTpl.DomainObject) {
	po := mapper.Single[model.DockerfileTplPO](do)
	repository.Insert(&po)
}

// Update 修改Dockerfile模板
func (repository dockerfileTplRepository) Update(id int, do dockerfileTpl.DomainObject) {
	po := mapper.Single[model.DockerfileTplPO](do)
	repository.Where("Id = ?", id).Update(po)
}

// Delete 删除Dockerfile模板
func (repository dockerfileTplRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
