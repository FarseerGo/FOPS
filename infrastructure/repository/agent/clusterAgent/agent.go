package clusterAgent

import "fops/infrastructure/repository/context"

// ToList 集群列表
func ToList() []PO {
	return context.NewContext().Cluster.Select("Id", "Name", "RuntimeEnvType").Asc("Sort").ToList()
}

// ToInfo 集群信息
func ToInfo(id int) PO {
	return context.NewContext().Cluster.Where("Id = ?", id).ToEntity()
}

// Count 集群数量
func Count() int64 {
	return context.NewContext().Cluster.Count()
}

// Add 添加集群
func Add(po PO) int {
	context.NewContext().Cluster.Insert(&po)
	return po.Id
}

// Update 修改集群
func Update(id int, po PO) int64 {
	return context.NewContext().Cluster.Where("Id = ?", id).Update(po)
}

// Delete 删除集群
func Delete(id int) int64 {
	return context.NewContext().Cluster.Where("Id = ?", id).Delete()
}
