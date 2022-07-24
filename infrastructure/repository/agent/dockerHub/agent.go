package dockerHub

import "fops/infrastructure/repository/context"

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList DockerHub列表
func (agent) ToList() []PO {
	return context.NewContext().DockerHub.Desc("Id").ToList()
}

// ToInfo DockerHub信息
func (agent) ToInfo(id int) PO {
	return context.NewContext().DockerHub.Where("Id = ?", id).ToEntity()
}

// Count DockerHub数量
func (agent) Count() int64 {
	return context.NewContext().DockerHub.Count()
}

// Add 添加仓库
func (agent) Add(po PO) {
	context.NewContext().DockerHub.Insert(&po)
}

// Update 修改仓库
func (agent) Update(id int, po PO) {
	context.NewContext().DockerHub.Where("Id = ?", id).Update(po)
}

// Delete 删除仓库
func (agent) Delete(id int) {
	context.NewContext().DockerHub.Where("Id = ?", id).Delete()
}
