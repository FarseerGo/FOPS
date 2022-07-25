package repository

import (
	"fops/domain/metaData/projectGroup"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() projectGroup.Repository {
		return &projectGroupRepository{data.Init[context.MysqlContext]().ProjectGroup}
	})
}

type projectGroupRepository struct {
	data.TableSet[model.ProjectGroupPO]
}

// ToList 项目组列表
func (repository *projectGroupRepository) ToList() []projectGroup.DomainObject {
	lstPO := repository.Desc("Id").ToList()
	return mapper.Array[projectGroup.DomainObject](lstPO)
}

// ToInfo 项目组信息
func (repository *projectGroupRepository) ToInfo(id int) *projectGroup.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	do := mapper.Single[projectGroup.DomainObject](po)
	return &do
}

//// Count 项目组数量
//func (repository *projectGroupRepository) Count() int64 {
//	return repository.Count()
//}

// Add 添加项目组
func (repository *projectGroupRepository) Add(do projectGroup.DomainObject) int {
	po := mapper.Single[model.ProjectGroupPO](do)
	repository.Insert(&po)
	return po.Id
}

// Update 修改项目组
func (repository *projectGroupRepository) Update(id int, do projectGroup.DomainObject) {
	po := mapper.Single[model.ProjectGroupPO](do)
	repository.Where("Id = ?", id).Update(po)
}

// Delete 删除项目组
func (repository *projectGroupRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
