package repository

import (
	"fops/domain/k8s/yamlTpl"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() yamlTpl.Repository { return &yamlTplRepository{data.Init[context.MysqlContext]().YamlTpl} })
}

type yamlTplRepository struct {
	data.TableSet[model.YamlTplPO]
}

// ToList Yaml模板列表
func (repository yamlTplRepository) ToList() []yamlTpl.DomainObject {
	lst := repository.Desc("Id").ToList()
	return mapper.Array[yamlTpl.DomainObject](lst)
}

// ToInfo Yaml模板信息
func (repository yamlTplRepository) ToInfo(id int) yamlTpl.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	return mapper.Single[yamlTpl.DomainObject](po)
}

//// Count Yaml模板数量
//func (repository yamlTplRepository) Count() int64 {
//	return repository.Count()
//}

// Add 添加Yaml模板
func (repository yamlTplRepository) Add(do yamlTpl.DomainObject) int {
	po := mapper.Single[model.YamlTplPO](do)
	repository.Insert(&po)
	return po.Id
}

// Update 修改Yaml模板
func (repository yamlTplRepository) Update(id int, do yamlTpl.DomainObject) {
	po := mapper.Single[model.YamlTplPO](do)
	repository.Where("Id = ?", id).Update(po)
}

// Delete 删除Yaml模板
func (repository yamlTplRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
