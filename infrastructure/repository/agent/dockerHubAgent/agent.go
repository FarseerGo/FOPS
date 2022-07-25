package dockerHubAgent

import "fops/infrastructure/repository/context"

// ToList DockerHub列表
func ToList() []PO {
	return context.NewContext().DockerHub.Desc("Id").ToList()
}

// ToInfo DockerHub信息
func ToInfo(id int) PO {
	return context.NewContext().DockerHub.Where("Id = ?", id).ToEntity()
}

// Count DockerHub数量
func Count() int64 {
	return context.NewContext().DockerHub.Count()
}

// Add 添加仓库
func Add(po PO) {
	context.NewContext().DockerHub.Insert(&po)
}

// Update 修改仓库
func Update(id int, po PO) {
	context.NewContext().DockerHub.Where("Id = ?", id).Update(po)
}

// Delete 删除仓库
func Delete(id int) {
	context.NewContext().DockerHub.Where("Id = ?", id).Delete()
}
