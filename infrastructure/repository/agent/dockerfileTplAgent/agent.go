package dockerfileTplAgent

import "fops/infrastructure/repository/context"

// ToList Dockerfile模板列表
func ToList() []PO {
	return context.NewContext().DockerfileTpl.ToList()
}

// ToInfo Dockerfile模板信息
func ToInfo(id int) PO {
	return context.NewContext().DockerfileTpl.Where("Id = ?", id).ToEntity()
}

// Count Dockerfile模板数量
func Count() int64 {
	return context.NewContext().DockerfileTpl.Count()
}

// Add 添加Dockerfile模板
func Add(po PO) {
	context.NewContext().DockerfileTpl.Insert(&po)
}

// Update 修改Dockerfile模板
func Update(id int, po PO) {
	context.NewContext().DockerfileTpl.Where("Id = ?", id).Update(po)
}

// Delete 删除Dockerfile模板
func Delete(id int) {
	context.NewContext().DockerfileTpl.Where("Id = ?", id).Delete()
}
