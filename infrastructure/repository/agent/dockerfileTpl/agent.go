package dockerfileTpl

import "fops/infrastructure/repository/context"

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList Dockerfile模板列表
func (agent) ToList() []PO {
	return context.NewContext().DockerfileTpl.ToList()
}

// ToInfo Dockerfile模板信息
func (agent) ToInfo(id int) PO {
	return context.NewContext().DockerfileTpl.Where("Id = ?", id).ToEntity()
}

// Count Dockerfile模板数量
func (agent) Count() int64 {
	return context.NewContext().DockerfileTpl.Count()
}

// Add 添加Dockerfile模板
func (agent) Add(po PO) {
	context.NewContext().DockerfileTpl.Insert(&po)
}

// Update 修改Dockerfile模板
func (agent) Update(id int, po PO) {
	context.NewContext().DockerfileTpl.Where("Id = ?", id).Update(po)
}

// Delete 删除Dockerfile模板
func (agent) Delete(id int) {
	context.NewContext().DockerfileTpl.Where("Id = ?", id).Delete()
}
