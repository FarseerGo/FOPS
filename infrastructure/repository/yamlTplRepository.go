package repository

import (
	"fops/domain/k8s/yamlTpl"
	"fops/infrastructure/repository/agent/yamlTplAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() yamlTpl.Repository { return &yamlTplRepository{} })
}

type yamlTplRepository struct {
}

// ToList Yaml模板列表
func (repository yamlTplRepository) ToList() []yamlTpl.DomainObject {
	lst := yamlTplAgent.ToList()
	return mapper.Array[yamlTpl.DomainObject](lst)
}

// ToInfo Yaml模板信息
func (repository yamlTplRepository) ToInfo(id int) yamlTpl.DomainObject {
	po := yamlTplAgent.ToInfo(id)
	return mapper.Single[yamlTpl.DomainObject](po)
}

// Count Yaml模板数量
func (repository yamlTplRepository) Count() int64 {
	return yamlTplAgent.Count()
}

// Add 添加Yaml模板
func (repository yamlTplRepository) Add(do yamlTpl.DomainObject) int {
	po := mapper.Single[yamlTplAgent.PO](do)
	return yamlTplAgent.Add(po)
}

// Update 修改Yaml模板
func (repository yamlTplRepository) Update(id int, do yamlTpl.DomainObject) {
	po := mapper.Single[yamlTplAgent.PO](do)
	yamlTplAgent.Update(id, po)
}

// Delete 删除Yaml模板
func (repository yamlTplRepository) Delete(id int) {
	yamlTplAgent.Delete(id)
}
