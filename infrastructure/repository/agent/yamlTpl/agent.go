package yamlTpl

import "fops/infrastructure/repository/context"

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList Yaml模板列表
func (agent) ToList() []PO {
	return context.NewContext().YamlTpl.Desc("Id").ToList()
}

// ToInfo Yaml模板信息
func (agent) ToInfo(id int) PO {
	return context.NewContext().YamlTpl.Where("Id = ?", id).ToEntity()
}

// Count Yaml模板数量
func (agent) Count() int64 {
	return context.NewContext().YamlTpl.Count()
}

// Add 添加Yaml模板
func (agent) Add(po PO) int {
	context.NewContext().YamlTpl.Insert(&po)
	return po.Id
}

// Update 修改Yaml模板
func (agent) Update(id int, po PO) {
	context.NewContext().YamlTpl.Where("Id = ?", id).Update(po)
}

// Delete 删除Yaml模板
func (agent) Delete(id int) {
	context.NewContext().YamlTpl.Where("Id = ?", id).Delete()
}
