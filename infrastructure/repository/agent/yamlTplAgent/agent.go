package yamlTplAgent

import "fops/infrastructure/repository/context"

// ToList Yaml模板列表
func ToList() []PO {
	return context.NewContext().YamlTpl.Desc("Id").ToList()
}

// ToInfo Yaml模板信息
func ToInfo(id int) PO {
	return context.NewContext().YamlTpl.Where("Id = ?", id).ToEntity()
}

// Count Yaml模板数量
func Count() int64 {
	return context.NewContext().YamlTpl.Count()
}

// Add 添加Yaml模板
func Add(po PO) int {
	context.NewContext().YamlTpl.Insert(&po)
	return po.Id
}

// Update 修改Yaml模板
func Update(id int, po PO) {
	context.NewContext().YamlTpl.Where("Id = ?", id).Update(po)
}

// Delete 删除Yaml模板
func Delete(id int) {
	context.NewContext().YamlTpl.Where("Id = ?", id).Delete()
}
